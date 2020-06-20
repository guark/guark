// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/guark/guark/cmd/guark/utils"
)

// Front end UI builder
type UIBuilder struct {

	// Task/Package manager.
	Pkg string

	// Main builder.
	Main *Build
}

func (b UIBuilder) Before() error {

	b.Main.Log.Update("Building app ui.")
	return nil
}

func (b UIBuilder) Build() error {

	cmd := exec.Command(b.Pkg, "build")
	cmd.Dir = utils.Path("ui")
	cmd.Env = append(os.Environ(), fmt.Sprintf("GUARK_BUILD_DIR=%s/ui", b.Main.Temp))

	out, err := cmd.CombinedOutput()

	if err != nil {

		b.Main.Log.Err(string(out))
		return err
	}

	b.Main.Log.Done("Guark ui builded 🙈")
	return nil
}

func (b UIBuilder) Cleanup() {

}
