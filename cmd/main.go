package main

import (
	"fmt"
	"os"

	"github.com/asyrawih/manga/config"
	"github.com/asyrawih/manga/handlers"
	"github.com/asyrawih/manga/pkg/dbconn"
	"github.com/asyrawih/manga/repositories"
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

func RunServer(cfg string, port string) error {

	e := echo.New()
	c := config.LoadConfig(cfg)

	db, err := dbconn.NewMySQLDB(c)
	if err != nil {
		return err
	}

	repositories.NewUserRepo(db)

	handler := handlers.NewHttpHandler()
	e.GET("/", handler.CreateUser)
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
