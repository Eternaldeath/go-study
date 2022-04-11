package main

import (
	"log"
	"net/http"
	"writing_web_application/server"
)

func main() {
	// 测试数据
	// p1 := &Page{"TestPage", []byte("这是测试数据")}
	// p1.savePage()

	// p2, err := loadPage("TestPage")
	// if err != nil {
	// 	fmt.Println("loadPage err: ", err)
	// }

	// fmt.Println(string(p2.Body))

	http.HandleFunc("/", server.RootHandler)
	// 注意，这里必须是“/personalIntro/”，像“/personalIntro”执行的是“/”的路径
	http.HandleFunc("/view/", server.MakeHandler(server.ViewHandler))
	http.HandleFunc("/edit/", server.MakeHandler(server.EditHandler))
	http.HandleFunc("/save/", server.MakeHandler(server.SaveHandler))

	// 启动监听本地 8080 端口号
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("http.ListenAndServe err: ", err)
	}
}
