package lexer

import (
	"panda/token"
	"strings"
	"testing"
)

func TestLexer(t *testing.T) {
	input := `var five = 5;
	var ten = 10;
	var add = function(x, y) {
		x + y;
	};
	
	var result = add(five, ten);
	!-/*5;
	5 < 10 > 5;
	
	if (5 < 0) {
		return true;
	} else {
		return false;
	}
	5 == 5;
	5 != 6;
	"foobar";
	"foo bar";
	[];
	cccc.call
	{ "foo"."bar" }
	[1:3]
	5 % 4
	import tests
	x && y
	x || y
	struct
	for
	100.001!=0.123
	///.. aaaaa 
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.NUMBER, "5"},
		{token.SEMI, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.NUMBER, "10"},
		{token.SEMI, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNC, "function"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMI, ";"},
		{token.RBRACE, "}"},
		{token.SEMI, ";"},
		{token.VAR, "var"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMI, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.DIV, "/"},
		{token.MUL, "*"},
		{token.NUMBER, "5"},
		{token.SEMI, ";"},
		{token.NUMBER, "5"},
		{token.LESS, "<"},
		{token.NUMBER, "10"},
		{token.GREATER, ">"},
		{token.NUMBER, "5"},
		{token.SEMI, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.NUMBER, "5"},
		{token.LESS, "<"},
		{token.NUMBER, "0"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMI, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMI, ";"},
		{token.RBRACE, "}"},
		{token.NUMBER, "5"},
		{token.EQUALS, "=="},
		{token.NUMBER, "5"},
		{token.SEMI, ";"},
		{token.NUMBER, "5"},
		{token.NOTEQUALS, "!="},
		{token.NUMBER, "6"},
		{token.SEMI, ";"},
		{token.STRING, "foobar"},
		{token.SEMI, ";"},
		{token.STRING, "foo bar"},
		{token.SEMI, ";"},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.SEMI, ";"},
		{token.IDENTIFIER, "cccc"},
		{token.DOT, "."},
		{token.IDENTIFIER, "call"},
		{token.LBRACE, "{"},
		{token.STRING, "foo"},
		{token.DOT, "."},
		{token.STRING, "bar"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.NUMBER, "1"},
		{token.COLON, ":"},
		{token.NUMBER, "3"},
		{token.RBRACKET, "]"},
		{token.NUMBER, "5"},
		{token.MOD, "%"},
		{token.NUMBER, "4"},
		{token.IMPORT, "import"},
		{token.IDENTIFIER, "tests"},
		{token.IDENTIFIER, "x"},
		{token.LOGICALAND, "&&"},
		{token.IDENTIFIER, "y"},
		{token.IDENTIFIER, "x"},
		{token.LOGICALOR, "||"},
		{token.IDENTIFIER, "y"},
		{token.STRUCT, "struct"},
		{token.FOR, "for"},
		{token.NUMBER, "100.001"},
		{token.NOTEQUALS, "!="},
		{token.NUMBER, "0.123"},
		{token.EOF, ""},
	}

	reader := strings.NewReader(input)
	l := New(reader)

	for i, tt := range tests {
		if i == 106 {
			i = 106
		}
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got %q", i, token.TokenTypeLiteral[tt.expectedType], token.TokenTypeLiteral[tok.Type])
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenvale wrong. expected=%q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
