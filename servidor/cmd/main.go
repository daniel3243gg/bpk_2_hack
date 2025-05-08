package main

import (
	"log"
	"sistema_artigos_bpk/api"
	"sistema_artigos_bpk/internal/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
	    log.Printf("Erro ao carregar o arquivo .env: %v", err)
	}

	config.LoadConfig("config.yml")
	e := echo.New()
	user, password, host, port, dbName := config.GetDatabaseConfig()
	db, err := config.GetDb(user, password, host, dbName, port)

	if err.Message != ""{
		e.StdLogger.Panic(err.Error)
	}
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	

	api.RegisterRoutes(e,db,[]byte("asasaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))

	e.Debug = true
	
	e.Logger.Fatal(e.Start(":3000"))
}