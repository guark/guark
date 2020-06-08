// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package stdio

import (
	"fmt"

	"github.com/guark/guark/cmd/guark/stdio/uilive"
	colors "github.com/logrusorgru/aurora"
)

type Output struct {
	Writer *uilive.Writer
}

func (o Output) Update(s string, icon string) {

	if icon == "" {
		icon = "●"
	}

	fmt.Fprintf(o.Writer, "%s %s", colors.Bold(colors.Cyan(icon)), colors.Bold(colors.Blue(s)))
}

func (o Output) End(s string, icon string) {

	if icon == "" {
		icon = "✔"
	}

	fmt.Println(fmt.Sprintf("%s %s", colors.Green(icon), colors.Cyan(s)))
	o.Writer.Stop()
	o.Writer.Start()
}

func NewWriter() *Output {

	w := uilive.New()
	w.Start()

	return &Output{
		Writer: w,
	}
}
