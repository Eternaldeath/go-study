# map排序问题

## 将map键值对对调

这里我直接用了《Go 入门指南》里面的代码

```
package main
import (
    "fmt"
)

var (
    barVal = map[string]int{"alpha": 34, "bravo": 56, "charlie": 23,
                            "delta": 87, "echo": 56, "foxtrot": 12,
                            "golf": 34, "hotel": 16, "indio": 87,
                            "juliet": 65, "kili": 43, "lima": 98}
)

func main() {
    invMap := make(map[int]string, len(barVal))
    for k, v := range barVal {
        invMap[v] = k
    }
    fmt.Println("inverted:")
    for k, v := range invMap {
        fmt.Printf("Key: %v, Value: %v / ", k, v)
    }
}
```

```ba
inverted:
Key: 34, Value: golf / Key: 23, Value: charlie / Key: 16, Value: hotel / Key: 87, Value: delta / Key: 98, Value: lima / Key: 12, Value: foxtrot / Key: 43, Value: kili / Key: 56, Value: bravo / Key: 65, Value: juliet /
```

