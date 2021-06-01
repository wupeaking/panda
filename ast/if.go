package ast

import (
	"bytes"
	"fmt"
	"panda/token"
)

type IFStatement struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IFStatement) StatementNode()       {}
func (is *IFStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IFStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf(" if(%s)", is.Condition.String()))
	out.WriteString(fmt.Sprintf("%s", is.Consequence.String()))
	if is.Alternative != nil {
		out.WriteString(" else ")
		out.WriteString(fmt.Sprintf("%s", is.Alternative.String()))
		out.WriteString(" \n")
	}
	return out.String()
}
