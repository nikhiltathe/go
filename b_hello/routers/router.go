package routers

import (
	"github.com/nikhil5987/b_hello/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
