// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	"github.com/guark/guark/cmd/guark/stdio"
	"github.com/guark/guark/internal/generator"
	"github.com/urfave/cli/v2"
	"github.com/zserge/webview"
)

var (
	buildTmpDir   string
	supportedOses = []string{"linux", "darwin", "windows"}
	BuildFlags    = []cli.Flag{
		&cli.StringFlag{
			Name:  "pkg",
			Usage: "Set your package manager.",
			Value: "yarn",
		},
		&cli.StringSliceFlag{
			Name:  "target",
			Usage: "Set build targets (linux, darwin, windows)",
			Value: cli.NewStringSlice(supportedOses...),
		},
	}
)

// SEE: https://github.com/zserge/webview/issues/22
// Install  binutils-mingw-w64
// sudo dnf install mingw64-gcc

func Build(c *cli.Context) (err error) {

	out := stdio.NewWriter()
	targets := c.StringSlice("target")
	defer out.Writer.Stop()

	for i := range targets {

		if err = checkTarget(targets[i]); err != nil {
			return
		}
	}

	buildTmpDir, err = ioutil.TempDir("", "guark")

	if err != nil {
		return
	}

	defer os.RemoveAll(buildTmpDir)

	out.End("Guark build initialized ⚙️", "")

	if err = buildUI(c.String("pkg"), buildTmpDir); err != nil {
		return
	}

	out.End("Guark UI builded 🙈", "")

	staticDir := filepath.Join(buildTmpDir, "static")

	if err = index(buildTmpDir, staticDir); err != nil {
		return
	}

	out.End("Guark UI indexed 🙉", "")

	for i := range targets {

		if err = build(targets[i], buildTmpDir); err != nil {
			return
		}

		out.End(fmt.Sprintf("Guark build for %s 🙊", targets[i]), "")
	}

	defer out.End("Guark build finished 🚀🚀", "")

	return
}

func build(target string, dir string) error {

	buildDir, err := getBuildDir(target, dir)

	if err != nil {
		return err
	}

	flags, env := getBuildFlagsAndEnvFor(target, buildDir)

	cmd := exec.Command("go", flags...)
	cmd.Env = append(os.Environ(), env...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func getDlls() string {
	return filepath.Join(os.Getenv("GOPATH"), "src", pkgPath(webview.New(true)), "dll")
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

func buildUI(pkg string, dir string) error {

	cmd := exec.Command(pkg, "build")
	cmd.Dir = path("ui")
	cmd.Env = append(os.Environ(), fmt.Sprintf("GUARK_BUILD_DIR=%s/ui", dir))
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func getBuildDir(target string, dir string) (string, error) {

	buildDir := filepath.Join(dir, target)

	if err := os.Mkdir(buildDir, 0740); err != nil {
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

func index(dir string, staticDir string) error {

	if err := os.Mkdir(staticDir, 0740); err != nil {
		return err
	}

	return generator.Assets(filepath.Join(dir, "ui"), staticDir, path("lib", "static.go"), "lib", filepath.Join(dir, "ui"))
}
