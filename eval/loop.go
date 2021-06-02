package eval

import (
	"panda/ast"
)

func (inter *Interpreter) evalForStatement(forStatement *ast.ForStatement) (interface{}, bool, error) {
	inter.scopeManager.Push(IFScope)
	defer inter.scopeManager.Pop()
	_, err := inter.evalStatement(forStatement.FirstCondition.(ast.Statement))
	if err != nil {
		return nil, false, err
	}
	for {
		v, err := inter.evalExpress(forStatement.MidCondition.(ast.Expression))
		if err != nil {
			return nil, false, err
		}
		var isContinue bool
		switch s := v.(type) {
		case int64:
			isContinue = s > 0
		case float64:
			isContinue = s > 0.0
		case bool:
			isContinue = s
		}
		if !isContinue {
			break
		}
		ret, hasRet, err := inter.evalASTNodes(forStatement.LoopBody.Statements)
		if err != nil {
			return nil, false, err
		}
		if hasRet {
			return ret, true, nil
		}

		// break 语句
		if ret != nil {
			if _, ok := ret.(*ast.BlockStatement); ok {
				return nil, false, nil
			}
		}

		// 继续执行
		_, err = inter.evalStatement(forStatement.LastCondition.(ast.Statement))
		if err != nil {
			return nil, false, err
		}
	}
	return nil, false, nil
}
