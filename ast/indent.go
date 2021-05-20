package ast

import "panda/token"

type IdentifierExpression struct {
	Token token.Token
	Value string
}

func (ide *IdentifierExpression) ExpressionNode()      {}
func (ide *IdentifierExpression) TokenLiteral() string { return ide.Token.Literal }
func (ide *IdentifierExpression) String() string {
	return ide.Token.Literal
}
