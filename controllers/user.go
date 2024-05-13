package controllers

import (
	"net/http"

	"github.com/google/uuid"
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

func (h UserHandler) HandleGetUser(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id empty")
	}
	if uuid.Validate(id) != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}
	idUUID, err := uuid.Parse(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error converting string to UUID. "+err.Error())
	}
	users, err := h.UserRepo.GetUserById(idUUID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}
