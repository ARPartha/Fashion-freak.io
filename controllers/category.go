package controllers

import beego "github.com/beego/beego/v2/server/web"

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {

	c.TplName = "category.tpl"
}
