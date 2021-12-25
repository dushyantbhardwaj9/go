package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "boom",
		Usage: "make an explosive entrance",
		Action: func(c *cli.Context) error {
			sayThis()
			return nil
		},
		// Action: sayThis(c * cli.Context),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func sayThis() {
	fmt.Println("Yuuuhhuuuu!")

}
