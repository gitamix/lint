package issue

// Issue represents a linting issue with a message and a type.
//
// It is used to encapsulate the details of an issue found during the linting process.
// The Issue struct contains a message that describes the issue and a type that indicates its severity.
type Issue struct {
	// msg is the description of the issue.
	msg string
	// typ indicates the severity of the issue (critical, warning, info).
	typ Type
}

// NewIssue creates a new Issue instance with the specified message and type.
func NewIssue(typ Type, msg string) Issue {
	return Issue{
		msg: msg,
		typ: typ,
	}
}

// NewCritical creates a new Issue with type Critical.
//
// It is a convenience function to create an critical issue.
// This function is used to create issues that indicate critical problems
// that need to be solved immediately.
func NewCritical(msg string) Issue {
	return NewIssue(Critical, msg)
}

// NewWarning creates a new Issue with type Warning.
//
// It is a convenience function to create a warning issue.
// This function is used to create issues that indicate potential problems
// that should be reviewed but are not critical.
func NewWarning(msg string) Issue {
	return NewIssue(Warning, msg)
}

// NewInfo creates a new Issue with type Info.
//
// It is a convenience function to create an informational issue.
// This function is used to create issues that provide additional information
// that does not require immediate action.
func NewInfo(msg string) Issue {
	return NewIssue(Info, msg)
}

// Message returns the message of the issue.
//
// It provides a human-readable description of the issue.
// This function is used to retrieve the message associated with the issue,
// which can be displayed to the user or logged for debugging purposes.
func (i Issue) Message() string {
	return i.msg
}

// Type returns the type of the issue.
//
// This function is used to categorize the issue based on its severity,
// allowing the linting process to handle different types of issues appropriately.
func (i Issue) Type() Type {
	return i.typ
}
