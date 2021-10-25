/*
 * @Author: your name
 * @Date: 2021-10-22 10:54:01
 * @LastEditTime: 2021-10-22 10:55:30
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \go-study\变量\init_study.go
 */
// 推断以下程序的输出，并解释你的答案，然后编译并执行它们
package varstudy

var a string

func main() {
	a = "G"
	print(a)
	f1()
	// GOG
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}
