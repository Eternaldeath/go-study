

# 题目参考

《Mastering go》

# 题目

使用 go 实现一个简单的 kv 存储，要满足以下四个功能：添加新元素，基于 key 值删除元素，根据 key 值查找元素，基于 key 修改元素，输入 Stop 停止程序。

# 解答

1. 我们设计一个关于种类的 kv 存储，key 为种类（如，人类，狗类等），value 为实例（如，人类下有张三，11 岁）
2. 引入相关包，设计这个结构体。使用 map 实现这个 kv 存储

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 此处定义 value 为结构体
type MyValue struct {
	Name string
	Age  int
}

// k-v 存储在这里用 map 类型实现
var Data = make(map[string]MyValue)
```
3. 增删查改的函数
```go
// 新增
func Add(k string, v MyValue) bool {
	if k == "" {
		fmt.Println("键名不能为空")
		return false
	}

	if LookUp(k) == nil {
		Data[k] = v
		fmt.Println("新增成功！")
		return true
	}

	fmt.Println("新增失败！")
	return false
}

// 删除
func Delete(k string) bool {
	if LookUp(k) != nil {
		delete(Data, k)
		fmt.Println("删除成功！")
		return true
	} else {
		fmt.Println("没有这个键名！")
		return false
	}
}

// 查询
func LookUp(k string) *MyValue {
	_, ok := Data[k]
	if ok {
		n := Data[k]
		return &n
	} else {
		return nil
	}
}

// 更改
func Change(k string, v MyValue) bool {
	Data[k] = v
	return true
}

// 打印全部信息
func Print() {
	for k, v := range Data {
		fmt.Printf("key: %s, value: %v\n", k, v)
	}
}
```
4. 从控制台接收输入的命令
```go
// 持续接收标准输入
func GetInput() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// 获取输入的命令
		text := scanner.Text()
		// 将命令中间的空格去掉
		text = strings.TrimSpace(text)
		// 将这些命令字段存入 tokens 中
		// Fields 返回的是切片类型
		// tokens 里面存储的是命令，比如 Add human zhangsan 14
		// 可以注意到，所有 token 最长为 4 个元素（如上面的 Add）
		tokens := strings.Fields(text)

		// 通过切片长度决定命令
		switch len(tokens) {
		// 因为 tokens 最长为 4 个元素，所以对于一些命令不足 4 个的，我们通过 append 补全
		case 0:
			continue
		case 1:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 2:
			tokens = append(tokens, "")
			tokens = append(tokens, "")
		case 3:
			tokens = append(tokens, "")
		}

		// 根据 tokens 判断命令
		switch tokens[0] {
		case "Print":
			Print()
		case "Stop":
			return
		case "Delete":
			if !Delete(tokens[1]) {
				fmt.Println("删除失败")
			}
		case "Add":
			// 将第三个字符串字段（年龄）转换成 int 类型
			if tokenAge, err := strconv.Atoi(tokens[3]); err == nil {
				v := MyValue{tokens[2], tokenAge}
				if !Add(tokens[1], v) {
					fmt.Println("新增失败！")
				}
			}
		case "LookUp":
			n := LookUp(tokens[1])
			if n != nil {
				fmt.Printf("%v\n", n)
			}
		case "Change":
			if tokenAge, err := strconv.Atoi(tokens[3]); err == nil {
				v := MyValue{tokens[2], tokenAge}
				if !Change(tokens[1], v) {
					fmt.Println("更新失败")
				}
			}
		default:
			fmt.Println("未知命令，重新输入！")
		}

	}
}
```

5. 测试结果

```bash
go run .
Print
Add human zhangsan 11
新增成功！
Add human a 11
新增成功！
LookUp human
{a 11}
key: human, value: {a 11}
&{a 11}
PS D:\code\go_project\study\gobasic_study\gocombination_study> go run .
Add human zhangsan 12
新增成功！
LookUp human
&{zhangsan 12}
Add dog dahuang 2 
新增成功！
Print
key: human, value: {zhangsan 12}
key: dog, value: {dahuang 2}
Delete dog
删除成功！
Print
key: human, value: {zhangsan 12}
Stop
```

