/*
	工具类：处理应用中名字的正则表达式
*/
package utils

import (
	"errors"
	"net/http"
	"regexp"
)

// 提供了 panic 的 compile
var ValidPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func GetTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	// FindStringSubmatch 验证我们的 r.URL.Path 是否满足 validPath 标准
	m := ValidPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}
