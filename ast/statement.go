package ast

import (
	"bytes"
	"fmt"
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

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) StatementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString("return ")
	out.WriteString(rs.Value.String())
	out.WriteString(" ;")
	return out.String()
}

// FunctionStatement  函数声明
type FunctionStatement struct {
	Token      token.Token
	Name       *IdentifierExpression
	Args       []Expression // 应该是变量表达式
	ReturnType Statement    // 保留 暂时设计的函数不需要声明返回类型
	FuncBody   *BlockStatement
}

func (fs *FunctionStatement) StatementNode()       {}
func (fs *FunctionStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *FunctionStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("< function %s(", fs.Name.Value))
	for i := range fs.Args {
		out.WriteString(fs.Args[i].String())
		out.WriteString(", ")
	}
	out.WriteString(") ")
	out.WriteString(fs.FuncBody.String())
	out.WriteString(" >")
	return out.String()
}

type BreakStatement struct {
	Token token.Token
}

func (bk *BreakStatement) StatementNode()       {}
func (bk *BreakStatement) TokenLiteral() string { return bk.Token.Literal }
func (bk *BreakStatement) String() string {
	var out bytes.Buffer
	out.WriteString("break;")
	return out.String()
}
