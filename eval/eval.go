package eval

import (
	"fmt"
	"panda/ast"
	"panda/parse"
)

type Interpreter struct {
	p *parse.Parser
}

func New(p *parse.Parser) *Interpreter {
	return &Interpreter{p}
}

func (inter *Interpreter) Eval() interface{} {
	exp := inter.p.ParseExpression(parse.LOWEST)
	fmt.Printf("%v\n", exp)
	return inter.eval(exp)
}
func (inter *Interpreter) eval(exp ast.Expression) interface{} {
	switch express := exp.(type) {
	case *ast.NumberExpression:
		// fmt.Println(express.Value)
		return express.Value
	case *ast.IdentifierExpression:
		fmt.Printf("Id(%s)", express.Value)
		return nil
	case *ast.InfixExpression:
		switch express.Operator {
		case "+":
			leftValue := inter.eval(express.Left)
			rightValue := inter.eval(express.Right)
			return leftValue.(int64) + rightValue.(int64)

		case "-":
			leftValue := inter.eval(express.Left)
			rightValue := inter.eval(express.Right)
			return leftValue.(int64) - rightValue.(int64)

		case "*":
			leftValue := inter.eval(express.Left)
			rightValue := inter.eval(express.Right)
			return leftValue.(int64) * rightValue.(int64)

		case "/":
			leftValue := inter.eval(express.Left)
			rightValue := inter.eval(express.Right)
			return leftValue.(int64) / rightValue.(int64)
		}
	default:
		panic(fmt.Errorf("未识别的错误"))
	}
	return nil
}
