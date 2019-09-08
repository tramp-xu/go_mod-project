package controllers

import (
	"fmt"
	"go_mod-project/cmd/blog/models"
)

type AlbumController struct {
	BaseController
}

func (this *AlbumController) Get() {
	albums, err := models.FindAllAlbums()
	if err != nil {
		fmt.Println(err)
	}
	this.Data["Album"] = albums
	this.TplName = "album.html"
}
