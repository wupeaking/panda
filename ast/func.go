package ast

import (
	"bytes"
	"panda/token"
)

/*
var fn = function(a, b, c, d) {
	a = b;
	b = c;
}

var fn = function(a, b, c, d)  {
	a = b;
	b = c;
}

*/

// 定义匿名函数表达式
type AnonymousFuncExpression struct {
	Token      token.Token
	Args       []Expression // 应该是变量表达式
	ReturnType Statement    // todo: 定义返回类型
	FuncBody   *BlockStatement
}

func (afe *AnonymousFuncExpression) ExpressionNode()      {}
func (afe *AnonymousFuncExpression) TokenLiteral() string { return afe.Token.Literal }
func (afe *AnonymousFuncExpression) String() string {
	var out bytes.Buffer
	out.WriteString("< function (")
	for i := range afe.Args {
		out.WriteString(afe.Args[i].String())
		out.WriteString(", ")
	}
	out.WriteString(") ")
	out.WriteString(afe.FuncBody.String())
	out.WriteString(" >")
	return out.String()
}
