// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"os"

	"github.com/guark/guark/app/utils"
	"github.com/guark/guark/internal/generator"
	"github.com/urfave/cli/v2"
)

func Generate(c *cli.Context) error {

	return BuildStatics("build/ui", "build/dist/datadir/static")
}

func BuildStatics(src string, dest string) error {

	if utils.IsDir(dest) {
		os.RemoveAll(dest)
	}

	os.MkdirAll(dest, 0760)

	return generator.Assets(src, dest, "lib/static.go", "lib", src)
}
