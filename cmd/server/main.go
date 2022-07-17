package main

import (
	"log"
	"net"
	"os"

	"github.com/urfave/cli/v2"
	"word-of-wisdom/internal/service/quotes"
	"word-of-wisdom/internal/workers/tcpserver"
)

func main() {
	app := &cli.App{
		Name:        "Server quotes",
		Description: "TCP server quotes, protected from DDOS attacks with the Prof of Work",
		Version:     "1.0.0",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "addr",
				Value:   ":12345",
				Usage:   "listen address",
				EnvVars: []string{"ADDR"},
			},
			&cli.IntFlag{
				Name:    "complexity",
				Value:   5,
				Usage:   "complexity algorithm for proof of work",
				EnvVars: []string{"COMPLEXITY"},
			},
		},
		Action: func(ctx *cli.Context) error {
			addr, err := net.ResolveTCPAddr("tcp", ctx.String("addr"))
			if err != nil {
				log.Fatal("resolve tcp address", err)
			}
			complexity := uint8(ctx.Int("complexity"))
			server := tcpserver.New(addr, quotes.New(), complexity)
			return server.Run(ctx.Context)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
