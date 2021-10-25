/*
 * @Author: your name
 * @Date: 2021-10-25 10:04:48
 * @LastEditTime: 2021-10-25 10:07:37
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \go-study\控制结构\测试多返回值函数的错误.go
 * 本题主要是感受函数存在错误返回值 error ，如何去接收 error
 * 以及利用 error 做进一步的处理
 */
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "ABC"
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// 感受 error
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
