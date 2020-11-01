// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"io/ioutil"

	"github.com/guark/guark/cmd/guark/builders"
	"github.com/guark/guark/utils"
	"github.com/melbahja/bundler/bundle"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

func Bundle(c *cli.Context) error {

	data, err := ioutil.ReadFile("bundler.yaml")
	if err != nil {
		return err
	}

	config := &builders.GuarkConfig{}
	if err = utils.UnmarshalGuarkFile(".", config); err != nil {
		return err
	}

	bundler := &bundle.Bundler{
		Data: map[string]interface{}{
			"Engine": config.EngineName,
		},
	}
	if err = yaml.Unmarshal(data, bundler); err != nil {
		return err
	}

	return bundler.Run()
}
