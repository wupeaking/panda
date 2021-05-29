package lexer

import (
	"io"
	"panda/token"
	"text/scanner"
)

type Lexer struct {
	*scanner.Scanner
}

func New(reader io.Reader) *Lexer {
	lexer := &Lexer{
		&scanner.Scanner{},
	}
	lexer.Init(reader)
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	lexer.skipWhitespace()
	ch := lexer.Next()
	switch ch {
	case '(', ')', '{', '}', '[', ']', ',', ';', ':', '.', '*', '%':
		tokenType := token.LiteralTokenType[string(ch)]
		return token.Token{Type: tokenType, Literal: string(ch), Line: lexer.Pos().Line, Position: lexer.Pos().Offset}

	case '/':
		// 考虑是否为注释
		if lexer.Peek() == '/' {
			lexer.skipComments()
			return lexer.NextToken()
		} else {
			tokenType := token.LiteralTokenType[string(ch)]
			return token.Token{Type: tokenType, Literal: string(ch), Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}

	case '=': // 赋值 或者 等于
		if lexer.Peek() == '=' {
			lexer.Next()
			return token.Token{Type: token.EQUALS, Literal: "==", Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		} else {
			return token.Token{Type: token.ASSIGN, Literal: "=", Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}

	case '+': // 加 或者累加
		if lexer.Peek() == '+' {
			lexer.Next()
			return token.Token{Type: token.PLUSPLUS, Literal: "++", Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		} else {
			return token.Token{Type: token.PLUS, Literal: "+", Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
	case '-': // 减 或者减减
		if lexer.Peek() == '-' {
			lexer.Next()
			return token.Token{Type: token.MINUSMINUS, Literal: "--", Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		} else {
			return token.Token{Type: token.MINUS, Literal: "-", Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
	case '!': //  ! !=
		if lexer.Peek() == '=' {
			lexer.Next()
			return token.Token{Type: token.NOTEQUALS, Literal: "!=", Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		} else {
			return token.Token{Type: token.BANG, Literal: "!", Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
	case '<', '>': // < <= << > >> >=
		peek := lexer.Peek()
		if peek == '=' {
			lexer.Next()
			letter := string(ch) + "="
			return token.Token{Type: token.LiteralTokenType[letter], Literal: letter, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		} else if peek == ch {
			lexer.Next()
			letter := string(ch) + string(ch)
			return token.Token{Type: token.LiteralTokenType[letter], Literal: letter, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		} else {
			letter := string(ch)
			return token.Token{Type: token.LiteralTokenType[letter], Literal: letter, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
	case '|', '&': // | || & &&
		peek := lexer.Peek()
		if peek == ch {
			lexer.Next()
			letter := string(ch) + string(ch)
			return token.Token{Type: token.LiteralTokenType[letter], Literal: letter, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		} else {
			letter := string(ch)
			return token.Token{Type: token.LiteralTokenType[letter], Literal: letter, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
	case -1:
		return token.Token{Type: token.EOF, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
	}
	return lexer.readMutliChars(ch)
}

func (lexer *Lexer) readMutliChars(char rune) token.Token {
	// 判断是否是字母
	switch {
	case char == -1:
		return token.Token{Type: token.EOF, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
	case isNumber(char):
		// 读取数字
		return lexer.readNumber(char)
	case isLetter(char):
		// 读取 变量名 或者关键字
		return lexer.readIdentifierOrKeywords(char)
	case isQuote(char):
		return lexer.readString(char)
	default:
		return token.Token{Type: token.ILLEGAL, Literal: string(char), Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
	}
}

func (lexer *Lexer) readNumber(char rune) token.Token {
	// 尝试继续往下读 读取到不是number为止 number 可以是 整数  也可以是浮点数
	switch {
	case char == '0':
		if lexer.Peek() != '.' {
			return token.Token{Type: token.ILLEGAL, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
		lexer.Next()
		str := string(char) + "."
		for {
			peek := lexer.Peek()
			if isNumber(peek) {
				lexer.Next()
				str += string(peek)
			} else {
				break
			}
		}
		if len(str) < 3 {
			return token.Token{Type: token.ILLEGAL, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
		return token.Token{Type: token.NUMBER, Literal: str, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}

	case isNumber(char):
		str := string(char)
		hasPoint := false
		for {
			peek := lexer.Peek()
			if peek == '.' && !hasPoint {
				hasPoint = true
				lexer.Next()
				str += string(peek)
				continue
			}
			if isNumber(peek) {
				lexer.Next()
				str += string(peek)
				continue
			}
			break
		}

		return token.Token{Type: token.NUMBER, Literal: str, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
	}
	return token.Token{Type: token.ILLEGAL, Literal: string(char), Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
}

// 暂时只支持"" 这种字符串表示
func (lexer *Lexer) readString(char rune) token.Token {
	var str string
	for {
		char := lexer.Next()
		if char == 0 {
			return token.Token{Type: token.ILLEGAL, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
		if char == '"' {
			// 说明读取完成
			return token.Token{Type: token.STRING, Literal: str, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
		}
		str += string(char)
	}
}

func (lexer *Lexer) readIdentifierOrKeywords(char rune) token.Token {
	// 尝试继续往下读 知道读取到不是letter为止
	str := string(char)
	for {
		peek := lexer.Peek()
		if isLetter(peek) {
			lexer.Next()
			str += string(peek)
		} else {
			break
		}
	}
	ty, ok := token.LiteralTokenType[str]
	if ok {
		return token.Token{Type: ty, Literal: str, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
	}
	return token.Token{Type: token.IDENTIFIER, Literal: str, Line: lexer.Pos().Line, Position: lexer.Pos().Offset}
}

func isNumber(char rune) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func isLetter(char rune) bool {
	if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_' {
		return true
	}
	return false
}

func isQuote(char rune) bool {
	return char == '"'
}

func (lexer *Lexer) skipWhitespace() {
	// 空格 table  回车 换行 忽略
	for {
		ch := lexer.Peek()
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			lexer.Next()
		} else {
			break
		}
	}
}

// skipComments 跳过注释
func (lexer *Lexer) skipComments() {
	// // 直到 \n 或者\n\r
	if lexer.Peek() == '/' {
		for {
			ch := lexer.Peek()
			if ch == scanner.EOF {
				break
			}
			if ch != '\n' {
				lexer.Next()
				continue
			} else {
				lexer.Next() // eat \n
				if lexer.Peek() == '\r' {
					lexer.Next() // eat \r
				}
				break
			}
		}
	}
}
