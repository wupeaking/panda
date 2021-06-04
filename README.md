# panda
学习编译原理 一步一步实现一个脚本语言

### 目前已经支持的表达式和语句

- 变量声明和赋值语句
- 四则运算和逻辑运算
- 函数声明
- 匿名函数
- return语句
- 函数调用
- 流程控制if语句
- for循环控制语句
- 数组索引

### 示例
```
## in:
    var a = 1 +2*3;
    println(a);
    a = a+12.1;
    println(a);
    var b = 1+a;
    println(b*a+1+2);
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
    println(calla(1,1,1,1));
    callb();
    callc()(2,2,3);
    println(a+callc()(2,2,3));
    function add(left, right) {
        return left+right;
    }
    add(1, 2.1*3);
    var b = 1;
    for(var a = 1; a<10; a=a+1){
        b = b +a;
        println(b);
        if (a>=5){
            break;
        }
    }
    var a = [1, 2, [3, 1, 2]];
    println(a[2][1]);
## out:
7 
19.1 
386.91 
4 
31.1 
2 
4 
7 
11 
16
1
```



