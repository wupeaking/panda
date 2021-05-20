package parse

import (
	"panda/ast"
	"panda/lexer"
	"panda/token"
)

const (
	_      = iota
	LOWEST // 最低
	SUM    // + -
	MUL    // / * %
	PREFIX // 前缀表达式的优先级
)

var tokenPrecedenceMap = map[token.TokenType]int{
	token.PLUS:  SUM,
	token.MINUS: SUM,
	token.MUL:   MUL,
	token.DIV:   MUL,
	token.MOD:   MUL,
}

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	nextToken token.Token
	errs      []error

	prefixExprParseFns map[token.TokenType]prefixExprParseFunc
	infixExprParseFns  map[token.TokenType]infixExprParseFunc
}

type (
	prefixExprParseFunc func() ast.Expression
	infixExprParseFunc  func(ast.Expression) ast.Expression
)

func New(l *lexer.Lexer) *Parser {
	parse := &Parser{l: l,
		prefixExprParseFns: make(map[token.TokenType]prefixExprParseFunc),
		infixExprParseFns:  make(map[token.TokenType]infixExprParseFunc),
	}
	// 注册前缀
	// number
	parse.registerPrefixExpr(token.NUMBER, parse.parseNumber)
	// + - 可以做前缀
	parse.registerPrefixExpr(token.PLUS, parse.paresPrefixExprssion)
	parse.registerPrefixExpr(token.MINUS, parse.paresPrefixExprssion)

	// 中缀
	// + - * /
	parse.registerInfixExpr(token.PLUS, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.MINUS, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.MUL, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.DIV, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.MOD, parse.paresInfixExprssion)

	parse.forwardToken()
	parse.forwardToken()
	return parse
}

func (p *Parser) registerPrefixExpr(ty token.TokenType, fn prefixExprParseFunc) {
	p.prefixExprParseFns[ty] = fn
}

func (p *Parser) registerInfixExpr(ty token.TokenType, fn infixExprParseFunc) {
	p.infixExprParseFns[ty] = fn
}

func (p *Parser) ParseExpression(precedence int) ast.Expression {
	prefix := p.prefixExprParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	for !p.nextTokenIs(token.SEMI) && precedence < p.nextTokenPrecedence() && !p.curTokenIs(token.EOF) {
		infix := p.infixExprParseFns[p.nextToken.Type]
		if infix == nil {
			return leftExp
		}
		p.forwardToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (p *Parser) paresPrefixExprssion() ast.Expression {
	exp := ast.PrefixExpression{}
	exp.Token = p.curToken
	exp.Operator = p.curToken.Literal
	// todo:: 前缀的优先级 后面需要定义一下
	p.forwardToken()
	exp.Right = p.ParseExpression(PREFIX)
	return &exp
}

func (p *Parser) paresInfixExprssion(left ast.Expression) ast.Expression {
	exp := ast.InfixExpression{}
	exp.Token = p.curToken
	exp.Operator = p.curToken.Literal
	exp.Left = left
	// 获取当前token的优先级
	pre := p.curTokenPrecedence()
	p.forwardToken()
	exp.Right = p.ParseExpression(pre)
	return &exp
}
