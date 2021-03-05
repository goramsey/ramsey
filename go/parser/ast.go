package parser

import "fmt"

type (
	AST interface {
		ToString() string
	}

	Expr interface {
		AST
		iExpr()
	}

	LiteralInt struct {
		bytes []byte
		val int
	}
)

func (l *LiteralInt) ToString() string {
	return fmt.Sprintf("%d", l.val)
}

func (l *LiteralInt) iExpr() {}

var _ Expr = (*LiteralInt)(nil)
