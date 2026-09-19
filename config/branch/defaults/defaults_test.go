package defaults_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config/branch/defaults"
)

func TestConfig_Name(t *testing.T) {
	t.Parallel()

	t.Run("with name", func(t *testing.T) {
		t.Parallel()
		want := "master"
		got := impl.
			NewConfig(impl.WithName("master")).
			Name()
		assert.Equal(t, want, got)
	})

	t.Run("with empty name", func(t *testing.T) {
		t.Parallel()
		var want string
		got := impl.
			NewConfig(impl.WithName("")).
			Name()
		assert.Equal(t, want, got)
	})

	t.Run("without options", func(t *testing.T) {
		t.Parallel()
		var want string
		got := impl.
			NewConfig().
			Name()
		assert.Equal(t, want, got)
	})

	t.Run("default config", func(t *testing.T) {
		t.Parallel()
		var want string
		var cfg impl.Config
		var got = cfg.Name()
		assert.Equal(t, want, got)
	})
}
