package parser

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	type testcase struct {
		query  string
		expect AST
	}
	yyDebug = 999
	yyErrorVerbose = true
	tests := []testcase{
		{query: "1", expect: &LiteralInt{bytes: []byte("1")}},
	}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			ast, err := Parse(tc.query)
			require.NoError(t, err)
			require.NotNil(t, ast)
		})
	}
}
