package ast

import (
	"bytes"
	"panda/token"
)

type Expression interface {
	Node
	ExpressionNode()
}

// InfixExpression 中缀表达式
type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Right    Expression
	Operator string // 操作符
}

func (ie *InfixExpression) ExpressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }
func (ie *InfixExpression) String() string {
	// return fmt.Sprintf(`Op(%s)
	// 	left: %s
	// 	right: %s
	// `, ie.Operator, ie.Left.String(), ie.Right.String())
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")
	return out.String()
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) ExpressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
