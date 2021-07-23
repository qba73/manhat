package manhat_test

import (
	"bytes"
	"testing"

	"github.com/qba73/manhat"
)

func TestCliCorrectInput(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Print version",
			args: []string{"-version"},
			want: "Version: \nGitRef: \nBuild Time: \n",
		},
		{
			name: "Calculate distance",
			args: []string{"-location", "12"},
			want: "3\n",
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

func TestCliInvalidInput(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Pass invalid value",
			args: []string{"-location", "0"},
		},
		{
			name: "Pass invalid value",
			args: []string{"-location", "-1"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			out := &bytes.Buffer{}

			if err := manhat.Cli(tc.args, out); err == nil {
				t.Fatal(err)
			}
		})
	}
}
