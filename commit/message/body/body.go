package body

import (
	"github.com/gitamix/types/commit"

	"github.com/gitamix/lint/config/commit/message/body"
)

// Body represents a linter that validates
// the commit message body against the configured length interval.
type Body struct {
	// body is the parsed commit message body inspected by the linter.
	body commit.Body

	// cfg is the configuration that defines
	// the length interval used to validate the commit message body.
	cfg body.Config
}

// NewBody creates a new Body linter
// with the provided commit message body and lint configuration.
func NewBody(
	body commit.Body,
	cfg body.Config,
) Body {
	return Body{
		body: body,
		cfg:  cfg,
	}
}
