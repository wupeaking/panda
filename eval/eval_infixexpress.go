package eval

import (
	"fmt"
	"panda/ast"
	"strings"
)

func (inter *Interpreter) evalAddExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left + right, nil
		case float64:
			return float64(left) + right, nil
		case string:
			return fmt.Sprintf("%d", left) + right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left + float64(right), nil
		case float64:
			return float64(left) + right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}

	case string:
		switch right := rightValue.(type) {
		case int64:
			return left + fmt.Sprintf("%d", right), nil
		case string:
			return left + right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

func (inter *Interpreter) evalSubExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left - right, nil
		case float64:
			return float64(left) - right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left - float64(right), nil
		case float64:
			return float64(left) - right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

func (inter *Interpreter) evalMulExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left * right, nil
		case float64:
			return float64(left) * right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left * float64(right), nil
		case float64:
			return float64(left) * right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}

	case string:
		switch right := rightValue.(type) {
		case int64:
			if right < 1 {
				return nil, fmt.Errorf("参数类型错误")
			}
			for i := int64(1); i < right; i++ {
				left += left
			}
			return left, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

func (inter *Interpreter) evalDivExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left / right, nil
		case float64:
			return float64(left) / right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left / float64(right), nil
		case float64:
			return float64(left) / right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

// >
func (inter *Interpreter) evalGtExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left > right, nil
		case float64:
			return float64(left) > right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left > float64(right), nil
		case float64:
			return float64(left) > right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case string:
		switch right := rightValue.(type) {
		case string:
			return strings.Compare(left, right) > 0, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

// <
func (inter *Interpreter) evalLeExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left < right, nil
		case float64:
			return float64(left) < right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left < float64(right), nil
		case float64:
			return float64(left) < right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case string:
		switch right := rightValue.(type) {
		case string:
			return strings.Compare(left, right) < 0, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

// >=
func (inter *Interpreter) evalGteExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left >= right, nil
		case float64:
			return float64(left) >= right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left >= float64(right), nil
		case float64:
			return float64(left) >= right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case string:
		switch right := rightValue.(type) {
		case string:
			return strings.Compare(left, right) >= 0, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

// <=
func (inter *Interpreter) evalLeeExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left <= right, nil
		case float64:
			return float64(left) <= right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left <= float64(right), nil
		case float64:
			return float64(left) <= right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case string:
		switch right := rightValue.(type) {
		case string:
			return strings.Compare(left, right) <= 0, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

func (inter *Interpreter) evalEQExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	switch left := leftValue.(type) {
	case int64:
		switch right := rightValue.(type) {
		case int64:
			return left == right, nil
		case float64:
			return float64(left) == right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case float64:
		switch right := rightValue.(type) {
		case int64:
			return left == float64(right), nil
		case float64:
			return float64(left) == right, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	case string:
		switch right := rightValue.(type) {
		case string:
			return strings.Compare(left, right) == 0, nil
		default:
			return nil, fmt.Errorf("参数类型错误")
		}
	default:
		return nil, fmt.Errorf("参数类型错误")
	}
}

// ||
func (inter *Interpreter) evalLogicORExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	left, ok := leftValue.(bool)
	if !ok {
		return nil, fmt.Errorf("参数类型错误 期待bool值")
	}
	right, ok := rightValue.(bool)
	if !ok {
		return nil, fmt.Errorf("参数类型错误 期待bool值")
	}
	return left || right, nil
}

// &&
func (inter *Interpreter) evalLogicANDExpression(express *ast.InfixExpression) (interface{}, error) {
	leftValue, err := inter.evalExpress(express.Left)
	if err != nil {
		return nil, err
	}
	rightValue, err := inter.evalExpress(express.Right)
	if err != nil {
		return nil, err
	}
	left, ok := leftValue.(bool)
	if !ok {
		return nil, fmt.Errorf("参数类型错误 期待bool值")
	}
	right, ok := rightValue.(bool)
	if !ok {
		return nil, fmt.Errorf("参数类型错误 期待bool值")
	}
	return left && right, nil
}
