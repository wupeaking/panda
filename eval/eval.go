package eval

import (
	"errors"
	"fmt"
	"panda/ast"
	"panda/parse"
)

var (
	returnError = errors.New("exit")
)

type Interpreter struct {
	p            *parse.Parser
	scopeManager *ScopeManager
}

func New(p *parse.Parser) *Interpreter {
	return &Interpreter{p: p, scopeManager: NewScopeManager()}
}

func (inter *Interpreter) Eval() interface{} {
	astTree := inter.p.ParserAST()
	if err := inter.p.Errors(); err != nil {
		panic(err)
	}
	// fmt.Printf("%v\n", astTree)
	return inter.evalProgram(astTree)
}

func (inter *Interpreter) evalProgram(astTree *ast.ProgramAST) interface{} {
	ret, err := inter.evalASTNodes(astTree.NodeTrees)
	if err != nil {
		panic(err)
	}
	return ret
}

func (inter *Interpreter) evalASTNodes(nodes []ast.Node) (interface{}, error) {
	for i := range nodes {
		switch x := nodes[i].(type) {
		case ast.Statement:
			v, err := inter.evalStatement(x)
			if err == returnError {
				return v, nil
			}
			if err != nil {
				return nil, err
			}

		case ast.Expression:
			v, err := inter.evalExpress(x)
			if err != nil {
				return nil, err
			}
			fmt.Printf("%v\n", v)
		}
	}
	return nil, nil
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
			inter.scopeManager.SetValue(statement.Name.Value, v, true)
			// VarMap[statement.Name.Value] = v
			return nil, nil
		} else {
			inter.scopeManager.SetValue(statement.Name.Value, nil, true)
		}
	case *ast.AssginStatement:
		// 判断变量是否存在
		// _, ok := VarMap[statement.Name.Value]
		if !inter.scopeManager.VarExists(statement.Name.Value) {
			return nil, fmt.Errorf("变量%s未定义", statement.Name.Value)
		}
		v, err := inter.evalExpress(statement.Value)
		if err != nil {
			return nil, err
		}
		//VarMap[statement.Name.Value] = v
		inter.scopeManager.SetValue(statement.Name.Value, v, false)

	case *ast.ExpressStatement:
		v, err := inter.evalExpress(statement.Expression)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%v\n", v)

	case *ast.ReturnStatement:
		// 返回语句
		v, err := inter.evalExpress(statement.Value)
		if err != nil {
			return nil, err
		}
		return v, returnError

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
		// fmt.Printf("Id(%s)\n", express.Value)
		v, scope := inter.scopeManager.GetValue(express.Value)
		// v, ok := VarMap[express.Value]
		if scope == nil {
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

	case *ast.AnonymousFuncExpression:
		// add(a, b){return a+b;};
		// 返回匿名函数表达式
		return express, nil

	case *ast.CallExpression:
		// add(1, 2)
		v, _, err := inter.evalFunctionCall(express)
		return v, err

	default:
		return nil, fmt.Errorf("未识别的表达式: %v", express)
	}
	return nil, nil
}

func (inter *Interpreter) evalFunctionCall(funcNode *ast.CallExpression) (interface{}, bool, error) {
	// 函数调用可以是匿名函数表达式的调用 也可以是函数语句调用
	funcExp, err := inter.evalExpress(funcNode.FuncName)
	if err != nil {
		return nil, false, err
	}

	anonyFunc, ok := funcExp.(*ast.AnonymousFuncExpression)
	if !ok {
		return nil, false, fmt.Errorf("%v不是函数", funcNode.FuncName)
	}
	if len(anonyFunc.Args) != len(funcNode.Arguments) {
		return nil, false, fmt.Errorf("实参和形参数量不一致")
	}
	inter.scopeManager.Push(FuncScope)
	defer inter.scopeManager.Pop()

	// 进行传值
	for i := range anonyFunc.Args {
		arg, ok := anonyFunc.Args[i].(*ast.IdentifierExpression)
		if !ok {
			return nil, false, fmt.Errorf("%v 参数类型错误", funcNode.FuncName)
		}
		argValue, err := inter.evalExpress(funcNode.Arguments[i])
		if err != nil {
			return nil, false, err
		}
		inter.scopeManager.SetValue(arg.Value, argValue, true)
	}
	//anonyFunc.FuncBody.Statements
	v, err := inter.evalASTNodes(anonyFunc.FuncBody.Statements)
	return v, true, err
}
