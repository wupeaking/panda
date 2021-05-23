package parse

import (
	"fmt"
	"panda/ast"
	"panda/token"
)

func (p *Parser) paresVarStatement() *ast.VarStatement {
	varStmt := ast.VarStatement{}
	varToken := p.curToken
	p.forwardToken()

	if p.nextTokenIs(token.ASSIGN) {
		varStmt.Token = varToken
		varStmt.Name = &ast.IdentifierExpression{
			Token: p.curToken,
			Value: p.curToken.Literal,
		}
		p.forwardToken()
		varStmt.Value = p.ParseExpression(LOWEST)
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
