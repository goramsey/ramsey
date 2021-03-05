package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	type r struct {
		tok Token
		lit string
	}
	tests := []struct {
		query    string
		expected []r
	}{{
		query: "1",
		expected: []r{{
			tok: DIGIT,
			lit: "1",
		}},
	}, {
		query: "1+2",
		expected: []r{{
			tok: DIGIT,
			lit: "1",
		}, {
			tok: PLUS,
			lit: "+",
		}, {
			tok: DIGIT,
			lit: "2",
		}},
	}, {
		query: "barfoo",
		expected: []r{{
			tok: IDENT,
			lit: "barfoo",
		}},
	}}

	for _, tc := range tests {
		t.Run(tc.query, func(t *testing.T) {
			scanner := NewScanner(strings.NewReader(tc.query))
			var res []r
			for {
				token, lit := scanner.Scan()
				if token == 0 {
					break
				}
				res = append(res, r{
					tok: token,
					lit: lit,
				})
			}
			assert.Equal(t, tc.expected, res)
		})
	}
}
