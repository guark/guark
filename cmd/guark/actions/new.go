// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions


import (
	"io"
	"os"
	"fmt"
	"github.com/urfave/cli/v2"

	"github.com/guark/guark/cmd/guark/stdio"
)


var (
	NewFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  "template",
			Aliases: []string{"from"},
			Usage: "init new project from a remote template.",
		},
		&cli.StringFlag{
			Name: "dest",
			Usage: "template destination path.",
		},
		&cli.StringFlag{
			Name: "mod",
			Usage: "Your app module name.",
		},
	}
)


func New(c *cli.Context) (err error) {

	var (
		out = stdio.NewWriter()
		dest = path(c.String("dest"))
		content []byte
		template string
	)

	if template = c.String("template"); template == "" {

		err = fmt.Errorf("Template required. for eg: `guark init --template vue`")
		return

	} else if isCleanDir(dest) == false {

		err = fmt.Errorf("Destination path %s is not empty!", dest)
		return
	}

	if IsUrl(template) == false {
		template = fmt.Sprintf("https://github.com/guark/%s", template)
	}

	out.Update("Checking remote template")

	// Validate remote repo is a valid guark template.
	if _, err = GitFile(template, "guark.yaml", os.Getenv("GUARK_GIT_AUTH")); err != nil {
		return
	}

	out.Done("Remote template is valid")

	fmt.Println(content)
	return nil

}


func isCleanDir(dir string) bool {

	var (
		d *os.File
		err error
	)

	if d, err = os.Open(dir); err != nil {
		return true
	}

	defer d.Close()

	_, err = d.Readdir(1)

	return err == io.EOF
}
