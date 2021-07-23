package manhat_test

import (
	"bytes"
	"testing"

	"github.com/qba73/manhat"
)

func TestCliAppVersion(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Print version",
			args: []string{"-version"},
			want: "Version: 0.1.0\nGitRef: c84cf\nBuild Time: 2021-07-19-09-20-42Z\n",
		},
		{
			name: "Calculate distance",
			args: []string{"-location", "12"},
			want: "3\n",
		},
		{
			name: "Pass invalid value",
			args: []string{"-location", "0"},
			want: "",
		},
		{
			name: "Pass invalid value",
			args: []string{"-location", "-1"},
			want: "",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := &bytes.Buffer{}

			if err := manhat.Cli(tc.args, out); err != nil {
				t.Fatal(err)
			}
			got := out.String()

			if tc.want != got {
				t.Errorf("want %q, got %q", tc.want, got)
			}

		})
	}
}
