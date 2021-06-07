package ast

import (
	"bytes"
	"panda/token"
)

type MapExpression struct {
	Token token.Token
	KV    map[Expression]Expression // key 目前只能是number string
}

func (me *MapExpression) ExpressionNode()      {}
func (me *MapExpression) TokenLiteral() string { return me.Token.Literal }

func (me *MapExpression) String() string {
	var out bytes.Buffer
	out.WriteString("{\n")
	for k, v := range me.KV {
		out.WriteString(" " + k.String())
		out.WriteString(" : ")
		out.WriteString(v.String())
		out.WriteString(",\n ")
	}
	out.WriteString("} \n")
	return out.String()
}
