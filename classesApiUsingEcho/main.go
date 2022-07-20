package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zahraakaraki/MyApp/handlers"

	_ "github.com/lib/pq"
)

func main() {

	var err error
	db, err := sql.Open("postgres", "user=postgres password=karaki217 dbname=stc sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/class", handlers.GetClasses(db))
	e.Logger.Fatal(e.Start(":8080"))

}
