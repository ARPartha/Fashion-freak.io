package routers

import (
	"Fashion-Freaks/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/category", &controllers.CategoryController{})
}
