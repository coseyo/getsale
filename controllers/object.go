package controllers

import (
	"getsale/models"

	"github.com/astaxie/beego"
	"github.com/robfig/cron"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

func init() {
	//	crawl()
	crontab()
}

func crontab() {
	c := cron.New()
	c.AddFunc("0 0 8 * * *", func() { crawl() })
	c.Start()
}

func crawl() {
	obj := &models.Object{}
	obj.TargetUrl = beego.AppConfig.String("targetUrl")
	obj.MailFrom = beego.AppConfig.String("mailFrom")
	obj.MailTo = beego.AppConfig.String("mailTo")
	obj.MailHost = beego.AppConfig.String("mailHost")
	obj.MailUser = beego.AppConfig.String("mailUser")
	obj.MailPassword = beego.AppConfig.String("mailPassword")
	obj.MailPort, _ = beego.AppConfig.Int("mailPort")
	obj.PageLimit, _ = beego.AppConfig.Int("pageLimit")
	obj.Go()
}

// @Title create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [get]
func (o *ObjectController) Get() {
	go crawl()
	o.Data["json"] = "开始抓取内容，请稍后查看邮箱"
	o.ServeJson()
}
