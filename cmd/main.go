package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"kratos_template/cmd/greeter"
)

func main() {
	app := cli.NewApp()
	app.Name = "kratos-template"
	app.Usage = "A DDD-style Kratos microservice template"
	app.Commands = []cli.Command{
		greeter.Command,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
