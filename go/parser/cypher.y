%{

package parser

var regs = make([]int, 26)
var base int

func setParseTree(yylex interface{}, result ASTNode) {
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
%type <expr> expr expr1

// same for terminals
%token LEX_ERROR EOF PLUS COMMA ILLEGAL 
%token RETURN WS

%token <bytes> DIGIT IDENT

%start start_statement

%%

start_statement: 
	expr
	{
	    setParseTree(yylex, $1)
	}

expr:
    expr1 PLUS expr1
    {
        $$ = &Add{Left: $1, Right: $3}
    }
|   expr1
    
expr1:   
    DIGIT
    {
        $$ = &LiteralInt{bytes: $1}
    }
|   IDENT
    {
        $$ = &Variable{name: string($1)}
    } 

%%
