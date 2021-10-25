# switch

## 问题一

[问题来源](https://learnku.com/docs/the-way-to-go/switch-structure/3594)

问题：请说出下面代码片段输出的结果

```go
k := 6
switch k {
    case 4: fmt.Println("was <= 4"); fallthrough;
    case 5: fmt.Println("was <= 5"); fallthrough;
    case 6: fmt.Println("was <= 6"); fallthrough;
    case 7: fmt.Println("was <= 7"); fallthrough;
    case 8: fmt.Println("was <= 8"); fallthrough;
    default: fmt.Println("default case")
}
```

```bash
was <= 6
was <= 7
was <= 8
default case
```

## 问题二

[问题来源](https://learnku.com/docs/the-way-to-go/switch-structure/3594)

问题：写一个 Season 函数，要求接受一个代表月份的数字，然后返回所代表月份所在季节的名称（不用考虑月份的日期）

```go
package main

import (
	"fmt"
)

func main() {
	var month int

	fmt.Println("Please enter your month: ")
    // 从键盘读取输入的值
	fmt.Scanln(&month)

	switch month {
	case 3, 4, 5:
		fmt.Println("春季")
	case 6, 7, 8:
		fmt.Println("夏季")
	case 9, 10, 11:
		fmt.Println("秋季")
	case 12, 1, 2:
		fmt.Println("冬季")
	default:
		fmt.Println("未知季节")
	}
}

```

```bash
Please enter your month:
11
秋季
```

