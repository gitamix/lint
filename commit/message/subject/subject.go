package subject

import (
	"github.com/gitamix/types/commit"

	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
)

// Subject represents a linter that validates a commit message subject
// by aggregating issues from its task, type, scope, and description sub-linters.
type Subject struct {
	// subj is the parsed commit message subject
	// inspected by the sub-linters.
	subj commit.Subject

	// cfg is the subject-level configuration
	// that provides the task and description sub-configurations
	// used by the corresponding sub-linters.
	cfg subject.Config

	// typcfg is the configuration that defines
	// the allowed commit types for the type sub-linter.
	typcfg types.Config

	// scpcfg is the configuration that defines
	// the pattern used to validate the commit scope
	// by the scope sub-linter.
	scpcfg scope.Config
}

// NewSubject creates a new Subject linter
// with the provided commit message subject and lint configurations.
//
//   - subj: the parsed commit message subject to validate.
//   - cfg: the subject-level configuration providing task and description sub-configs.
//   - typcfg: the configuration defining the allowed commit types.
//   - scpcfg: the configuration defining the pattern used to validate the commit scope.
func NewSubject(
	subj commit.Subject,
	cfg subject.Config,
	typcfg types.Config,
	scpcfg scope.Config,
) Subject {
	return Subject{
		subj:   subj,
		cfg:    cfg,
		typcfg: typcfg,
		scpcfg: scpcfg,
	}
}
