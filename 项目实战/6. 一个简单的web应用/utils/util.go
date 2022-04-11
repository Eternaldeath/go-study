package utils

import (
	"html/template"
	"net/http"
	"writing_web_application/model"
)

// must 方法包裹了一层 panic，在这里是合适的
var templates = template.Must(template.ParseFiles("./views/edit.html", "./views/view.html"))

/**
 * @description: 渲染 html 页面
 * @param {http.ResponseWriter} w
 * @param {string} tmpl
 * @param {*model.Page} p
 * @return {*}
 */
func RenderTemplate(w http.ResponseWriter, tmpl string, p *model.Page) {
	// t, err := template.ParseFiles("./views/" + tmpl + ".html")
	// http.StatusInternalServerError = 500
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	// err = t.Execute(w, p)

	// ExecuteTemplate 会把内容 p 渲染到第二个参数指定的视图页面
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
