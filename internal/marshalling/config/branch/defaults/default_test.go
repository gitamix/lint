package defaults_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/lint/config/branch/defaults"
	impl "github.com/gitamix/lint/internal/marshalling/config/branch/defaults"
)

func TestDefault_Config(t *testing.T) {
	t.Parallel()

	t.Run("returns zero config when empty", func(t *testing.T) {
		t.Parallel()
		var want defaults.Config
		var cfg impl.Default
		got := cfg.Config()
		assert.Equal(t, want, got)
	})

	t.Run("converts name into defaults config", func(t *testing.T) {
		t.Parallel()
		want := defaults.NewConfig(
			defaults.WithName("master"),
		)
		cfg := impl.Default{
			Name: "master",
		}
		got := cfg.Config()
		assert.Equal(t, want, got)
	})
}
