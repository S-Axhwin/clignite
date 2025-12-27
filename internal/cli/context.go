package cli

import (
	"bytes"
	"os"
	"os/exec"
)

func currentContext() (string, error) {
	return os.Getwd()
}

// getRepoRoot returns the top-level directory for the current repository.
// If git is present, use git root. Otherwise fallback to current directory.

func getRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	cmd.Dir, _ = os.Getwd()
	out, err := cmd.Output()
	if err == nil {
		return string(bytes.TrimSpace(out)), nil // <--- trim newline/whitespace
	}
	// fallback to current directory if not in git repo
	cwd, err := os.Getwd()
	return cwd, err
}
