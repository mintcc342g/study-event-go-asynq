package controllers

import (
	"net/http"
	"strings"

	"study-event-go-asynq/applications"
	"study-event-go-asynq/applications/dto"

	"github.com/labstack/echo/v4"
)

// AnnouncementController ...
type AnnouncementController struct {
	announcementSvc *applications.AnnouncementService
}

// NewAnnouncementController ...
func NewAnnouncementController(announcementSvc *applications.AnnouncementService) *AnnouncementController {
	return &AnnouncementController{
		announcementSvc: announcementSvc,
	}
}

// Schedule ...
func (a *AnnouncementController) Schedule(c echo.Context) (err error) {
	// TODO: change logger

	ctx := c.Request().Context()

	var request struct {
		From    string `json:"from"`
		Message string `json:"message"`
	}
	if err = c.Bind(&request); err != nil {
		c.Logger().Error("AnnouncementController Bind", "err", err)
		return response(c, http.StatusBadRequest, "invalid request", nil)
	}

	announcementDTO := dto.Announcement{
		Message: strings.TrimSpace(request.Message),
		From:    request.From,
	}

	res, err := a.announcementSvc.Schedule(ctx, announcementDTO)
	if err != nil {
		c.Logger().Error("AnnouncementController Schedule", "err", err)
		return response(c, http.StatusInternalServerError, "internal server error", err.Error())
	}

	return response(c, http.StatusOK, "Schedule OK", res)
}
