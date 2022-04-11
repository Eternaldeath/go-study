/*
	定义页面结构
*/
package model

import "os"

type Page struct {
	Title string
	Body  []byte
}

/**
 * @description: 存储你的信息到 txt 文件中
 * @param {*}
 * @return {*}
 */
func (p *Page) SavePage() error {
	filename := p.Title + ".txt"
	// win 用户可以不用在意第三个参数
	// WriteFile 会根据 filename 将数据写入文件中
	// 如果不存在 filename.txt 则会自己创建
	return os.WriteFile(filename, p.Body, 0600)
}

/**
 * @description: 从 txt 读取数据
 * @param {string} title
 * @return {*}
 */
func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
