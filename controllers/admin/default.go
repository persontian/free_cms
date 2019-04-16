package admin

import (
	"fmt"
	"free_cms/models"
	"free_cms/pkg/d"
	"runtime"
	"strconv"
)

type MainController struct {
	BaseController
}

func (c *MainController) Admin() {
	c.TplName = "admin/default/index.html"
}

func (c *MainController) Main() {
	if c.Ctx.Input.IsAjax() {
		var system = make(map[string]string)
		system["server"] = runtime.GOARCH + " " + runtime.GOOS //系统
		system["version"] = runtime.Version() //go版本
		system["numcpu"] = strconv.Itoa(runtime.NumCPU()) //逻辑cpu
		system["numgoroutine"] = strconv.Itoa(runtime.NumGoroutine())//当前go携程数
		fmt.Println(models.GetMysqlMsg())
		c.Data["json"] = d.ReturnJson(200, "ok", system)
		c.ServeJSON()
	}
	c.TplName = "admin/default/main.html"
}
