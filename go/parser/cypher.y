%{

package parser

import "fmt"

var regs = make([]int, 26)
var base int

%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union {
	bytes []byte
	
	val int
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <val> expr number

// same for terminals
%token LEX_ERROR EOF ASTERISK COMMA ILLEGAL IDENT
%token RETURN WS

%token <val> DIGIT LETTER

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%start expr

%%

list	: /* empty */
	| list stat '\n'
	;

stat	:    expr
		{
			fmt.Printf( "%d\n", $1 );
		}
	|    LETTER '=' expr
		{
			regs[$1]  =  $3
		}
	;

expr	:    number
	;

number	:    DIGIT
		{
			$$ = $1;
			if $1==0 {
				base = 8
			} else {
				base = 10
			}
		}
	|    number DIGIT
		{ $$ = base * $1 + $2 }
	;

%%      /*  start  of  programs  */
