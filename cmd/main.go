package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/asyrawih/manga/config"
	handler "github.com/asyrawih/manga/handlers"
	"github.com/asyrawih/manga/internal/services"
	"github.com/asyrawih/manga/pkg/dbconn"
	"github.com/asyrawih/manga/repositories"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

var commands = []*cli.Command{
	{
		Name: "start",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "locate the configuration file",
				Required: true,
			},
			&cli.IntFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "runnig port",
				Required: true,
			},
		},
		Action: func(ctx *cli.Context) error {
			config := ctx.String("config")
			port := ctx.Int("port")
			return RunServer(config, fmt.Sprintf(":%d", port))
		},
	},
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return err

	}
	return nil
}

func RunServer(cfg string, port string) error {
	// Prepare
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HideBanner = true
	c := config.LoadConfig(cfg)

	db, err := dbconn.NewMySQLDB(c)
	if err != nil {
		return err
	}

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "oke")
	})

	ur := repositories.NewUserRepo(db)
	us := services.NewUserServie(ur, c)

	// Users
	userHandler := handler.NewHttpHandler(us, c)
	userHandler.Routes(e)

	return e.Start(port)
}

func main() {

	app := &cli.App{
		Name:     "manga",
		Version:  "0.0.1",
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		// Make Sure Stdout must exit
		log.Err(err).Caller().Msg("")
		os.Exit(1)
	}
}