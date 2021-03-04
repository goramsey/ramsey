package parser

import (
	"io"
	"testing"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestHelloWorld(c *C) {
	c.Assert(42, Equals, "42")
	c.Assert(io.ErrClosedPipe, ErrorMatches, "io: .*on closed pipe")
	c.Check(42, Equals, 42)
}


//func TestParse(t *testing.T) {
//	type testcase struct {
//		query  string
//		expect AST
//	}
//
//	tests := []testcase{
//		{query: "1", expect: &LiteralInt{1}},
//	}
//	for _, tc := range tests {
//		ast, err := Parse(tc.query)
//		require.N
//	}
//}