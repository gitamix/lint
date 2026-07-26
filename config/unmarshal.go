package config

import (
	yaml "gopkg.in/yaml.v3"

	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/ticket"
	"github.com/gitamix/lint/config/ticket/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/issue"
)

type (
	yamlConfig struct {
		Branch struct {
			Ticket struct {
				Issue   yamlIssue `yaml:"issue"`
				Pattern string    `yaml:"pattern"`
			} `yaml:"ticket"`
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
				branch.WithTicket(
					ticket.NewConfig(
						ticket.WithID(
							id.NewConfig(
								value.NewString(
									issue.ParseOr(
										out.Branch.Ticket.Issue.Level,
										issue.Warning,
									),
									out.Branch.Ticket.Pattern,
								),
							),
						),
					),
				),
				branch.WithName(
					name.NewConfig(
						value.NewString(
							issue.ParseOr(
								out.Branch.Name.Issue.Level,
								issue.Warning,
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
