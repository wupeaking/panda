package stdlib

import "fmt"

type BuildFunc func([]interface{}) (interface{}, bool, error)

func println(args []interface{}) (interface{}, bool, error) {
	for _, v := range args {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
	return nil, false, nil
}

func length(args []interface{}) (interface{}, bool, error) {
	if len(args) != 0 {
		arg := args[0]
		switch v := arg.(type) {
		case []interface{}:
			return int64(len(v)), true, nil
		default:
			return 0, false, fmt.Errorf("参数类型错误 len函数只支持数组类型")
		}
	}
	return int64(0), true, nil
}
