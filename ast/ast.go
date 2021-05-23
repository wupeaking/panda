package ast

// 定义Node接口

type Node interface {
	String() string       // 返回Node的文字表示
	TokenLiteral() string // 返回token的字面值
}

type ProgramAST struct {
	NodeTrees []Node
}
