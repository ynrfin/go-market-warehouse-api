package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func (h UserHandler) HandleCreateUser(c echo.Context) error {
	type UserCreateRequest struct {
		Name  string `validate:"required"`
		Email string `validate:"required,email"`
	}

	request := UserCreateRequest{}

	if err := c.Bind(&request); err != nil {
		log.Println("error binding request")

		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(request); err != nil {
		log.Println("error validating request")
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	_, err := h.UserRepo.GetUserByEmail(request.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusBadRequest, "email has been used "+err.Error())
	}

	// TODO
	// ID still has chance of collision, solution
	// - if the ID exists, retry 5 times
	insertRes := h.UserRepo.Db.
		Create(repositories.
			User{
			ID:    uuid.New(),
			Name:  request.Name,
			Email: request.Email})
	if insertRes.Error != nil {
		log.Println("error creating user")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "create user")
}
