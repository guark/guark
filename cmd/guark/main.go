// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package main

import (
	"log"
	"os"

	"github.com/guark/guark/cmd/guark/actions"
	"github.com/guark/guark/cmd/guark/utils"
	"github.com/urfave/cli/v2"
)

var app *cli.App

func init() {

	app = &cli.App{
		Name:  "Guark",
		Usage: "Guark framework CLI tool.",
		Commands: []*cli.Command{
			{
				Name:   "build",
				Usage:  "Build guark app.",
				Flags:  actions.BuildFlags,
				Before: utils.CheckWorkingDir,
				Action: actions.Build,
			},
			{
				Name:   "bundle",
				Usage:  "Bundle guark app.",
				Before: utils.CheckWorkingDir,
				Action: actions.Bundle,
			},
			{
				Name:   "run",
				Usage:  "Start and run dev app.",
				Flags:  actions.DevFlags,
				Before: utils.CheckWorkingDir,
				Action: actions.Dev,
			},
			{
				Name:   "init",
				Usage:  "Initialize a new guark project.",
				Flags:  actions.NewFlags,
				Action: actions.New,
			},
		},
	}
}

func main() {

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
