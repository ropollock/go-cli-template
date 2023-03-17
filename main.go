package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "cli-template",
		Usage: "Make a cool cli app.",
		Action: func(*cli.Context) error {
			fmt.Println("Hello world!")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
