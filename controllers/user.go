package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"ynrfin.github.com/golang-warehouse-marketplace-api/repositories"
)

type (
	UserHandler struct {
		UserRepo repositories.UserRepository
	}
)

func HandleCreateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "HandleCreateUser")
}

func (h UserHandler) HandleListUser(c echo.Context) error {
	users, err := h.UserRepo.GetUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
