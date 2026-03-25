package echohandler

import (
	"baobab/internal/preserver"
	"baobab/internal/user"
	"net/http"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	repo preserver.UserRepository
}

func NewUserHandler(repo preserver.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/users", h.CreateUser)
	e.GET("/users", h.GetUsers)
}

func (h *UserHandler) CreateUser(c *echo.Context) error {
	user := new(user.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if err := h.repo.Create(user); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUsers(c *echo.Context) error {
	users, _ := h.repo.GetAll()
	return c.JSON(http.StatusOK, users)
}
