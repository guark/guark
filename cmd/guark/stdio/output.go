// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package stdio

import (
	"fmt"
	"time"

	colors "github.com/logrusorgru/aurora"
	"github.com/theckman/yacspin"
)

var cfg = &yacspin.Config{
	Frequency: 100 * time.Millisecond,
	CharSet:   yacspin.CharSets[14],
}

type Output struct {
	spinner *yacspin.Spinner
}

func (o Output) Update(s string) {

	if o.spinner.Active() == false {
		o.spinner.Start()
	}

	o.spinner.Message(fmt.Sprintf(" %s", colors.Bold(colors.Blue(s))))
}

func (o Output) Done(s string) {
	o.spinner.Stop()
	fmt.Println(fmt.Sprintf("%s %s", colors.Green("✔"), colors.Cyan(s)))
}

func (o Output) Err(s string) {
	o.spinner.Stop()
	fmt.Println(colors.Red(fmt.Sprintf("✘ %s", s)))
}

func (o Output) Stop() {
	o.spinner.Stop()
}

func NewWriter() *Output {

	y, err := yacspin.New(*cfg)

	if err != nil {
		panic(err)
	}

	return &Output{
		spinner: y,
	}
}
