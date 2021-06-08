package eval

import (
	"fmt"
	"panda/ast"
)

func (inter *Interpreter) evalAssginStatement(stat *ast.AssginStatement) (interface{}, error) {
	// 判断变量是否存在
	v, s := inter.scopeManager.GetValue(stat.Name.Value)
	if s == nil {
		return nil, fmt.Errorf("变量%s未定义", stat.Name.Value)
	}

	assginValue, err := inter.evalExpress(stat.Value)
	if err != nil {
		return nil, err
	}
	if stat.Index != nil {
		index, err := inter.evalExpress(stat.Index)
		if err != nil {
			return nil, err
		}
		if v == nil {
			return nil, fmt.Errorf("变量%s必须是数组或者map才能索引", stat.Name.Value)
		}
		// 说明是索引赋值
		// 返回值 应该只能是数组和map
		switch realValue := v.(type) {
		case map[interface{}]interface{}:
			realValue[index] = assginValue
			v = realValue
		case []interface{}:
			i, ok := index.(int64)
			if !ok {
				return nil, fmt.Errorf("数组只容许数字索引")
			}
			realValue[i] = assginValue
			v = realValue
		default:
			return nil, fmt.Errorf("当前类型不容许进行索引赋值")
		}
	} else {
		v = assginValue
	}

	ok := inter.scopeManager.SetValue(stat.Name.Value, v, false)
	if !ok {
		return nil, fmt.Errorf("变量%s不存在", stat.Name.Value)
	}
	return nil, nil
}
