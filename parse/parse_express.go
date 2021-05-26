package parse

import (
	"fmt"
	"panda/ast"
	"panda/token"
)

func (p *Parser) paresAnonymousFunctionExprssion() ast.Expression {
	exp := ast.AnonymousFuncExpression{}
	exp.Token = p.curToken // function
	if !p.nextTokenIs(token.LPAREN) {
		p.errs = append(p.errs,
			fmt.Errorf("line: %d, pos: %d 期待( 实际是%s ",
				p.curToken.Line, p.curToken.Position, token.TokenType2Name(p.curToken.Type)))
		return nil
	}
	p.forwardToken() // (
	p.forwardToken() // a, b, c

	args := make([]ast.Expression, 0)
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

func (p *Parser) paresCallFunctionExprssion(funcName ast.Expression) ast.Expression {
	exp := ast.CallExpression{}
	exp.FuncName = funcName
	exp.Token = p.curToken // (
	p.forwardToken()       // exp1, exp2, exp3

	args := make([]ast.Expression, 0)
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
