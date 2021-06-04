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
	LOGIC  // 逻辑运算
	SUM    // + -
	MUL    // / * %
	PREFIX // 前缀表达式的优先级
	INDEX  // 数组和map索引
	CALL   // 当进入前缀表达式时  要保证调用表达式比他级别高 比如 +add() 如果不设置这个级别 会返回 +add 而不是+add()
)

// 其实只有中缀的情况 才会需要设置优先级
var tokenPrecedenceMap = map[token.TokenType]int{
	token.PLUS:            SUM,
	token.MINUS:           SUM,
	token.MUL:             MUL,
	token.DIV:             MUL,
	token.MOD:             MUL,
	token.EQUALS:          LOGIC,
	token.NOTEQUALS:       LOGIC,
	token.LESS:            LOGIC,
	token.LESSOREQUALS:    LOGIC,
	token.GREATER:         LOGIC,
	token.GREATEROREQUALS: LOGIC,
	token.LPAREN:          CALL,
	token.LBRACKET:        INDEX,
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
	// () 将左括号为前缀 (做前缀
	parse.registerPrefixExpr(token.LPAREN, parse.parseGroupExpression)
	// ! 作为前缀 但是这里设计! 优先级 是最低的  那么就不能用paresPrefixExprssion 函数 要重新写一个
	parse.registerPrefixExpr(token.BANG, parse.paresBangExprssion)
	// 变量表达式 和解析number类似
	parse.registerPrefixExpr(token.IDENTIFIER, parse.parseIdent)
	// 匿名函数表达式
	parse.registerPrefixExpr(token.FUNC, parse.parseAnonymousFunctionExprssion)
	// [ 作为前缀的 数组表达式
	parse.registerPrefixExpr(token.LBRACKET, parse.parseArrayExprssion)

	// 中缀
	// + - * /
	parse.registerInfixExpr(token.PLUS, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.MINUS, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.MUL, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.DIV, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.MOD, parse.paresInfixExprssion)

	// > < >= <= == !=
	parse.registerInfixExpr(token.EQUALS, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.NOTEQUALS, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.LESS, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.LESSOREQUALS, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.GREATER, parse.paresInfixExprssion)
	parse.registerInfixExpr(token.GREATEROREQUALS, parse.paresInfixExprssion)

	// ( 说明是函数调用 当做中缀表达式的时候 注意 优先级  如果调用到中缀 说明前面肯定有其他前缀表达式 这个时候 他应该是比前缀表达式的优先级还高
	// 否则 对于+add() 会处理错误
	parse.registerInfixExpr(token.LPAREN, parse.parseCallFunctionExprssion)

	// [ 索引表达式 同理 它应该比前缀的优先级高  +a[1]--> + a[1]   1*a[]--> 1 * a[] call()[1]--> call() [1] 所以 index索引优先级
	// 应该是比前缀的高 但是要比函数调用低
	parse.registerInfixExpr(token.LBRACKET, parse.parseIndexExpression)

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
		stmt := p.parserASTNode()
		if stmt == nil {
			panic(p.errs)
		}
		root.NodeTrees = append(root.NodeTrees, stmt)
		p.forwardToken()
	}
	return &root
}

func (p *Parser) parserASTNode() ast.Node {
	switch p.curToken.Type {
	case token.VAR:
		// 解析声明表达式
		return p.parseVarStatement()
	case token.RETURN:
		// return 语句
		return p.parseReturnStatement()

	case token.FUNC:
		// 函数声明
		return p.parseFunctionStatement()

	case token.IF:
		// if语句
		return p.parseIfStatement()

	case token.FOR:
		// for
		return p.parseForStatement()
	case token.BREAK:
		// break
		return p.parseBreakStatement()

	case token.IDENTIFIER:
		if p.nextTokenIs(token.ASSIGN) {
			// todo:: 后面支持 数组和map的赋值
			return p.parseAssginStatement()
		}
		fallthrough
	default:
		return p.parseExpressStatement()
		// p.errs = append(p.errs, fmt.Errorf("未知的语句处理 token: %v",
		// 	token.TokenType2Name(p.curToken.Type)))
	}
}

func (p *Parser) parseExpressStatement() *ast.ExpressStatement {
	// 解析一个表达式语句  比如 a; 1+2+3; 这些表达式作为单独的语句 没有被接收
	startToken := p.curToken
	exp := p.ParseExpression(LOWEST)
	p.forwardToken()
	if !p.curTokenIs(token.SEMI) {
		p.errs = append(p.errs,
			fmt.Errorf("期待一个;在line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
		return nil
	}
	es := ast.ExpressStatement{Token: startToken, Expression: exp}

	return &es
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
