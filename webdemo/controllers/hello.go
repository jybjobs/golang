package controllers

import (
	"encoding/json"
	"webdemo/models"

	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (this *HelloController) Get() {
	this.Ctx.WriteString("hello")
}

func (this *HelloController) Post() {
	var err error
	var ob models.User
	if err = json.Unmarshal(this.Ctx.Input.RequestBody, &ob); err == nil {
		this.Data["json"] = ob
	} else {
		this.Data["json"] = err.Error()
	}
	this.ServeJSON()
}
