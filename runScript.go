package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gobuffalo/packr"
)

var customFolder = packr.NewBox("./custom")

func runScript(outputEnabled bool) {
	// load script
	script, err := customFolder.FindString("main.ps1")
	fatalErr(err)
	// tabs trigger tab-completion. use spaces instead.
	// make sure we have a trailing newline so last line will execute.
	script = strings.ReplaceAll(script, "\t", " ") + "\n"

	// launch powershell
	cmd := exec.Command("powershell")
	cmd.Stdin = strings.NewReader(script)
	if outputEnabled {
		// connect output to real output
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	cmd.Run()

	// script is over
	if outputEnabled {
		// powershell leaves us at a prompt (no newline)
		fmt.Println()
	}
	os.Exit(0)
}
