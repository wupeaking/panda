package ast

import (
	"bytes"
	"panda/token"
)

// 数组表达式
type ArrayExpression struct {
	Token  token.Token
	Exprs  []Expression // 暂定数组支持的类型有 字符串 数组 数字 map 匿名函数
	Values []interface{}
}

func (ae *ArrayExpression) ExpressionNode()      {}
func (ae *ArrayExpression) TokenLiteral() string { return ae.Token.Literal }

func (ae *ArrayExpression) String() string {
	var out bytes.Buffer
	out.WriteString("[")
	for i := range ae.Exprs {
		out.WriteString(ae.Exprs[i].String())
		out.WriteString(", ")
	}
	out.WriteString("] ")
	return out.String()
}
