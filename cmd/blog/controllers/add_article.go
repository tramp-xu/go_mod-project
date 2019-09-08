package controllers

import (
	"fmt"
	"go_mod-project/cmd/blog/models"
	"time"
)

type AddArticleController struct {
	BaseController
}

/*
当访问/add路径的时候回触发AddArticleController的Get方法
响应的页面是通过TpName
*/

func (c *AddArticleController) Get()  {
	c.TplName = "write_article.html"
}

func (c *AddArticleController) Post()  {
	title := c.GetString("title")
	tags := c.GetString("tags")
	short := c.GetString("short")
	content := c.GetString("content")
	fmt.Printf("title: %s, tags: %s \n", title, tags)

	// 实例化数据
	art := models.Article{0, title, tags, short, content, "xuuu", time.Now().Unix()}

	_, err := models.AddArticle(art)
	var response map[string]interface{}
	if err == nil {
		response = map[string]interface{}{"code": 1, "message": "ok"}
	} else {
		response = map[string]interface{}{"code": 0, "message": "error"}
	}

	c.Data["json"] = response
	c.ServeJSON()
}
