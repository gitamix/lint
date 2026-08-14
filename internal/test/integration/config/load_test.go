//go:build integration
// +build integration

package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	impl "github.com/gitamix/lint/config"
	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/name"
	"github.com/gitamix/lint/config/task"
	"github.com/gitamix/lint/config/task/id"
	"github.com/gitamix/lint/config/value"
	"github.com/gitamix/lint/internal/test/testdata"
	"github.com/gitamix/lint/issue"
)

func TestLoad(t *testing.T) {
	t.Parallel()
	t.Run("load config file", func(t *testing.T) {
		t.Parallel()
		fpath := testdata.PathToFile("lint.yml")
		require.FileExistsf(
			t,
			fpath,
			"Test data configuration file does not exist",
		)
		got, gotErr := impl.Load(fpath)
		assert.NoError(t, gotErr)
		assert.Equal(
			t,
			impl.NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithTask(
							task.NewConfig(
								task.WithID(
									id.NewConfig(
										value.NewString(
											issue.Info,
											`(TASK|PROJ|BUG)-[0-9]+`,
										),
									),
								),
							),
						),
						branch.WithName(
							name.NewConfig(
								value.NewString(
									issue.Warning,
									`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
								),
							),
						),
					),
				),
			),
			got,
		)
	})

	t.Run("file not exists", func(t *testing.T) {
		t.Parallel()
		fpath := "not_exists.yml"
		require.NoFileExists(t, fpath)
		got, gotErr := impl.Load(fpath)
		assert.ErrorIs(t, gotErr, os.ErrNotExist)
		assert.Nil(t, got)
	})

	t.Run("not yaml", func(t *testing.T) {
		t.Parallel()
		fpath := testdata.PathToFile("not_yaml.txt")
		require.FileExistsf(
			t,
			fpath,
			"The required file that is supposed to have incorrect format does not exist",
		)
		got, gotErr := impl.Load(fpath)
		assert.Error(t, gotErr)
		assert.Nil(t, got)
	})
}
