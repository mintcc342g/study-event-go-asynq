package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"study-event-go-asynq/cmd/containers"
	"study-event-go-asynq/conf"
	"study-event-go-asynq/consts"
	"study-event-go-asynq/workers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

const (
	banner = "\n" +
		"     ___           ___           ___           ___           ___                    ___           ___           ___           ___           ___     \n" +
		"    /\\  \\         /\\  \\         /\\__\\         /\\  \\         |\\__\\                  /\\  \\         /\\  \\         |\\__\\         /\\__\\         /\\  \\    \n" +
		"   /::\\  \\        \\:\\  \\       /:/  /        /::\\  \\        |:|  |                /::\\  \\       /::\\  \\        |:|  |       /::|  |       /::\\  \\   \n" +
		"  /:/\\ \\  \\        \\:\\  \\     /:/  /        /:/\\:\\  \\       |:|  |               /:/\\:\\  \\     /:/\\ \\  \\       |:|  |      /:|:|  |      /:/\\:\\  \\  \n" +
		" _\\:\\~\\ \\  \\       /::\\  \\   /:/  /  ___   /:/  \\:\\__\\      |:|__|__            /::\\~\\:\\  \\   _\\:\\~\\ \\  \\      |:|__|__   /:/|:|  |__    \\:\\~\\:\\  \\ \n" +
		"/\\ \\:\\ \\ \\__\\     /:/\\:\\__\\ /:/__/  /\\__\\ /:/__/ \\:|__|     /::::\\__\\          /:/\\:\\ \\:\\__\\ /\\ \\:\\ \\ \\__\\     /::::\\__\\ /:/ |:| /\\__\\    \\:\\ \\:\\__\\\n" +
		"\\:\\ \\:\\ \\/__/    /:/  \\/__/ \\:\\  \\ /:/  / \\:\\  \\ /:/  /    /:/~~/~             \\/__\\:\\/:/  / \\:\\ \\:\\ \\/__/    /:/~~/~    \\/__|:|/:/  /     \\:\\/:/  /\n" +
		" \\:\\ \\:\\__\\     /:/  /       \\:\\  /:/  /   \\:\\  /:/  /    /:/  /                    \\::/  /   \\:\\ \\:\\__\\     /:/  /          |:/:/  /       \\::/  / \n" +
		"  \\:\\/:/  /     \\/__/         \\:\\/:/  /     \\:\\/:/  /     \\/__/                     /:/  /     \\:\\/:/  /     \\/__/           |::/  /        /:/  /  \n" +
		"   \\::/  /                     \\::/  /       \\::/__/                               /:/  /       \\::/  /                      /:/  /        /:/  /   \n" +
		"    \\/__/                       \\/__/         ~~                                   \\/__/         \\/__/                       \\/__/         \\/__/    \n" +
		" => Starting listen %s\n"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	configs := conf.Configs

	r := echoInit(configs)
	signal := sigInit(r)
	// db := dbInit(r, configs)
	asynqClient := initAsynqClient(configs)

	repoContainer := containers.InitInfrastructureContainer(asynqClient)
	svcContainer := containers.InitApplicationContainer(repoContainer)
	ctrlContainer := containers.InitControllerContainer(svcContainer, repoContainer)

	if err := initHandler(r, ctrlContainer, signal); err != nil {
		r.Logger.Error("initHandler Error")
		os.Exit(1)
	}

	if err := initAsynqServer(r, configs, repoContainer); err != nil {
		r.Logger.Error("initAsynqServer Error")
		os.Exit(1)
	}

	if !configs.GetBool("use_docker_compose") {
		if err := initAsynqWebUIServer(r, configs); err != nil {
			r.Logger.Error("initAsynqWebUIServer Error")
			os.Exit(1)
		}
	}

	startServer(configs, r)
}

func echoInit(configs *conf.ViperConfig) (r *echo.Echo) {
	r = echo.New()

	// Middleware
	r.Use(middleware.RemoveTrailingSlash())
	r.Use(middleware.Recover())

	// CORS
	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.POST, echo.GET, echo.PUT, echo.DELETE},
	}))

	r.HideBanner = true

	return r
}

func sigInit(r *echo.Echo) chan os.Signal {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)
	go func() {
		sig := <-quit
		r.Logger.Error("Got signal", sig)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := r.Shutdown(ctx); err != nil {
			r.Logger.Fatal(err)
		}
		signal.Stop(quit)
		close(quit)
	}()
	return quit
}

func initAsynqClient(configs *conf.ViperConfig) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr: configs.GetString("redis_host"),
	})
}

func initAsynqServer(r *echo.Echo, configs *conf.ViperConfig, repoContainer *containers.InfrastructureContainer) error {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr: configs.GetString("redis_host"),
			// DB:   0, // fixme: test
		},
		asynq.Config{
			Concurrency: 10,
		},
	)

	announcementWorker := workers.NewAnnouncementWorker(repoContainer.AnnouncementRepo)

	mux := asynq.NewServeMux()
	mux.HandleFunc(consts.AnnouncementTaskKey, announcementWorker.Announce)
	mux.HandleFunc(consts.AnnouncementTimeTaskKey, announcementWorker.AnnounceWithTime)

	go func() {
		if err := srv.Run(mux); err != nil {
			r.Logger.Fatal("asynq server error", err)
			panic(err)
		}
	}()

	return nil
}

func initAsynqWebUIServer(r *echo.Echo, configs *conf.ViperConfig) error {
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring", // RootPath specifies the root for asynqmon app
		RedisConnOpt: asynq.RedisClientOpt{Addr: configs.GetString("redis_host")},
	})

	// Note: We need the tailing slash when using net/http.ServeMux.
	http.Handle(h.RootPath()+"/", h)

	// Go to http://localhost:8080/monitoring to see asynqmon homepage.
	go func() {
		if err := http.ListenAndServe(configs.GetString("asynqmon_host"), nil); err != nil {
			r.Logger.Fatal("asynqmon error", err)
			panic(err)
		}
	}()

	return nil
}

func initHandler(r *echo.Echo, ctrlContainer *containers.ControllerContainer, signal <-chan os.Signal) error {
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())

	api := r.Group("/api")
	ver := api.Group("/v1")
	sys := ver.Group("/study-asynq")

	sys.POST("/announcement/schedule", ctrlContainer.AnnouncementCtrl.Schedule)

	return nil
}

func startServer(configs *conf.ViperConfig, r *echo.Echo) {
	apiServer := fmt.Sprintf("0.0.0.0:%d", configs.GetInt("port"))
	r.Logger.Debugf("Starting server, Listen[%s]", apiServer)

	fmt.Printf(banner, apiServer)
	if err := r.Start(apiServer); err != nil {
		r.Logger.Fatal(err)
	}
}
