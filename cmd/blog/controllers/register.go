package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_mod-project/cmd/blog/models"
	"go_mod-project/cmd/blog/utils"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Get() {
	r.TplName = "register.html"
}

func (r *RegisterController) Post() {
	fmt.Println(r)
	username := r.GetString("username")
	password := r.GetString("password")
	repassword := r.GetString("repassword")
	fmt.Println(username, password, repassword)

	// 注册之前判断是否已被注册
	id := models.QueryUserWithUsername(username)
	fmt.Println("id: ", id)
	if id > 0 {
		r.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已存在"}
		r.ServeJSON()
		return
	}
	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = utils.MD5(password)
	fmt.Println("md5后: ", password)
	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		r.Data["json"] = map[string]interface{}{"code": 0, "message": "注册失败"}
	} else {
		r.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	}
	r.ServeJSON()
}
