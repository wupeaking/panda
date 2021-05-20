package ast

import "panda/token"

// 数字表达式
type NumberExpression struct {
	Token token.Token
	Value interface{}
}

func (ne *NumberExpression) ExpressionNode()      {}
func (ne *NumberExpression) TokenLiteral() string { return ne.Token.Literal }

func (ne *NumberExpression) String() string {
	return ne.Token.Literal
}
