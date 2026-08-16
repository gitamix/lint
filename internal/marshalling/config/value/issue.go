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
//
// When empty, the zero issue type is returned.
// Otherwise the level string is parsed,
// defaulting to warning when it does not match a known level.
func (i Issue) Config() issue.Type {
	if i.Empty() {
		var typ issue.Type
		return typ
	}
	return issue.ParseOr(i.Level, issue.Warning)
}
