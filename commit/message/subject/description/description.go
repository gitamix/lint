package description

import (
	"github.com/gitamix/types/commit"

	config "github.com/gitamix/lint/config/commit/message/subject/description"
)

// Description represents a linter that validates the length
// of the commit message subject description against the configured interval.
type Description struct {
	// desc is the parsed commit message subject description
	// inspected for its length.
	desc commit.Description

	// cfg is the configuration that defines
	// the length interval used to validate the subject description.
	cfg config.Config
}

// NewDescription creates a new Description linter
// with the provided commit message subject description
// and lint configuration that defines the length interval
// used to validate the subject description.
func NewDescription(
	desc commit.Description,
	cfg config.Config,
) Description {
	return Description{
		desc: desc,
		cfg:  cfg,
	}
}
