package message

import (
	"github.com/gitamix/types/commit"

	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
)

// Message represents a linter that validates a complete commit message
// by aggregating issues from its task, subject, and body sub-linters.
type Message struct {
	// msg is the parsed commit message inspected by the linter.
	msg commit.Message

	// cfg is the message-level configuration
	// that provides the subject and body sub-configurations
	// used by the corresponding sub-linters.
	cfg message.Config

	// typcfg is the configuration that defines
	// the allowed commit types for the type sub-linter.
	typcfg types.Config

	// scpcfg is the configuration that defines
	// the pattern used to validate the commit scope
	// by the scope sub-linter.
	scpcfg scope.Config
}

// NewMessage creates a new Message linter
// with the provided commit message and lint configurations.
//
//   - msg: the parsed commit message to validate.
//   - cfg: the message-level configuration providing subject and body sub-configs.
//   - typcfg: the configuration defining the allowed commit types.
//   - scpcfg: the configuration defining the pattern used to validate the commit scope.
func NewMessage(
	msg commit.Message,
	cfg message.Config,
	typcfg types.Config,
	scpcfg scope.Config,
) Message {
	return Message{
		msg:    msg,
		cfg:    cfg,
		typcfg: typcfg,
		scpcfg: scpcfg,
	}
}
