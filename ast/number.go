package ast

import (
	"fmt"
	"panda/token"
)

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

type AddessExpression struct {
	Token token.Token
	Value string // 实际的地址
}

func (ide *AddessExpression) ExpressionNode()      {}
func (ide *AddessExpression) TokenLiteral() string { return ide.Token.Literal }
func (ide *AddessExpression) String() string {
	return ide.Token.Literal
}

type ThisExpression struct {
	Token token.Token
	Value string // 指向当前合约地址
}

func (te *ThisExpression) ExpressionNode()      {}
func (te *ThisExpression) TokenLiteral() string { return te.Token.Literal }
func (te *ThisExpression) String() string {
	return te.Token.Literal + fmt.Sprintf("(%s)", te.Value)
}
