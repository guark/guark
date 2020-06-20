// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package builders

import "github.com/guark/guark/cmd/guark/stdio"

// Builders/Compilers interface.
type Builder interface {

	// setup before build.
	Before() error

	// build the thing.
	Build() error

	// Cleanup build.
	Cleanup()
}

type Build struct {

	// App build info.
	Info struct {
		GuarkFileVersion string `yaml:"guark"`
		Version          string `yaml:"version"`
		ID               string `yaml:"id"`
		Name             string `yaml:"name"`
		License          string `yaml:"license"`
		LogLevel         string `yaml:"logLevel"`
	}

	// build temp dir.
	Temp string

	// build logs output.
	Log *stdio.Output

	// Build dest.
	Dest string

	// Clean dest dir if exists.
	Clean bool

	// Build targets.
	Targets []string

	// Builders
	Builders []Builder

	BeforeFunc  func(*Build) error
	BuildFunc   func(*Build) error
	CleanupFunc func(*Build)
}

func (b *Build) Before() error {
	return b.BeforeFunc(b)
}

func (b *Build) Build() error {
	return b.BuildFunc(b)
}

func (b *Build) Cleanup() {
	b.CleanupFunc(b)
}

func Run(builder Builder) (err error) {

	defer builder.Cleanup()

	if err = builder.Before(); err != nil {
		return
	}

	return builder.Build()
}
