// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package log

import (
	"github.com/sirupsen/logrus"
)

type Log struct {
	entry *logrus.Entry
}

func (l Log) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l Log) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l Log) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l Log) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l Log) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l Log) Panic(args ...interface{}) {
	l.entry.Panic(args...)
}

func (l Log) SetLevel(n string) {

	var level logrus.Level
	switch n {
	case "debug":
		level = logrus.DebugLevel
		break

	case "info":
		level = logrus.InfoLevel
		break

	case "error":
		level = logrus.ErrorLevel
		break

	case "fatal":
		level = logrus.FatalLevel
		break

	case "panic":
		level = logrus.PanicLevel
		break

	default:
		level = logrus.WarnLevel
	}

	logrus.SetLevel(level)
}

func New(label string) Logger {
	return &Log{
		entry: logrus.WithFields(logrus.Fields{"lebel": label}),
	}
}
