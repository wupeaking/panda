package parse

import (
	"fmt"
	"panda/ast"
	"panda/token"
)

func (p *Parser) parseVarStatement() *ast.VarStatement {
	varStmt := ast.VarStatement{}
	varToken := p.curToken
	p.forwardToken()
	if !p.curTokenIs(token.IDENTIFIER) {
		p.errs = append(p.errs,
			fmt.Errorf("变量声明后面应该是标识符, line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
		return nil
	}

	if p.nextTokenIs(token.ASSIGN) {
		varStmt.Token = varToken
		varStmt.Name = &ast.IdentifierExpression{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		p.forwardToken()
		p.forwardToken()
		varStmt.Value = p.ParseExpression(LOWEST)
		p.forwardToken()
		if !p.curTokenIs(token.SEMI) {
			p.errs = append(p.errs,
				fmt.Errorf("期待一个;在line: %d, pos: %d",
					p.curToken.Line, p.curToken.Position,
				),
			)
			return nil
		}
		return &varStmt

	} else if p.nextTokenIs(token.SEMI) {
		varStmt.Token = varToken
		varStmt.Name = &ast.IdentifierExpression{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		p.forwardToken()
		return &varStmt
	} else {
		p.errs = append(p.errs,
			fmt.Errorf("变量声明后面应该是等号或者分号, line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
		return nil
	}
}

func (p *Parser) parseAssginStatement() *ast.AssginStatement {
	assginStmt := ast.AssginStatement{}
	idToken := p.curToken
	p.forwardToken()
	assginStmt.Name = &ast.IdentifierExpression{
		Token: idToken,
		Value: idToken.Literal,
	}
	p.forwardToken()
	assginStmt.Value = p.ParseExpression(LOWEST)
	p.forwardToken()
	if !p.curTokenIs(token.SEMI) {
		p.errs = append(p.errs,
			fmt.Errorf("期待一个分号在line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
	}
	return &assginStmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	retStmt := ast.ReturnStatement{}
	retStmt.Token = p.curToken
	p.forwardToken()
	retStmt.Value = p.ParseExpression(LOWEST)
	p.forwardToken()
	if !p.curTokenIs(token.SEMI) {
		p.errs = append(p.errs,
			fmt.Errorf("期待一个分号在line: %d, pos: %d",
				p.curToken.Line, p.curToken.Position,
			),
		)
	}
	return &retStmt
}
