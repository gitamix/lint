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
	"github.com/gitamix/lint/config/commit"
	"github.com/gitamix/lint/config/commit/message"
	"github.com/gitamix/lint/config/commit/message/body"
	"github.com/gitamix/lint/config/commit/message/body/mandate"
	"github.com/gitamix/lint/config/commit/message/subject"
	"github.com/gitamix/lint/config/commit/message/subject/description"
	"github.com/gitamix/lint/config/commit/scope"
	"github.com/gitamix/lint/config/commit/types"
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
		want := impl.NewConfig(
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
			impl.WithCommit(
				commit.NewConfig(
					commit.WithMessage(
						message.NewConfig(
							message.WithSubject(
								subject.NewConfig(
									subject.WithDescription(
										description.NewConfig(
											description.WithLength(
												value.NewRange(
													issue.Warning,
													10,
													72,
												),
											),
										),
									),
									subject.WithTask(
										task.NewConfig(
											task.WithID(
												id.NewConfig(
													value.NewString(
														issue.Critical,
														`(TASK|PROJ|BUG)-[0-9]+`,
													),
												),
											),
										),
									),
								),
							),
							message.WithBody(
								body.NewConfig(
									body.WithLength(
										value.NewRange(
											issue.Info,
											10,
											255,
										),
									),
									body.WithMandate(
										mandate.NewConfig(
											mandate.WithTypes(
												value.NewStrings(
													issue.Critical,
													"fix",
													"chore",
													"refactor",
													"perf",
												),
											),
										),
									),
								),
							),
						),
					),
					commit.WithScope(
						scope.NewConfig(
							value.NewString(
								issue.Warning,
								`^[A-Za-z _-]+$`,
							),
						),
					),
					commit.WithTypes(
						types.NewConfig(
							value.NewStrings(
								issue.Critical,
								"feat",
								"fix",
								"chore",
								"refactor",
								"perf",
								"test",
								"docs",
							),
						),
					),
				),
			),
		)
		assert.NoError(t, gotErr)
		assert.Equal(t, want, got)
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
