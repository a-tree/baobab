package echohandler

import (
	"baobab/internal/place"
	"baobab/internal/preserver"
	"net/http"

	"github.com/labstack/echo/v5"
)

type PlaceHandler struct {
	repo preserver.PlaceRepository
}

func NewPlaceHandler(repo preserver.PlaceRepository) *PlaceHandler {
	return &PlaceHandler{repo: repo}
}

func (h *PlaceHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/places", h.CreatePlace)
	e.GET("/places", h.GetPlaces)
}

func (h *PlaceHandler) CreatePlace(c *echo.Context) error {
	place := new(place.Place)
	if err := c.Bind(place); err != nil {
		return err
	}
	if err := h.repo.Create(place); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, place)
}

func (h *PlaceHandler) GetPlaces(c *echo.Context) error {
	places, _ := h.repo.GetAll()
	return c.JSON(http.StatusOK, places)
}
