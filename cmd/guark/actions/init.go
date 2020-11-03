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
		out      = utils.NewWriter()
		dest     = utils.Path(c.String("dest"))
		template string
	)

	defer out.Stop()

	if template = c.String("template"); template == "" {

		err = fmt.Errorf("Template required. for eg: `guark new --template vue`")
		return

	} else if isCleanDir(dest) == false {

		err = fmt.Errorf("Destination path %s is not empty!", dest)
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

	prompt := promptui.Prompt{
		Label:     "Change the app module name to yours:",
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

	err = runSetupCommands(out, dest)

	if err == nil {
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

func runSetupCommands(out *utils.Output, dir string) error {

	data, err := ioutil.ReadFile(filepath.Join(dir, "guark-build.yaml"))

	if err != nil {
		return err
	}

	sup := SetupCommands{}

	if err = yaml.Unmarshal(data, &sup); err != nil {
		return err
	}

	if len(sup.Commands) > 0 {
		err = confirmAndRunSetupCommands(out, sup, dir)
	}

	return err
}

func confirmAndRunSetupCommands(out *utils.Output, s SetupCommands, dest string) error {

	fmt.Println("⏺ Setup commands (Verify them before confirm):")

	for _, v := range s.Commands {
		ln := fmt.Sprintf(`  - "%s"`, v.Cmd)
		if v.Dir != "" {
			ln = fmt.Sprintf("%s (dir: %s)", ln, v.Dir)
		}
		fmt.Println(ln)
	}

	prompt := promptui.Prompt{
		Label:     "Do you want to run this commands on your machine (Enter for No)",
		IsConfirm: true,
		Validate: func(v string) error {

			if v == "y" {
				return fmt.Errorf("Are you sure? type uppercase Y.")
			}

			return nil
		},
	}

	yes, err := prompt.Run()

	if err != nil {

		return err

	} else if yes == "Y" {

		out.Update("Running setup commands...")

		for _, v := range s.Commands {
			out.Update(fmt.Sprintf("Running setup command: %s", v.Cmd))
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
