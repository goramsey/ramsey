package parser

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	scanner := NewScanner(strings.NewReader("1"))
	token, lit := scanner.Scan()
	assert.EqualValues(t, token, DIGIT)
	assert.Equal(t, lit, "1")
}