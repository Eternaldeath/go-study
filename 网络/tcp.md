# tcp 服务器

## 问题一

实现一个简单的 tcp 服务器，要求至少有以下几个功能：

1. 客户端发送数据，服务器接收数据并展示
2. 当客户端发送 “Q” 时，客户端断开链接
3. 当客户端发送 “SH” 时，服务器关闭
4. 服务器记录交互的客户端名字，当客户端发送 “WHO”时，展示哪些客户端在线，哪些服务器离线
5. 自定义 ```checkError``` 函数，包装错误检查

这里直接采用 《Go 入门指南》中的代码

```go
// server
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Map of the clients: contains: clientname - 1 (active) / 0 - (inactive)
var mapUsers map[string]int

func main() {
	var listener net.Listener
	var error error
	var conn net.Conn
	mapUsers = make(map[string]int)

	fmt.Println("Starting the server ...")

	// create listener:
	listener, error = net.Listen("tcp", "localhost:50000")
	checkError(error)
	// listen and accept connections from clients:
	for {
		conn, error = listener.Accept()
		checkError(error)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	var buf []byte
	var error error

	for {
		buf = make([]byte, 512)
		_, error = conn.Read(buf)
		checkError(error)
		input := string(buf)
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down.")
			os.Exit(0)
		}
		// op commando WHO:  write out mapUsers
		if strings.Contains(input, ": WHO") {
			DisplayList()
		}
		// extract clientname:
		ix := strings.Index(input, "says")
		clName := input[0 : ix-1]
		//fmt.Printf("The clientname  is ---%s---\n", string(clName))
		// set clientname active in mapUsers:
		mapUsers[string(clName)] = 1
		fmt.Printf("Received data: --%v--", string(buf))
	}
}

// advantage: code is cleaner,
// disadvantage:  the server process has to stop at any error:
//                a simple return continues in the function where we came from!
func checkError(error error) {
	if error != nil {
		panic("Error: " + error.Error()) // terminate program
	}
}

func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}
```

```go
// client
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	var conn net.Conn
	var error error
	var inputReader *bufio.Reader
	var input string
	var clientName string

	// maak connectie met de server:
	conn, error = net.Dial("tcp", "localhost:50000")
	checkError(error)

	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ = inputReader.ReadString('\n')
	// fmt.Printf("CLIENTNAME %s",clientName)
	trimmedClient := strings.Trim(clientName, "\r\n") // "\r\n" voor Windows, "\n" voor Linux

	for {
		fmt.Println("What to send to the server? Type Q to quit. Type SH to shutdown server.")
		input, _ = inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\r\n")
		// fmt.Printf("input:--%s--",input)
		// fmt.Printf("trimmedInput:--%s--",trimmedInput)
		if trimmedInput == "Q" {
			return
		}
		_, error = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
		checkError(error)
	}
}

func checkError(error error) {
	if error != nil {
		panic("Error: " + error.Error()) // terminate program
	}
}
```

