package controllers

import (
	"go_mod-project/cmd/blog/models"
)

type UpdateArticleController struct {
	BaseController
}

func (this *UpdateArticleController) Get()  {
	id, _ := this.GetInt("id")
	art := models.QueryArticleWithId(id)
	this.Data["Title"] = art.Title
	this.Data["Tags"] = art.Tags
	this.Data["Short"] = art.Short
	this.Data["Content"] = art.Content
	this.Data["Id"] = art.Id
	this.TplName = "write_article.html"
}

func (this *UpdateArticleController) Post()  {
	id, _ := this.GetInt("id")
	title := this.GetString("title")
	tags := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")
	//实例化model，修改数据库
	art := models.Article{id, title, tags, short, content, "", 0}
	_, err := models.UpdateArticle(art)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "更新失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	}
	this.ServeJSON()
}