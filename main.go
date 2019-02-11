package main

import (
	"os"

	"github.com/graphql-services/idp/cmd"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GraphQL IDP"
	app.Usage = "..."
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		cmd.ServerCommand(),
	}

	app.Run(os.Args)
}
