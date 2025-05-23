package main

import (
	"log"
	"os"

	"github.com/labbs/mynotes/pkg/cmd/migration"
	"github.com/labbs/mynotes/pkg/cmd/server"
	"github.com/urfave/cli/v2"
)

var version = "development"

func main() {
	app := cli.NewApp()
	app.Name = "MyNotes"
	app.Usage = ""
	app.Version = version

	app.Commands = []*cli.Command{
		server.NewInstance(),
		migration.NewInstance(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
