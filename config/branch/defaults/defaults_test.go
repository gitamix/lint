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

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("ok if name is not empty", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithName("main"),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if name is numeric", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithName("123"),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok if name is not trimmed", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithName(" main "),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("error if name is empty", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithName(""),
			).
			Validate()
		assert.EqualError(t, gotErr, "empty default branch name")
	})

	t.Run("error if name is just a whitespace", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithName(" "),
			).
			Validate()
		assert.EqualError(t, gotErr, "empty default branch name")
	})
}
