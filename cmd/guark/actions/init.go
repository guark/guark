// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/guark/guark/cmd/guark/utils"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

var (
	NewFlags = []cli.Flag{
		&cli.StringFlag{
			Name:    "template",
			Aliases: []string{"from"},
			Usage:   "init new project from a remote template.",
		},
		&cli.StringFlag{
			Name:  "dest",
			Usage: "template destination path.",
		},
		&cli.StringFlag{
			Name:  "mod",
			Usage: "Your app module name.",
		},
		&cli.BoolFlag{
			Name:  "no-setup",
			Usage: "Do not run template setup commands",
		},
	}
)

// Setup commands
type SetupCommands struct {
	Commands []struct {
		Cmd string   `yaml:"cmd"`
		Dir string   `yaml:"dir,omitempty"`
		Env []string `yaml:"env"`
	} `yaml:"setup"`
}

func New(c *cli.Context) (err error) {

	var (
		out              = utils.NewWriter()
		dest             = utils.Path(c.String("dest"))
		module, template string
	)

	defer out.Stop()

	if template = c.String("template"); template == "" {

		err = fmt.Errorf("Template required. for eg: `guark new --template vue`")
		return

	} else if isCleanDir(dest) == false {

		err = fmt.Errorf("Destination path %s is not empty!", dest)
		return

	} else if module = c.String("mod"); module == "" {

		err = fmt.Errorf("App module name required. for eg: `guark new --mod example.com/usr/pkg`")
		return
	}

	if utils.IsUrl(template) == false {
		template = fmt.Sprintf("https://github.com/guark/%s", template)
	}

	out.Update("Checking remote template")

	// Validate remote repo.
	if _, err = utils.GitFile(template, "guark.yaml", os.Getenv("GUARK_GIT_AUTH")); err != nil {
		return
	}

	out.Done("Remote template validated")
	out.Update(fmt.Sprintf("Downloading: %s", template))

	if err = clone(template, dest); err != nil {
		return
	}

	out.Done("Template downloaded successfully.")

	if err = refactorMod(strings.TrimSpace(module), dest); err != nil {
		return
	}

	out.Done("App module name refactored")

	if err = runSetupCommands(out, dest, c.Bool("no-setup") == false); err == nil {
		out.Done(fmt.Sprintf("Done! cd to %s and run `guark run`.", dest))
	}

	return
}

func clone(repo string, dir string) error {
	cmd := exec.Command("git", "clone", repo, dir)
	return cmd.Run()
}

func refactorMod(mod string, dest string) error {

	return filepath.Walk(dest, func(path string, f os.FileInfo, err error) error {

		if err != nil {

			return err

		} else if f.IsDir() || strings.Contains(path, ".git") {

			return nil
		}

		b, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		return ioutil.WriteFile(path, bytes.ReplaceAll(b, []byte("{{AppPkg}}"), []byte(mod)), f.Mode())
	})
}

func runSetupCommands(out *utils.Output, dest string, runComamnds bool) error {

	data, err := ioutil.ReadFile(filepath.Join(dest, "guark-build.yaml"))

	if err != nil {
		return err
	}

	scmd := SetupCommands{}

	if err = yaml.Unmarshal(data, &scmd); err != nil {
		return err
	}

	if len(scmd.Commands) > 0 {
		err = showAndRunSetupCommands(out, scmd, dest, runComamnds)
	}

	return err
}

func showAndRunSetupCommands(out *utils.Output, s SetupCommands, dest string, runComamnds bool) (err error) {

	fmt.Println("⏺ Setup commands:")

	for _, v := range s.Commands {
		ln := fmt.Sprintf(`  - "%s"`, v.Cmd)
		if v.Dir != "" {
			ln = fmt.Sprintf("%s (dir: %s)", ln, v.Dir)
		}
		fmt.Println(ln)
	}

	if runComamnds {

		out.Update("Running setup commands...")

		for _, v := range s.Commands {
			out.Done(fmt.Sprintf("Running setup command: %s", v.Cmd))
			if err = runSetupCommand(filepath.Join(dest, v.Dir), strings.Split(v.Cmd, " "), v.Env); err != nil {
				break
			}
		}
	}

	return err
}

func runSetupCommand(dir string, c []string, env []string) error {
	cmd := exec.Command(c[0], c[1:]...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), env...)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func isCleanDir(dir string) bool {

	var (
		d   *os.File
		err error
	)

	if d, err = os.Open(dir); err != nil {
		return true
	}

	defer d.Close()

	_, err = d.Readdir(1)

	return err == io.EOF
}
