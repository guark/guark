// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import "github.com/guark/guark/cmd/guark/utils"

// Builders/Compilers/Packagers interface.
type Builder interface {

	// Setup before run.
	Before() error

	// Run the builder.
	Run() error

	// Cleanup build.
	Cleanup()
}

type GuarkConfig struct {
	GuarkFileVersion string `yaml:"guark"`
	Version          string `yaml:"version"` // App version
	ID               string `yaml:"id"`
	Name             string `yaml:"name"`
	License          string `yaml:"license"`
	LogLevel         string `yaml:"logLevel"`
	EngineName       string `yaml:"engineName"`
}

type Build struct {

	// App build info.
	Info GuarkConfig

	// Build Config
	Config struct {

		// RC for linux.
		Linux struct {
			CC      string
			CXX     string
			Ldflags string
		}

		// RC for darwin.
		Darwin struct {
			CC      string
			CXX     string
			Ldflags string
		}

		// Build config for windows.
		Windows struct {
			CC      string
			CXX     string
			Windres string
			Ldflags string
		}
	}

	// build temp dir.
	Temp string

	// build logs output.
	Log *utils.Output

	// Build dest.
	Dest string

	// Clean dest dir if exists.
	Clean bool

	// Build targets.
	Targets []string

	// Builders
	Builders []Builder

	// Builder funcs
	BeforeFunc  func(*Build) error
	RunFunc     func(*Build) error
	CleanupFunc func(*Build)
}

func (b *Build) Before() error {
	return b.BeforeFunc(b)
}

func (b *Build) Run() error {
	return b.RunFunc(b)
}

func (b *Build) Cleanup() {
	b.CleanupFunc(b)
}
