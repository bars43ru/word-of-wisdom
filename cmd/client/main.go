package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"word-of-wisdom/internal/workers/tcpclient"
)

func main() {
	app := &cli.App{
		Name:        "Quote client",
		Description: "Client for server quotes",
		Version:     "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Value:   "0.0.0.0:12345",
				Usage:   "host address",
				EnvVars: []string{"HOST"},
			},
		},
		Action: func(ctx *cli.Context) error {
			client := tcpclient.New(ctx.String("host"))
			return client.Run(ctx.Context)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
