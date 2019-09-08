package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_mod-project/cmd/blog/models"
)

type DeleteArticleController struct {
	beego.Controller
}

func (this *DeleteArticleController) Get() {
	id, _ := this.GetInt("id")
	_, err := models.DeleteArticle(id)
	if err != nil {
		fmt.Println(err)
	}
	this.Redirect("/", 302)
}
