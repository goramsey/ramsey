%{

package parser

import "fmt"

var regs = make([]int, 26)
var base int

func setParseTree(yylex interface{}, result AST) {
  yylex.(*Scanner).ast = result
}

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union {
	bytes []byte
	expr Expr
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <expr> expr

// same for terminals
%token LEX_ERROR EOF ASTERISK COMMA ILLEGAL IDENT
%token RETURN WS

%token <bytes> DIGIT LETTER

%start start_statement

%%

start_statement: 
	expr
	{
	    setParseTree(yylex, $1)
	}

expr:
    DIGIT
    {
    	fmt.Println($1)
        $$ = &LiteralInt{bytes: $1}
    }

%%
