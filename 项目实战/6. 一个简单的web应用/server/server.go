package server

import (
	"fmt"
	"net/http"
	"writing_web_application/model"
	"writing_web_application/utils"
)

/**
 * @description: / 路由
 * @param {http.ResponseWriter} w: 存放从服务器发往客户机的数据
 * @param {*http.Request} r: 存放从客户机发往服务器的请求
 * @return {*}
 */
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "你好，你正在访问%s", r.URL.Path[1:])
}

/**
 * @description: 个人介绍处理函数
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {*}
 */
// func personalIntroHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.URL.Path[len("/personalIntro/"):]
// 	p, err := loadPage(title)
// 	if err != nil {
// 		fmt.Println("loadPage err: ", err)
// 	}
// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }

/**
 * @description: 编辑页面
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {*}
 */
func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := model.LoadPage(title)
	if err != nil {
		p = &model.Page{Title: title}
	}
	// fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	// 	"<form action=\"/save/%s\" method=\"POST\">"+
	// 	"<textarea name=\"body\">%s</textarea><br>"+
	// 	"<input type=\"submit\" value=\"Save\">"+
	// 	"</form>",
	// 	p.Title, p.Title, p.Body)

	// 这里我们用 template 替换硬编码
	// t, _ := template.ParseFiles("./views/edit.html")
	// t.Execute(w, p)

	// 这里我们甚至把渲染模板封装成了一个函数
	utils.RenderTemplate(w, "edit", p)
}

/**
 * @description: 展示页处理函数
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 * @return {*}
 */
func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := model.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	utils.RenderTemplate(w, "view", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// FormValue 可以获取？后面的参数（url，a 标签等）
	body := r.FormValue("body")
	// 像 model 中装填数据
	p := &model.Page{title, []byte(body)}
	// 存储数据
	err := p.SavePage()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 跳转到 view 视图
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

/**
 * @description: 统一处理函数
 * @param {*} http
 * @param {*} http
 * @param {*} string
 * @return {*}
 */
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := utils.ValidPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
