package routers

import (
	"github.com/wzh700/beego_study/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
