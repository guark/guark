// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/guark/guark/cmd/guark/builders"
	. "github.com/guark/guark/cmd/guark/utils"
	"github.com/guark/guark/utils"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
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
		&cli.BoolFlag{
			Name:  "keep-tmp",
			Usage: "Remove dest path if exists.",
		},
	}
)

func before(b *builders.Build) (err error) {

	// Unmarshal guark file.
	if err = utils.UnmarshalGuarkFile(".", &b.Info); err != nil {
		return
	} else if err = unmarshalBuildFile(&b.Config); err != nil {
		return
	}

	// Check targets.
	for i := range b.Targets {

		if err = checkTarget(b.Targets[i]); err != nil {
			b.Log.Err("Invalid target name.")
			return
		}
	}

	// Handle if dest already exists.
	if utils.IsDir(b.Dest) {

		if !b.Clean {
			return fmt.Errorf(`Dest "%[1]s/" already exists. try with: "--rm" flag, or remove "%[1]s/".`, b.Dest)
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

	for _, target := range b.Targets {

		switch target {
		case "linux":
			b.Builders = append(b.Builders, &builders.LinuxBuilder{
				Build: b,
			})
			break
		case "darwin":
			b.Builders = append(b.Builders, &builders.DarwinBuilder{
				Build: b,
			})
			break
		case "windows":
			b.Builders = append(b.Builders, &builders.WindowsBuilder{
				Build: b,
			})
		}
	}

	b.Log.Done("Build Initialized   🔨")
	return
}

func build(b *builders.Build) (err error) {

	for _, builder := range b.Builders {

		if err = run(builder); err != nil {
			return
		}
	}

	return nil
}

func Build(c *cli.Context) (err error) {

	b := &builders.Build{
		Log:        NewWriter(),
		Dest:       c.String("dest"),
		Clean:      c.Bool("rm"),
		Targets:    c.StringSlice("target"),
		BeforeFunc: before,
		RunFunc:    build,
		CleanupFunc: func(b *builders.Build) {
			if b.Temp != "" && !c.Bool("keep-tmp") {
				os.RemoveAll(b.Temp)
			}
		},
	}

	b.Builders = []builders.Builder{
		&builders.UIBuilder{
			Pkg:   c.String("pkg"),
			Build: b,
		},
		&builders.EmbedBuilder{
			Build: b,
		},
		&builders.MetaBuilder{
			Build: b,
		},
	}

	return run(b)
}

func run(builder builders.Builder) (err error) {

	defer builder.Cleanup()

	if err = builder.Before(); err != nil {
		return
	}

	return builder.Run()
}

func unmarshalBuildFile(c interface{}) error {

	cnf, err := ioutil.ReadFile("guark-build.yaml")

	if err != nil {
		return nil
	}

	return yaml.Unmarshal(cnf, c)
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
