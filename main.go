package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "yaml2excel",
		Usage: "yaml2excel <yaml filename>",
		Action: func(c *cli.Context) error {
			fmt.Println("Test")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

