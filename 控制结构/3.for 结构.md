# for 结构

## 问题一

1. 使用 for 结构创建一个简单的循环。要求循环 15 次然后使用 fmt 包来打印计数器的值
2. 使用 goto 语句重写循环，要求不能使用 for 关键字

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(i)
	}

	fmt.Println("------分界线------")

	i := 0
START:
	fmt.Printf("%d\n", i)
	i++
	if i < 15 {
		goto START
	}
}

```

```bash
0
1
2
3
4
5
6
7
8
9
10
11
12
13
14
------分界线------
0
1
2
3
4
5
6
7
8
9
10
11
12
13
14
```

## 问题二

创建一个程序，要求能够打印类似下面的结果（直到每行 25 个字符时为止）

```go
G
GG
GGG
GGGG
GGGGG
GGGGGG
```

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 25; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("G")
		}
		fmt.Printf("\n")
	}
}
```

```bash
G
GG
GGG
GGGG
GGGGG
GGGGGG
GGGGGGG
GGGGGGGG
GGGGGGGGG
GGGGGGGGGG
GGGGGGGGGGG
GGGGGGGGGGGG
GGGGGGGGGGGGG
GGGGGGGGGGGGGG
GGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGGGGGGGGG
GGGGGGGGGGGGGGGGGGGGGGGG
```

## 问题三

写一个从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。遇到 5 的倍数时，打印 Buzz 而不是相应的数字。对于同时为 3 和 5 的倍数的数，打印 FizzBuzz。

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%s ", "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Printf("%s ", "Fizz")
		} else if i%5 == 0 {
			fmt.Printf("%s ", "Buzz")
		} else {
			fmt.Printf("%d ", i)
		}
	}
}

```

```bash
FizzBuzz 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz ...
```

## 问题四

使用 `*` 符号打印宽为 20，高为 10 的矩形

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}
}
```

```bash
********************
********************
********************
********************
********************
********************
********************
********************
********************
********************
```

## 问题五

以下程序的输出结果

```go
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}
}
```

```bash
0 0 0 0 0
```



