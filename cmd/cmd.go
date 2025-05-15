package main

import (
	"log"
	"os"

	"github.com/labbs/mynotion/pkg/cmd/migration"
	"github.com/labbs/mynotion/pkg/cmd/server"
	"github.com/urfave/cli/v2"
)

var version = "development"

func main() {
	app := cli.NewApp()
	app.Name = "MyNotion"
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
