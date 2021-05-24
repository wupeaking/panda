package ast

import (
	"bytes"
	"panda/token"
)

type Statement interface {
	Node
	StatementNode()
}

//ExpressStatement 表达式语句 实际是一个表达式 但是放在单独的语句执行 意思是返回值没有被接收
type ExpressStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressStatement) StatementNode()       {}
func (es *ExpressStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressStatement) String() string {
	var out bytes.Buffer
	out.WriteString("(express ")
	out.WriteString(es.Expression.String())
	out.WriteString(")")
	return out.String()
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
	out.WriteString(" = ")
	out.WriteString(as.Value.String())
	out.WriteString(")")
	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Node
}

func (bs *BlockStatement) StatementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for i := range bs.Statements {
		out.WriteString(bs.Statements[i].String())
		out.WriteString("\n")
	}
	out.WriteString("}")
	return out.String()
}
