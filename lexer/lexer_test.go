package lexer

import (
	"testing"

	"github.com/tlauffs/go_interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){}`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	in := New(input)
	for i, tt := range tests {
		tok := in.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("token type test failed: expected: %q ; got %q", tt.expectedLiteral, tok.Type)
		}
	}
}
