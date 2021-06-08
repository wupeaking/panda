# panda
一个玩具性质的脚本解释器

### 目前已经支持的语法特性

- 变量声明和赋值语句
- 四则运算和逻辑运算
- 函数声明
- 匿名函数
- return语句
- 函数调用
- 流程控制if语句
- for循环控制语句
- 数组和map类型

### 内建函数
- println
- len
- append

### 使用方法
```shell
NAME:
   Panda - 一个玩具性质的脚本解释器

USAGE:
   main [global options] command [command options] 源文件路径

VERSION:
   0.0.1

COMMANDS:
   ast, ast  打印出抽象语法树
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

### 执行示例脚本
```
> panda code/example.pd
```
### 代码示例
```
    // 变量申明和赋值
    var a = 1 +2*3;
    // 答应变量
    println(a);
    // 变量类型自动转换  int->float
    a = a+12.1;
    println(a);

    // 匿名函数表达式
    var calla = function(a, b, c, d) {
        return a+b+c+d;
    };

    var callb = function() {
        return 1+2+3;
    };

    // 匿名函数返回值为匿名函数
    var callc = function() {
        return function(a, b,c) {
            return a*b*c;
        };
    };

    // 匿名函数调用
    println(calla(1,1,1,1));
    callb();
    println(a+callc()(2,2,3));

    // 函数声明
    function add(left, right) {
        return left+right;
    }
    // 函数调用
    add(1, 2.1*3);

    // 数组和map赋值
    var a = [1, 2, [3, 1, 2]];
    var b = {
        "name": "wpx",
        "age" : 10+20,
        "sex": true
    };

    // 遍历数组元素
    for(var i = 0; i < len(a); i=i+1) {
        if(i==2) {
            println("a=", a[i][1]);
        }else{
            println("a=", a[i]);
        }
    }

    // 数组和map的修改
    a[0] = 333;
    b["age"] = 31;
    b["address"] = "address";

```



