// Copyright 2020 Mohammed El Bahja. All rights reserved.
// Use of this source code is governed by a MIT license.

package actions

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/guark/guark/app/utils"
	"github.com/guark/guark/cmd/guark/stdio"
	. "github.com/guark/guark/cmd/guark/utils"
	"github.com/guark/guark/internal/embed"
	"github.com/urfave/cli/v2"
)

var (
	DevFlags = []cli.Flag{
		&cli.StringFlag{
			Name:  "pkg",
			Usage: "Set your package manager.",
			Value: "yarn",
		},
	}
)

func Dev(c *cli.Context) error {

	var (
		err      error
		out      = stdio.NewWriter()
		sig      = make(chan os.Signal)
		lock     = Path("ui", "guark.lock")
		cmd      *exec.Cmd
		cancel   context.CancelFunc
		teardown = func(c *exec.Cmd, cancel context.CancelFunc) {
			kill(c)
			cancel()
		}
	)

	if err = embed.Embed([]string{"guark.yaml"}, Path("lib", "embed.go"), "lib", ""); err != nil {
		return err
	}

	port, err := utils.GetNewPort()

	if err != nil {
		return err
	}

	out.Update("Waiting for UI dev server to start...")

	os.Remove(lock)

	cmd, cancel = serve(c.String("pkg"), port)
	defer teardown(cmd, cancel)

	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println()
		teardown(cmd, cancel)
		out.Done("Cleanup before exit.")
		os.Exit(1)
	}()

	out.Stop()

	for {

		time.Sleep(500 * time.Millisecond)

		if utils.IsPortOpen(fmt.Sprintf("127.0.0.1:%s", port), 5) && utils.IsFile(lock) {
			break
		}
	}

	out.Done("UI server started successfully.")

	return start(port, out)
}

func serve(pkg string, port string) (*exec.Cmd, context.CancelFunc) {

	var (
		ctx, cancel = context.WithCancel(context.Background())
	)

	cmd := exec.CommandContext(ctx, pkg, "run", "serve", "--host", "127.0.0.1", "--port", port)
	cmd.Dir = Path("ui")
	cmd.Stderr = os.Stderr
	cmd.Start()

	return cmd, cancel
}

func start(port string, out *stdio.Output) error {

	cmd := exec.Command("go", "run", "-tags", "dev", "app.go")
	cmd.Env = append(os.Environ(), fmt.Sprintf("GUARK_DEBUG_PORT=%s", port))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	out.Done("Starting guark dev app...")
	return cmd.Run()
}

func kill(cmd *exec.Cmd) {

	if cmd == nil {
		return
	}

	cmd.Process.Signal(syscall.SIGTERM)
}
