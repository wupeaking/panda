package stdlib

var StdlibFuncs = map[string]BuildFunc{}

func init() {
	StdlibFuncs["println"] = println
	StdlibFuncs["len"] = length
	StdlibFuncs["append"] = appendArr
}
