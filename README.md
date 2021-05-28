# panda
学习编译原理 一步一步实现一个脚本语言

### 目前已经支持的表达式和语句

- 变量声明和赋值语句
- 四则运算和逻辑运算
- 匿名函数
- return语句
- 函数调用

### 示例
```
## in:
    var a = 1 +2*3;
    a;
    a = a+12;
    a;
    var b = 1+a;
    b*a+1+2;
    var calla = function(a, b, c, d) {
        return a+b+c+d;
    };
    var callb = function() {
        return 1+2+3;
    };
    var callc = function() {
        return function(a, b,c) {
            return a*b*c;
        };
    };

    calla(1,1,1,1);
    callb();
    callc()(2,2,3);
## out:
7
19
383
4
6
12
```



