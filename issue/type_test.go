package issue_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/lint/issue"
)

func TestType_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		t    impl.Type
		want string
	}{
		{
			name: "Critical",
			t:    impl.Critical,
			want: "critical",
		},
		{
			name: "Warning",
			t:    impl.Warning,
			want: "warning",
		},
		{
			name: "Info",
			t:    impl.Info,
			want: "info",
		},
		{
			name: "custom 255",
			t:    impl.Type(255),
			want: "",
		},
		{
			name: "custom zero",
			t:    impl.Type(0),
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(
				t,
				tt.want,
				tt.t.String(),
			)
		})
	}
}
