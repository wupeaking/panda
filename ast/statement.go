package ast

import (
	"bytes"
	"panda/token"
)

type Statement interface {
	Node
	StatementNode()
}

// VarStatement 变量声明
type VarStatement struct {
	Token token.Token
	Name  *IdentifierExpression
	Value Expression
}

func (vs *VarStatement) StatementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }
func (vs *VarStatement) String() string {
	var out bytes.Buffer
	out.WriteString("(var ")
	out.WriteString(vs.Name.TokenLiteral())
	if vs.Value != nil {
		out.WriteString(" = ")
		out.WriteString(" " + vs.Value.String() + " ")
	}
	out.WriteString(")")
	return out.String()
}

// AssginStatement 赋值语句
type AssginStatement struct {
	Token token.Token
	Name  *IdentifierExpression
	Value Expression
}

func (as *AssginStatement) StatementNode()       {}
func (as *AssginStatement) TokenLiteral() string { return as.Token.Literal }
func (as *AssginStatement) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(as.Name.TokenLiteral())
	out.WriteString(" " + as.Token.Literal + " ")
	out.WriteString(")")
	return out.String()
}
