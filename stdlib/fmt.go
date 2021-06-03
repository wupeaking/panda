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
