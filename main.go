package main

import (
	"html/template"
	"net/http"
	_ "gowork/initial"
	_ "gowork/routers"
	_ "gowork/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"

)

func main() {
	logs.SetLogger(logs.AdapterFile,`{"filename":"logs/gowork.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.ErrorHandler("404", page_not_found)
	beego.ErrorHandler("401", page_note_permission)
	beego.Run()
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("userLogin").(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}

func page_note_permission(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("401.tpl").ParseFiles("views/401.tpl")
	data := make(map[string]interface{})
	t.Execute(rw, data)
}
