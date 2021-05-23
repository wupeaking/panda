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
	if err := inter.p.Errors(); err != nil {
		panic(err)
	}
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
			// todo:: 待支持浮点数
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
		default:
			panic(fmt.Errorf("中缀表达式: %s 不支持%s 操作符", express.String(), express.Operator))
		}

	case *ast.PrefixExpression:
		switch express.Operator {
		case "-":
			rightValue := inter.eval(express.Right)
			return -(rightValue.(int64))
		case "+":
			rightValue := inter.eval(express.Right)
			return -(rightValue.(int64))
		case "!":
			// 逻辑取反
			rightValue := inter.eval(express.Right)
			switch x := rightValue.(type) {
			case int64:
				return !(x > 0)
			case float64:
				return !(x > 0)
			case bool:
				return !x
			default:
				panic(fmt.Errorf("!运算不支持此类型的值: %v", x))
			}
		}
	default:
		panic(fmt.Errorf("未识别的表达式: %v", express))
	}
	return nil
}
