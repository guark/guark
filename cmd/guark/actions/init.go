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

	"github.com/guark/guark/cmd/guark/stdio"
	"github.com/manifoldco/promptui"
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
	}
)

// UI setup commands
type setup struct {
	Version  string `yaml:"guark"`
	Commands []struct {
		Cmd  string   `yaml:"cmd"`
		Args []string `yaml:"args"`
	} `yaml:"setup"`
}

func New(c *cli.Context) (err error) {

	var (
		out      = stdio.NewWriter()
		dest     = "" //path(c.String("dest"))
		template string
	)

	defer out.Stop()

	if template = c.String("template"); template == "" {

		err = fmt.Errorf("Template required. for eg: `guark new --template vue`")
		return

	} else if isCleanDir(dest) == false {

		err = fmt.Errorf("Destination path %s is not empty (you can add `--dest new_dir`)!", dest)
		return
	}

	// if IsUrl(template) == false {
	// 	template = fmt.Sprintf("https://github.com/guark/%s", template)
	// }

	out.Update("Checking remote template")

	// // Validate remote repo.
	// if _, err = GitFile(template, "guark.yaml", os.Getenv("GUARK_GIT_AUTH")); err != nil {
	// 	return
	// }

	out.Done("Remote template validated")
	out.Update(fmt.Sprintf("Downloading: %s", template))

	if err = clone(template, dest); err != nil {
		return
	}

	out.Done("Template downloaded successfully.")

	prompt := promptui.Prompt{
		Label:     "Type your app module name",
		Default:   "github.com/melbahja/myapp",
		AllowEdit: true,
		Templates: &promptui.PromptTemplates{
			Valid:   "⏺ {{ . | cyan }}: ",
			Success: `{{ green "✔"}} {{ cyan "App module name:" }} `,
		},
	}

	mod, err := prompt.Run()

	if err != nil {
		return
	}

	err = refactorMod(strings.TrimSpace(mod), dest)

	out.Done("App module name refactored")

	err = setupUI(dest)

	if err == nil {
		out.Done(fmt.Sprintf("Done! cd to %s and run `guark dev`.", dest))
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

		} else if f.IsDir() || strings.Contains(path, ".git/") {

			return nil
		}

		b, err := ioutil.ReadFile(path)

		if err != nil {
			return err
		}

		return ioutil.WriteFile(path, bytes.ReplaceAll(b, []byte("{{AppPkg}}"), []byte(mod)), f.Mode())
	})
}

func setupUI(dir string) error {

	data, err := ioutil.ReadFile(filepath.Join(dir, "ui", "guark-setup.yaml"))

	if err != nil {
		return err
	}

	sup := setup{}

	if err = yaml.Unmarshal(data, &sup); err != nil {
		return err
	}

	if len(sup.Commands) > 0 {
		err = confirmAndRun(sup, dir)
	}

	return err
}

func confirmAndRun(s setup, dest string) error {

	fmt.Println("⏺ UI setup commands (Verify them before confirm):")

	for _, v := range s.Commands {
		fmt.Println(fmt.Sprintf("  - %s %s", v.Cmd, strings.Join(v.Args, " ")))
	}

	prompt := promptui.Prompt{
		Label:     "Do you want to run setup commands on your machine (Enter for N)",
		IsConfirm: true,
		Validate: func(v string) error {

			if v == "y" {
				return fmt.Errorf("Are you sure? type uppercase Y.")
			}

			return nil
		},
		Templates: &promptui.PromptTemplates{
			Success: `{{ green "✔"}} {{ cyan "You allowed setup commands:" }} `,
		},
	}

	yes, err := prompt.Run()

	if err != nil {

		return err

	} else if yes == "Y" {

		for _, v := range s.Commands {
			if err = runSetupCommand(dest, v.Cmd, v.Args); err != nil {
				break
			}
		}
	}

	return err
}

func runSetupCommand(dir string, c string, args []string) error {
	cmd := exec.Command(c, args...)
	cmd.Dir = filepath.Join(dir, "ui")
	cmd.Stdout = os.Stdout
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
