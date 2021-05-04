package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed custom/main.ps1
var script string

func runScript(outputEnabled bool) {
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
