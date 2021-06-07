package parse

import (
	"fmt"
	"panda/ast"
	"panda/token"
	"strconv"
)

func (p *Parser) parseNumber() ast.Expression {
	exp := &ast.NumberExpression{}
	exp.Token = p.curToken
	var value interface{}
	v, err := strconv.ParseInt(p.curToken.Literal, 10, 64)
	if err == nil {
		value = v
	}
	if err != nil {
		fv, err := strconv.ParseFloat(p.curToken.Literal, 64)
		if err != nil {
			p.errs = append(p.errs, fmt.Errorf("解析number表达式出错 line: %d, pos: %d", p.curToken.Line, p.curToken.Position))
			return nil
		}
		value = fv
	}
	exp.Value = value
	return exp
}

func (p *Parser) parseIdent() ast.Expression {
	exp := &ast.IdentifierExpression{}
	exp.Token = p.curToken
	exp.Value = p.curToken.Literal
	return exp
}

func (p *Parser) parseString() ast.Expression {
	exp := &ast.StringExpression{}
	exp.Token = p.curToken
	exp.Value = p.curToken.Literal
	return exp
}

func (p *Parser) parseBool() ast.Expression {
	exp := &ast.BoolExpression{}
	exp.Token = p.curToken
	if p.curTokenIs(token.TRUE) {
		exp.Value = true
		return exp
	} else if p.curTokenIs(token.FALSE) {
		exp.Value = false
		return exp
	}
	return exp
}
