package value

import "github.com/gitamix/lint/issue"

// Issue is the transport representation
// of the issue severity level used by the gitamix linter
// to report problems.
type Issue struct {
	// Level stores the severity level
	// of the issue as a string read from the config file.
	Level string `yaml:"level,omitempty"`
}

// Empty reports whether the Issue is empty,
// i.e. no severity level is set.
func (i Issue) Empty() bool {
	return i.Level == ""
}

// Config converts the Issue into the domain value.
func (i Issue) Config() issue.Type {
	return issue.Parse(i.Level)
}
