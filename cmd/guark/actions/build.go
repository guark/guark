// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/guark/guark"
	"github.com/guark/guark/app/utils"
	"github.com/guark/guark/cmd/guark/builders"
	"github.com/guark/guark/cmd/guark/stdio"
	"github.com/manifoldco/promptui"
	"github.com/urfave/cli/v2"
	"github.com/zserge/webview"
)

var (
	supportedOses = []string{"linux", "darwin", "windows"}
	BuildFlags    = []cli.Flag{
		&cli.StringFlag{
			Name:  "pkg",
			Usage: "Set your package manager.",
			Value: "yarn",
		},
		&cli.StringSliceFlag{
			Name:  "target",
			Usage: "Set build targets",
			Value: cli.NewStringSlice(runtime.GOOS),
		},
		&cli.StringFlag{
			Name:  "dest",
			Usage: "Build to a specific destination.",
			Value: "dist",
		},
		&cli.BoolFlag{
			Name:  "rm",
			Usage: "Remove dest path if exists.",
		},
	}
)

func before(b *builders.Build) (err error) {

	b.Log.Done("Starting...")

	// Unmarshal guark file.
	if err = guark.UnmarshalGuarkFile(&b.Info); err != nil {
		return
	}

	// check targets.
	for i := range b.Targets {

		if err = checkTarget(b.Targets[i]); err != nil {
			b.Log.Err("Invalid target name.")
			return
		}
	}

	// handle if dest already exists.
	if utils.IsDir(b.Dest) {

		if b.Clean {

		} else if err = confirmDeleteDest(b.Dest); err != nil {
			return
		}

		os.RemoveAll(b.Dest)
	}

	// Create new dest.
	if err = os.MkdirAll(b.Dest, 0754); err != nil {
		return
	}

	// Create tmp.
	if b.Temp, err = ioutil.TempDir("", "guark"); err != nil {
		return
	}

	// Append builders.

	b.Log.Done("Guark build initialized ⚙️")
	return
}

func build(b *builders.Build) (err error) {

	for _, builder := range b.Builders {

		if err = builders.Run(builder); err != nil {
			return
		}
	}

	return nil
}

func cleanup(b *builders.Build) {

	if b.Temp != "" {
		os.RemoveAll(b.Temp)
	}
}

func Build(c *cli.Context) (err error) {

	b := &builders.Build{
		Log:         stdio.NewWriter(),
		Dest:        c.String("dest"),
		Clean:       c.Bool("rm"),
		Targets:     c.StringSlice("target"),
		BeforeFunc:  before,
		BuildFunc:   build,
		CleanupFunc: cleanup,
	}

	b.Builders = []builders.Builder{
		&builders.UIBuilder{
			Pkg:  c.String("pkg"),
			Main: b,
		},
		&builders.EmbedBuilder{
			Main: b,
		},
		&builders.MetaBuilder{
			Main: b,
		},
	}

	return builders.Run(b)
}

func confirmDeleteDest(dest string) error {

	prompt := promptui.Prompt{
		Label:     fmt.Sprintf("Confirm deleting: %s", dest),
		IsConfirm: true,
		Validate: func(v string) error {

			if v == "y" {
				return fmt.Errorf("Are you sure? type uppercase Y.")
			}

			return nil
		},
		Templates: &promptui.PromptTemplates{
			Success: `{{ green "✔"}} {{ cyan "Delete existing dest:" }} `,
		},
	}

	if yes, err := prompt.Run(); yes != "Y" || err != nil {
		return fmt.Errorf("aborted")
	}

	return nil
}

// TODO: change value of "x64" to be based on build
func getDlls() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", pkgPath(webview.New(true)), "dll", "x64")
}

// this function code was stolen from:
// https://stackoverflow.com/a/60846213/5834438
func pkgPath(v interface{}) string {
	if v == nil {
		return ""
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		return val.Elem().Type().PkgPath()
	}
	return val.Type().PkgPath()
}

func getBuildDir(target string, dir string) (string, error) {

	buildDir := filepath.Join(dir, target)

	if err := os.Mkdir(buildDir, 0754); err != nil {
		return "", err
	}

	return buildDir, nil
}

func checkTarget(target string) error {

	for i := range supportedOses {
		if supportedOses[i] == target {
			return nil
		}
	}

	return fmt.Errorf("target: %s not supported yet!", target)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
