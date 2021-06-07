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

type StringExpression struct {
	Token token.Token
	Value interface{}
}

func (se *StringExpression) ExpressionNode()      {}
func (se *StringExpression) TokenLiteral() string { return se.Token.Literal }

func (se *StringExpression) String() string {
	return se.Token.Literal
}

type BoolExpression struct {
	Token token.Token
	Value interface{}
}

func (be *BoolExpression) ExpressionNode()      {}
func (be *BoolExpression) TokenLiteral() string { return be.Token.Literal }

func (be *BoolExpression) String() string {
	return be.Token.Literal
}
