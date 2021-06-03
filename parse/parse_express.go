package parse

import (
	"fmt"
	"panda/ast"
	"panda/token"
)

func (p *Parser) parseAnonymousFunctionExprssion() ast.Expression {
	exp := ast.AnonymousFuncExpression{}
	exp.Token = p.curToken // function
	if !p.nextTokenIs(token.LPAREN) {
		p.errs = append(p.errs,
			fmt.Errorf("line: %d, pos: %d 期待( 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() // (
	p.forwardToken() // a, b, c 或者)

	args := make([]ast.Expression, 0)
	if p.curTokenIs(token.RPAREN) {
		exp.Args = args
		goto parseBody
	}

	for {
		argExp := p.ParseExpression(LOWEST)
		if argExp == nil {
			return nil
		}
		args = append(args, argExp)

		if p.nextTokenIs(token.RPAREN) {
			break
		}
		if p.nextTokenIs(token.COMMA) {
			p.forwardToken() // ,
			p.forwardToken() //
			continue
		} else {
			p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待,或者) 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
			return nil
		}
	}
	exp.Args = args
	p.forwardToken() // )

parseBody:
	if !p.nextTokenIs(token.LBRACE) {
		p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待{ 实际是%s ",
			p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	exp.FuncBody = &ast.BlockStatement{Token: p.curToken}
	p.forwardToken() // {

	p.forwardToken()
	nodes := make([]ast.Node, 0)
	for {
		node := p.parserASTNode()
		if node == nil {
			return nil
		}
		nodes = append(nodes, node)
		if p.nextTokenIs(token.RBRACE) {
			break
		}
		p.forwardToken()
	}
	exp.FuncBody.Statements = nodes
	p.forwardToken()

	return &exp
}

// paresCallFunctionExprssion 解析函数调用  add() add()()()
func (p *Parser) parseCallFunctionExprssion(funcName ast.Expression) ast.Expression {
	exp := p.parseCallFunctionExprssionHelper(funcName)
	for p.nextTokenIs(token.LPAREN) {
		p.forwardToken()
		exp = p.parseCallFunctionExprssionHelper(exp)
	}
	return exp
}

func (p *Parser) parseCallFunctionExprssionHelper(funcName ast.Expression) ast.Expression {
	exp := ast.CallExpression{}
	// 可能的表达式是 add()
	// 也可能是 addFn()() 所以funcName 也能的值是id表达式和调用表达式
	exp.FuncName = funcName
	exp.Token = p.curToken // (
	p.forwardToken()       // exp1, exp2, exp3 或者 )

	args := make([]ast.Expression, 0)
	if p.curTokenIs(token.RPAREN) {
		exp.Arguments = args
		return &exp
	}

	for {
		argExp := p.ParseExpression(LOWEST)
		if argExp == nil {
			return nil
		}
		args = append(args, argExp)

		if p.nextTokenIs(token.RPAREN) {
			break
		}
		if p.nextTokenIs(token.COMMA) {
			p.forwardToken() // ,
			p.forwardToken() // next exp
			continue
		} else {
			p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待,或者) 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
			return nil
		}
	}
	exp.Arguments = args
	p.forwardToken() // )

	return &exp
}

func (p *Parser) parseArrayExprssion() ast.Expression {
	array := &ast.ArrayExpression{}
	array.Token = p.curToken
	p.forwardToken() //
	array.Exprs = make([]ast.Expression, 0)
	if p.nextTokenIs(token.RBRACKET) {
		p.forwardToken() //]
		return array

	}
	for {
		exp := p.ParseExpression(LOWEST)
		array.Exprs = append(array.Exprs, exp)
		if p.nextTokenIs(token.RBRACKET) {
			break
		}
		if !p.nextTokenIs(token.COLON) {
			p.errs = append(p.errs, fmt.Errorf("line: %d, pos: %d 期待, 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.nextToken.Type)))
			return nil
		}
	}
	p.forwardToken() // ]
	return array
}
