# tutor

## 多模块工作区
[实例](https://gocn.vip/topics/qQP9Gvcyw3#%E5%A4%9A%E6%A8%A1%E5%9D%97%E5%B7%A5%E4%BD%9C%E5%8C%BA)

## 单元测试规则

- 单元测试文件名必须以 xxx_test.go 命名
- 方法必须是 TestXxx 开头，建议风格保持一致（驼峰或者下划线）
- 方法参数必须 t *testing.T
- 测试文件和被测试文件必须在一个包中
### 例子
```go
//add.go
package calc

import (
	"fmt"
	"math"
	"reflect"
)

func Add(a int, b int) int {
	return a + b
}

//add_test.go
package calc

import "testing"

func TestAdd(t *testing.T) {
	println(Add(3, 5))
}

 

```
## 可见性规则
Go语言没有像其它语言一样有`public、protected、private`等访问控制修饰符，它是通过字母大小写来控制可见性的，如果定义的常量、变量、类型、接口、结构、函数等的名称是**大写字母开头**表示能被其它包访问或调用（相当于public），非大写开头就只能在包内使用（相当于private，变量或常量也可以下划线开头）