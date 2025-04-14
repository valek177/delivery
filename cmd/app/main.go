package main

import (
	"context"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gitlab.com/microarch-ru/delivery/cmd"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()

	port := getEnvVariable("HTTP_PORT")

	app := cmd.NewCompositionRoot(
		ctx,
	)
	startWebServer(app, port)
}

func startWebServer(compositionRoot cmd.CompositionRoot, port string) {
	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%s", port)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getEnvVariable(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	env := os.Getenv(key)
	return env
}
