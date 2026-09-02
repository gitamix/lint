package errs

import (
	"errors"
)

var (
	// ErrGitFailed is returned when a git command execution fails.
	ErrGitFailed = errors.New("git command failed")
)
