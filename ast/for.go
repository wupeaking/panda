package ast

import (
	"bytes"
	"fmt"
	"panda/token"
)

type ForStatement struct {
	Token          token.Token
	FirstCondition Node
	MidCondition   Node
	LastCondition  Node
	LoopBody       *BlockStatement
}

func (fs *ForStatement) StatementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer
	out.WriteString(fmt.Sprintf("for("))
	if fs.FirstCondition != nil {
		out.WriteString(fs.FirstCondition.String())
		out.WriteString("; ")
	}
	if fs.MidCondition != nil {
		out.WriteString(fs.MidCondition.String())
		out.WriteString("; ")
	}
	if fs.LastCondition != nil {
		out.WriteString(fs.LastCondition.String())
		out.WriteString(") ")
	}
	out.WriteString(fs.LoopBody.String())
	return out.String()
}
