package parser

import "fmt"

type (
	ASTNode interface {
		ToString() string
	}

	Expr interface {
		ASTNode
		iExpr()
	}

	Add struct {
		Left, Right Expr
	}

	LiteralInt struct {
		bytes []byte
	}

	Variable struct {
		name string
	}
)

func (l *LiteralInt) ToString() string {
	return fmt.Sprintf("%s", string(l.bytes))
}

func (a *Add) ToString() string {
	return fmt.Sprintf("%s + %s", a.Left.ToString(), a.Right.ToString())
}
func (a *Variable) ToString() string {
	return a.name
}

func (*LiteralInt) iExpr() {}
func (*Add) iExpr()        {}
func (*Variable) iExpr()   {}

var _ Expr = (*LiteralInt)(nil)
