package config

import (
	yaml "gopkg.in/yaml.v3"

	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

type (
	yamlConfig struct {
		Branch struct {
			Task struct {
				Issue   yamlIssue `yaml:"issue"`
				Pattern string    `yaml:"pattern"`
			} `yaml:"task"`
			Name struct {
				Issue   yamlIssue `yaml:"issue"`
				Pattern string    `yaml:"pattern"`
			} `yaml:"name"`
		} `yaml:"branch"`
	}

	yamlIssue struct {
		Level string `yaml:"level"`
	}
)

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
			branch.NewConfig(
				branch.WithTask(
					task.NewConfig(
						task.WithID(
							id.NewConfig(
								value.NewString(
									issue.Parse(
										out.Branch.Task.Issue.Level,
									),
									out.Branch.Task.Pattern,
								),
							),
						),
					),
				),
				branch.WithName(
					name.NewConfig(
						value.NewString(
							issue.Parse(
								out.Branch.Name.Issue.Level,
							),
							out.Branch.Name.Pattern,
						),
					),
				),
			),
		),
	)
	return cfg, nil
}
