package main

import (
	"fmt"
	"github.com/victorsalaun/poc-go-releaser/pkg/cmd/factory"
	"github.com/victorsalaun/poc-go-releaser/pkg/cmd/root"
	"os"
)

type exitCode int

const (
	exitOK    exitCode = 0
	exitError exitCode = 1
)

func main() {
	code := mainRun()
	os.Exit(int(code))
}

func mainRun() exitCode {
	cmdFactory := factory.New()

	rootCmd := root.NewCmdRoot(cmdFactory)

	if _, err := rootCmd.ExecuteC(); err != nil {
		_, _ = fmt.Fprintln(cmdFactory.IOStreams.ErrOut, err.Error())
		return exitError
	}

	return exitOK
}
