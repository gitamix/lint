package issue

import "slices"

const (
	// Critical indicates a critical issue that needs to be solved immediately.
	//
	// It is used to represent issues that prevent the application from functioning correctly.
	// This type of issue should be addressed as a priority.
	// It is typically used for issues that indicate a failure in the linting process
	// or a violation of essential rules that must be corrected.
	Critical Type = 1

	// Warning indicates a potential issue that should be reviewed but is not critical.
	//
	// It is used to represent issues that may not prevent the application from functioning,
	// but could lead to problems if not addressed.
	// This type of issue is typically used for linting rules that are not strictly enforced
	// but are recommended for best practices.
	Warning Type = 2

	// Info indicates informational messages that do not require action.
	//
	// It is used to provide additional context or information about the linting process.
	// This type of issue is typically used for messages that provide insights
	// or details about the linting process, such as configuration settings or
	// general information that may be useful for the user but does not indicate a problem.
	Info Type = 3
)

// Type represents the severity of an issue in the linting process
// and is used to categorize issues found during linting.
//
// It indicates the severity of the issue (critical, warning, info),
// allowing the linting process to handle different types of issues appropriately.
type Type uint8

// String returns a string representation of the Type.
func (t Type) String() string {
	switch t {
	case Critical:
		return "critical"
	case Warning:
		return "warning"
	case Info:
		return "info"
	default:
		return ""
	}
}

// Unspecified defines whether the type is unspecified.
func (t Type) Unspecified() bool {
	return t == Type(0)
}

// In defines whether the type exists in provided list of types.
func (t Type) In(typs ...Type) bool {
	return slices.Contains(typs, t)
}

// Unknown defines whether the type is unknown.
func (t Type) Unknown() bool {
	return !t.In(Critical, Warning, Info)
}

// Parse parses the provided string into an issue type
// or returns zero type if not parsed.
func Parse(s string) Type {
	return ParseOr(s, Type(0))
}

// ParseOr parses the provided string into an issue Type,
// or returns the provided default if the value does not match any known type.
func ParseOr(s string, def Type) Type {
	switch s {
	case "critical":
		return Critical
	case "warning":
		return Warning
	case "info":
		return Info
	default:
		return def
	}
}
