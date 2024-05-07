package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ynrfin.github.com/golang-warehouse-marketplace-api/models"
)

func main() {
	e := echo.New()
	log.Println("request received")
	// ENV_PGUSER := os.Getenv("PGUSER")
	ENV_PGPASS := os.Getenv("PGPASSWORD")

	ENV_PGHOST := os.Getenv("PGHOST")
	ENV_PGPORT := os.Getenv("PGPORT")
	ENV_PGDB := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", ENV_PGHOST, "postgres", ENV_PGPASS, ENV_PGDB, ENV_PGPORT)
	// dsn := fmt.Sprint("host=local-pg-16 user=%s password=%s dbname=%s port=%s sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("error connecting to pg", err)
	}

	e.GET("/", func(c echo.Context) error {
		// Assuming that database has table called users like this

		// CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
		// CREATE TABLE users (
		//     id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
		//     name VARCHAR NOT NULL
		// );
		//
		//
		// INSERT INTO users (name) VALUES ('harry');
		// INSERT INTO users (name) VALUES ('jack');
		// INSERT INTO users (name) VALUES ('winston');
		users := []models.User{}
		db.Find(&users)
		return c.JSON(http.StatusOK, users)
	})

	e.Start(":8080")
}
