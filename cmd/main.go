package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ynrfin.github.com/golang-warehouse-marketplace-api/controllers"
	"ynrfin.github.com/golang-warehouse-marketplace-api/models"
	"ynrfin.github.com/golang-warehouse-marketplace-api/repositories"
)

func main() {
	e := echo.New()
	log.Println("request received")
	ENV_PGUSER := os.Getenv("PGUSER")
	ENV_PGPASS := os.Getenv("PGPASSWORD")

	ENV_PGHOST := os.Getenv("PGHOST")
	ENV_PGPORT := os.Getenv("PGPORT")
	ENV_PGDB := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", ENV_PGHOST, ENV_PGUSER, ENV_PGPASS, ENV_PGDB, ENV_PGPORT)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("error connecting to pg", err)
	}

	UserRepo := repositories.UserRepository{Db: db}
	UserHandler := controllers.UserHandler{UserRepo: UserRepo}

	e.GET("/", func(c echo.Context) error {
		users := []models.User{}
		db.Find(&users)
		return c.JSON(http.StatusOK, users)
	})
	e.GET("/users", UserHandler.HandleListUser)

	e.Start(":8080")
}
