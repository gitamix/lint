package issue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/issue"
)

func TestIssue_Message(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		i    impl.Issue
		want string
	}{
		{
			name: "message with critical constructor",
			i:    impl.NewCritical("test message"),
			want: "test message",
		},
		{
			name: "message with warning constructor",
			i:    impl.NewWarning("test message"),
			want: "test message",
		},
		{
			name: "message with info constructor",
			i:    impl.NewInfo("test message"),
			want: "test message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.i.Message(),
			)
		})
	}
}

func TestIssue_Type(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		i    impl.Issue
		want impl.Type
	}{
		{
			name: "critical issue with its contstructor",
			i:    impl.NewCritical("test message"),
			want: impl.Critical,
		},
		{
			name: "warning issue with its contstructor",
			i:    impl.NewWarning("test message"),
			want: impl.Warning,
		},
		{
			name: "info issue with its contstructor",
			i:    impl.NewInfo("test message"),
			want: impl.Info,
		},
		{
			name: "info issue with issue contstructor",
			i:    impl.NewIssue(impl.Info, "test message"),
			want: impl.Info,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.i.Type(),
			)
		})
	}
}
