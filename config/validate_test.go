package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/config"
	"github.com/gitamix/lint/config/branch"
	"github.com/gitamix/lint/config/branch/defaults"
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
	"github.com/gitamix/lint/issue"
)

func TestConfig_Validate(t *testing.T) {
	t.Parallel()

	t.Run("err for empty config", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.NewConfig().Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid default branch: empty default branch name",
		)
	})

	t.Run("err for empty default branch with valid commit", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithCommit(
					commit.NewConfig(
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Critical,
									"feat",
									"fix",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid default branch: empty default branch name",
		)
	})

	t.Run("err for blank default branch name", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("   "),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid default branch: empty default branch name",
		)
	})

	t.Run("ok with default branch only", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with default branch and empty commit", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("ok with fully valid branch and commit", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
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
													value.NewRange(issue.Warning, 10, 72),
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
										body.WithMandate(
											mandate.NewConfig(
												mandate.WithTypes(
													value.NewStrings(
														issue.Critical,
														"fix",
														"chore",
													),
												),
											),
										),
										body.WithLength(
											value.NewRange(issue.Info, 10, 255),
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
								),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("err for default branch even if commit is invalid", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithCommit(
					commit.NewConfig(
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Type(-1),
									"feat",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid default branch: empty default branch name",
		)
	})

	t.Run("err for branch name before commit scope", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
						branch.WithName(
							name.NewConfig(
								value.NewString(
									issue.Type(-1),
									`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
								),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithScope(
							scope.NewConfig(
								value.NewString(
									issue.Type(-1),
									`^[a-z]+$`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid name: invalid config: unknown issue level",
		)
	})

	t.Run("err for branch task before commit message", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
						branch.WithTask(
							task.NewConfig(
								task.WithID(
									id.NewConfig(
										value.NewString(
											issue.Type(-1),
											`(TASK|PROJ|BUG)-[0-9]+`,
										),
									),
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
													value.NewRange(issue.Warning, 100, 1),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid task: invalid config: unknown issue level",
		)
	})

	t.Run("err for commit types when branch is valid", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
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
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Type(-1),
									"feat",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid types: invalid config: unknown issue level",
		)
	})

	t.Run("err for commit message before commit scope", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
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
													value.NewRange(issue.Warning, 100, 1),
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
									issue.Type(-1),
									`^[a-z]+$`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid message: invalid subject: invalid description: invalid config: min is greater than max",
		)
	})

	t.Run("err for commit scope before commit types", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithScope(
							scope.NewConfig(
								value.NewString(
									issue.Type(-1),
									`^[a-z]+$`,
								),
							),
						),
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Type(-1),
									"feat",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid scope: invalid config: unknown issue level",
		)
	})

	t.Run("err for unknown branch name level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
						branch.WithName(
							name.NewConfig(
								value.NewString(
									issue.Type(-1),
									`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid name: invalid config: unknown issue level",
		)
	})

	t.Run("err for unknown branch task level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
						branch.WithTask(
							task.NewConfig(
								task.WithID(
									id.NewConfig(
										value.NewString(
											issue.Type(-1),
											`(TASK|PROJ|BUG)-[0-9]+`,
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid task: invalid config: unknown issue level",
		)
	})

	t.Run("err for empty default with valid name and task", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithName(
							name.NewConfig(
								value.NewString(
									issue.Warning,
									`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
								),
							),
						),
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
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid default branch: empty default branch name",
		)
	})

	t.Run("err for commit description length with min greater than max", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
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
													value.NewRange(issue.Warning, 100, 1),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid message: invalid subject: invalid description: invalid config: min is greater than max",
		)
	})

	t.Run("err for commit subject task with unknown level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
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
										subject.WithTask(
											task.NewConfig(
												task.WithID(
													id.NewConfig(
														value.NewString(
															issue.Type(-1),
															`(TASK|PROJ|BUG)-[0-9]+`,
														),
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid message: invalid subject: invalid task: invalid config: unknown issue level",
		)
	})

	t.Run("err for commit subject before commit body", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
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
													value.NewRange(issue.Warning, 100, 1),
												),
											),
										),
									),
								),
								message.WithBody(
									body.NewConfig(
										body.WithLength(
											value.NewRange(issue.Warning, -1, 100),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid message: invalid subject: invalid description: invalid config: min is greater than max",
		)
	})

	t.Run("err for commit body length with negative min", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithMessage(
							message.NewConfig(
								message.WithBody(
									body.NewConfig(
										body.WithLength(
											value.NewRange(issue.Warning, -1, 100),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid message: invalid body: invalid length: invalid config: min is negative",
		)
	})

	t.Run("err for commit body length with negative max", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithMessage(
							message.NewConfig(
								message.WithBody(
									body.NewConfig(
										body.WithLength(
											value.NewRange(issue.Warning, 0, -5),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid message: invalid body: invalid length: invalid config: max is negative",
		)
	})

	t.Run("err for commit body mandate with unknown level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithMessage(
							message.NewConfig(
								message.WithBody(
									body.NewConfig(
										body.WithMandate(
											mandate.NewConfig(
												mandate.WithTypes(
													value.NewStrings(
														issue.Type(-1),
														"fix",
													),
												),
											),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid message: invalid body: invalid mandate: invalid config: unknown issue level",
		)
	})

	t.Run("err for commit scope with unknown level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithScope(
							scope.NewConfig(
								value.NewString(
									issue.Type(-1),
									`^[a-z]+$`,
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid scope: invalid config: unknown issue level",
		)
	})

	t.Run("err for commit types with unknown level", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Type(-1),
									"feat",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid types: invalid config: unknown issue level",
		)
	})

	t.Run("unspecified levels do not break valid branch", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
						branch.WithName(
							name.NewConfig(
								value.NewString(
									issue.Type(0),
									`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
								),
							),
						),
						branch.WithTask(
							task.NewConfig(
								task.WithID(
									id.NewConfig(
										value.NewString(
											issue.Type(0),
											`(TASK|PROJ|BUG)-[0-9]+`,
										),
									),
								),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Type(0),
									"feat",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("empty scope pattern does not break valid types", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithScope(
							scope.NewConfig(
								value.NewString(
									issue.Warning,
									"",
								),
							),
						),
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Critical,
									"feat",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("zero body length does not break valid description length", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
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
													value.NewRange(issue.Warning, 1, 72),
												),
											),
										),
									),
								),
								message.WithBody(
									body.NewConfig(
										body.WithLength(
											value.NewRange(issue.Info, 0, 0),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("equal bounds are valid while other range is invalid", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
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
													value.NewRange(issue.Warning, 10, 10),
												),
											),
										),
									),
								),
								message.WithBody(
									body.NewConfig(
										body.WithLength(
											value.NewRange(issue.Warning, 100, 1),
										),
									),
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid message: invalid body: invalid length: invalid config: min is greater than max",
		)
	})

	t.Run("invalid branch name is reported even if commit is valid", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
						branch.WithName(
							name.NewConfig(
								value.NewString(
									issue.Type(-1),
									`^(feature|bugfix|hotfix)/[A-Z]+-\d+`,
								),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Critical,
									"feat",
									"fix",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid name: invalid config: unknown issue level",
		)
	})

	t.Run("last WithBranch wins with valid branch", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(),
				),
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("last WithBranch wins with invalid branch", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithBranch(
					branch.NewConfig(),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid branch: invalid default branch: empty default branch name",
		)
	})

	t.Run("last WithCommit wins with valid commit", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Type(-1),
									"feat",
								),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(),
				),
			).
			Validate()
		assert.NoError(t, gotErr)
	})

	t.Run("last WithCommit wins with invalid commit", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
				impl.WithCommit(
					commit.NewConfig(),
				),
				impl.WithCommit(
					commit.NewConfig(
						commit.WithTypes(
							types.NewConfig(
								value.NewStrings(
									issue.Type(-1),
									"feat",
								),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid types: invalid config: unknown issue level",
		)
	})

	t.Run("WithBranch does not reset invalid commit", func(t *testing.T) {
		t.Parallel()
		gotErr := impl.
			NewConfig(
				impl.WithCommit(
					commit.NewConfig(
						commit.WithScope(
							scope.NewConfig(
								value.NewString(
									issue.Type(-1),
									`^[a-z]+$`,
								),
							),
						),
					),
				),
				impl.WithBranch(
					branch.NewConfig(
						branch.WithDefault(
							defaults.NewConfig(
								defaults.WithName("master"),
							),
						),
					),
				),
			).
			Validate()
		assert.EqualError(
			t,
			gotErr,
			"invalid commit: invalid scope: invalid config: unknown issue level",
		)
	})
}
