// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package utils

import (
	"net"
	"strings"
	"time"
)

// Check is a port open.
func IsPortOpen(addr string, timeout int) bool {

	_, e := net.DialTimeout("tcp", addr, time.Duration(timeout)*time.Second)

	return e == nil
}

// Get unused port number
func GetNewPort() (string, error) {

	ln, err := net.Listen("tcp", "127.0.0.1:0")

	if err != nil {
		return "", err
	}

	defer ln.Close()

	parts := strings.Split(ln.Addr().String(), ":")

	return parts[len(parts)-1], nil
}
