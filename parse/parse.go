package parse

import (
	"fmt"
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
	// () 将左括号祖册为前缀
	parse.registerPrefixExpr(token.LPAREN, parse.parseGroupExpression)
	// ! 作为前缀 但是这里设计! 优先级 是最低的  那么就不能用paresPrefixExprssion 函数 要重新写一个
	parse.registerPrefixExpr(token.BANG, parse.paresBangExprssion)

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

func (p *Parser) ParserAST() *ast.ProgramAST {
	root := ast.ProgramAST{NodeTrees: []ast.Node{}}
	for p.curToken.Type != token.EOF {
		stmt := p.ParseStament()
		if stmt == nil {
			panic(p.errs)
		}
		root.NodeTrees = append(root.NodeTrees, stmt)
		p.forwardToken()
	}
	return &root
}

func (p *Parser) ParseStament() ast.Statement {
	switch p.curToken.Type {
	case token.VAR:
		// 解析声明表达式
		return p.paresVarStatement()
	default:
		p.errs = append(p.errs, fmt.Errorf("未知的语句处理 token: %v"))
	}
	return nil
}

func (p *Parser) ParseExpression(precedence int) ast.Expression {
	prefix := p.prefixExprParseFns[p.curToken.Type]
	if prefix == nil {
		p.errs = append(p.errs, fmt.Errorf("不支持的前缀符号: %s", token.TokenType2Name(p.curToken.Type)))
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

func (p *Parser) parseGroupExpression() ast.Expression {
	// () 解析表达式 直到返回的是)
	p.forwardToken()

	// 应该是只有到了发现 ) 没有注册的解析函数才会返回
	exp := p.ParseExpression(LOWEST)
	if !p.nextTokenIs(token.RPAREN) {
		p.errs = append(p.errs, fmt.Errorf("期待右括号 结果是: %v", token.TokenType2Name(p.nextToken.Type)))
		return nil
	}
	p.forwardToken()
	return exp
}

func (p *Parser) paresBangExprssion() ast.Expression {
	exp := ast.PrefixExpression{}
	exp.Token = p.curToken
	exp.Operator = p.curToken.Literal
	// !操作符 我认为优先级最低
	p.forwardToken()
	exp.Right = p.ParseExpression(LOWEST)
	return &exp
}