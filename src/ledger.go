package main

import (
	"os/exec"
)

// Execute ledger binary and return the output.
func runLedger(parameters ...string) (string, error) {
	// Ref:
	// https://pkg.go.dev/os/exec

	result, err := combinedOutput(parameters...)

	return result, err
}

func getCmd(parameters ...string) *exec.Cmd {
	cmd := exec.Command("ledger", parameters...)
	return cmd
}

func combinedOutput(parameters ...string) (string, error) {
	cmd := getCmd(parameters...)
	out, err := cmd.CombinedOutput()

	result := string(out)

	return result, err
}
