package parse

import "panda/token"

func (p *Parser) forwardToken() {
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(ty token.TokenType) bool {
	return p.curToken.Type == ty
}

func (p *Parser) nextTokenIs(ty token.TokenType) bool {
	return p.curToken.Type == ty
}

func (p *Parser) curTokenPrecedence() int {
	pre, ok := tokenPrecedenceMap[p.curToken.Type]
	if ok {
		return pre
	}
	return LOWEST
}

func (p *Parser) nextTokenPrecedence() int {
	pre, ok := tokenPrecedenceMap[p.nextToken.Type]
	if ok {
		return pre
	}
	return LOWEST
}
