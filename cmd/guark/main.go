// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/guark/guark/cmd/guark/actions"
	"github.com/urfave/cli/v2"
)

var app *cli.App

func init() {

	app = &cli.App{
		Name:  "Guark",
		Usage: "Guark framework command line interface.",
		// Flags: []cli.Flag{
		// 	&cli.StringFlag{
		// 		Name:  "get",
		// 		Usage: "Get from config.", // moved to `guark config` `guark config id` `guark config window width`
		// 	},
		// },
		Commands: []*cli.Command{
			{
				Name:   "build",
				Usage:  "Build guark app.",
				Flags:  actions.BuildFlags,
				Before: actions.CheckWorkingDir,
				Action: actions.Build,
			},
			{
				Name:   "run",
				Usage:  "Build and run guark app.",
				Before: actions.CheckWorkingDir,
				Action: run,
			},
			{
				Name:   "dev",
				Usage:  "Start dev app.",
				Flags:  actions.DevFlags,
				Before: actions.CheckWorkingDir,
				Action: actions.Dev,
			},
			{
				Name:   "generate",
				Usage:  "Generate embedable static files and assets.",
				Before: actions.CheckWorkingDir,
				Action: actions.Generate,
			},
		},
	}
}

func run(c *cli.Context) error {
	fmt.Println("run!")
	return nil
}

func main() {

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
