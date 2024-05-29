package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"watchy/internal/model"
	"watchy/internal/service"

	"github.com/labstack/echo/v4"
)

type Controller interface {
	GetWatchEvents(c echo.Context) error
	CreateWatchEvent(c echo.Context) error
}

type WatchEventController struct {
	Service service.WatchEventService
}

func NewWatchEventController(WatchEventService service.WatchEventService) Controller {
	return WatchEventController{
		Service: WatchEventService,
	}
}

// function to get specific user watch events
func (e WatchEventController) GetWatchEvents(c echo.Context) error {
	userID := c.Param("user_id")

	resp, err := e.Service.GetWatchEvent(c.Request().Context(), userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, resp)
}

// function to insert user watch events to DB
func (e WatchEventController) CreateWatchEvent(c echo.Context) error {
	b, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	m := model.WatchEvent{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	m, err = e.Service.CreateWatchEvent(c.Request().Context(), m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, m)
}
