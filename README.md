# Study Asynq

- Go 1.18
- Asynq V0

<br/>

## Run Server

- for MacOS and Windows

```bash
# 1. run redis
$ docker run -p 6379:6379 redis 

# 2. compile the project
$ make all


# 3. run the server for MacOS. Or You can use
# the debug mod of VSCode to click the button 'Launch' 
$ ./bin/study-event-go-asynq
```

<br/>

## Event Test

- cURL example
```bash
$ curl --location --request POST 'http://127.0.0.1:4569/api/v1/study-asynq/announcement/schedule' \
--header 'Content-Type: application/json' \
--data-raw '{
    "from": "The kirin",
    "message": "WAKARIMASU"
}'
```

- result on console
```bash
[ANNOUNCEMENT] GOT A NEW ANNOUNCEMENT FROM The kirin.
[ANNOUNCEMENT] THE MESSAGE IS ...
[ANNOUNCEMENT] "WAKARIMASU"
```

- Web UI : [http://localhost:8080/monitoring](http://localhost:8080/monitoring)
