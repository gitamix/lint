package config

import (
	yaml "gopkg.in/yaml.v3"

	"github.com/gitamix/lint/internal/marshalling/config/branch"
	"github.com/gitamix/lint/internal/marshalling/config/commit"
)

// yamlConfig accumulates all transport config parts in one.
//
// NOTE: this cannot be moved to internal/marshalling package
// because it will call cycle import. So that is why this struct
// placed here.
type yamlConfig struct {
	Branch branch.Branch `yaml:"branch"`
	Commit commit.Commit `yaml:"commit"`
}

// Unmarshal unmarshals YAML content and returns a Config instance
// containing the configuration, or returns an error if unmarshaling fails.
func Unmarshal(bb []byte) (*Config, error) {
	var out yamlConfig
	err := yaml.Unmarshal(bb, &out)
	if err != nil {
		return nil, err
	}
	cfg := NewConfig(
		WithBranch(
			out.Branch.Config(),
		),
		WithCommit(
			out.Commit.Config(),
		),
	)
	return cfg, nil
}
