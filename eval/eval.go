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
	astTree := inter.p.ParserAST()
	if err := inter.p.Errors(); err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", astTree)
	return inter.evalProgram(astTree)
}

func (inter *Interpreter) evalProgram(astTree *ast.ProgramAST) interface{} {
	for i := range astTree.NodeTrees {
		switch x := astTree.NodeTrees[i].(type) {
		case ast.Statement:
			_, err := inter.evalStatement(x)
			if err != nil {
				panic(err)
			}
		case ast.Expression:
			v, err := inter.evalExpress(x)
			if err != nil {
				panic(err)
			}
			fmt.Printf("%v\n", v)
		}
	}
	return nil
}

var VarMap = map[string]interface{}{}

func (inter *Interpreter) evalStatement(stmt ast.Statement) (interface{}, error) {
	switch statement := stmt.(type) {
	case *ast.VarStatement:
		// todo::存入变量表
		if statement.Value != nil {
			// 将计算结果存入变量
			v, err := inter.evalExpress(statement.Value)
			if err != nil {
				return nil, err
			}
			VarMap[statement.Name.Value] = v
			return nil, nil
		}
	default:
		return nil, fmt.Errorf("暂时未处理%v 语句", statement)
	}
	return nil, nil
}

func (inter *Interpreter) evalExpress(exp ast.Expression) (interface{}, error) {
	switch express := exp.(type) {
	case *ast.NumberExpression:
		// fmt.Println(express.Value)
		return express.Value, nil
	case *ast.IdentifierExpression:
		// todo:: 去变量表 找到对应的值 返回
		fmt.Printf("Id(%s)", express.Value)
		v, ok := VarMap[express.Value]
		if !ok {
			return nil, fmt.Errorf("未定义的变量: %v", express.Value)
		}
		return v, nil

	case *ast.InfixExpression:
		switch express.Operator {
		case "+":
			// todo:: 待支持浮点数
			leftValue, _ := inter.evalExpress(express.Left)
			rightValue, _ := inter.evalExpress(express.Right)
			return leftValue.(int64) + rightValue.(int64), nil

		case "-":
			leftValue, _ := inter.evalExpress(express.Left)
			rightValue, _ := inter.evalExpress(express.Right)
			return leftValue.(int64) - rightValue.(int64), nil

		case "*":
			leftValue, _ := inter.evalExpress(express.Left)
			rightValue, _ := inter.evalExpress(express.Right)
			return leftValue.(int64) * rightValue.(int64), nil

		case "/":
			leftValue, _ := inter.evalExpress(express.Left)
			rightValue, _ := inter.evalExpress(express.Right)
			return leftValue.(int64) / rightValue.(int64), nil
		default:
			panic(fmt.Errorf("中缀表达式: %s 不支持%s 操作符", express.String(), express.Operator))
		}

	case *ast.PrefixExpression:
		switch express.Operator {
		case "-":
			rightValue, _ := inter.evalExpress(express.Right)
			return -(rightValue.(int64)), nil
		case "+":
			rightValue, _ := inter.evalExpress(express.Right)
			return -(rightValue.(int64)), nil
		case "!":
			// 逻辑取反
			rightValue, _ := inter.evalExpress(express.Right)
			switch x := rightValue.(type) {
			case int64:
				return !(x > 0), nil
			case float64:
				return !(x > 0), nil
			case bool:
				return !x, nil
			default:
				return nil, fmt.Errorf("!运算不支持此类型的值: %v", x)
			}
		}
	default:
		return nil, fmt.Errorf("未识别的表达式: %v", express)
	}
	return nil, nil
}
