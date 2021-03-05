package parser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	type testcase struct {
		in, out string
	}
	yyErrorVerbose = true
	tests := []testcase{
		{"1", "1"},
		{"1+2", "1 + 2"},
		{"foobar", "foobar"},
	}

	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			ast, err := Parse(tc.in)
			require.NoError(t, err)
			require.Equal(t, tc.out, ast.ToString())
		})
	}
}
