package app

import (
	"fmt"
	"io"

	"github.com/leo3nava/leo3nava_ai/tools-go/leo3nava-ai/internal/phasestatus"
)

// Run executes the phase-status-validator logic.
// It parses args to extract a change root path, resolves the phase status,
// and writes the frozen JSON wire document to stdout.
// Errors are written to stderr.
func Run(stdout, stderr io.Writer, args []string) error {
	if len(args) < 1 || args[0] == "" {
		fmt.Fprintln(stderr, "error: change root path is required")
		return fmt.Errorf("usage: phase-status-validator <change-root-path>")
	}

	changeRoot := args[0]

	status, err := phasestatus.Resolve(phasestatus.Options{ChangeRoot: changeRoot})
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return err
	}

	jsonBytes, err := phasestatus.ProjectWire(status)
	if err != nil {
		fmt.Fprintf(stderr, "error: %v\n", err)
		return err
	}

	_, err = stdout.Write(jsonBytes)
	if err != nil {
		fmt.Fprintf(stderr, "error: writing output: %v\n", err)
		return err
	}

	// Trailing newline for clean terminal output.
	fmt.Fprintln(stdout)

	return nil
}
