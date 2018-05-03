package controllers

import (
	"fmt"
	. "gowork/models/groups"
	. "gowork/models/messages"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
	IsLogin      bool
	UserUserId   int64
	UserUsername string
	UserAvatar   string
}

func (this *BaseController) Prepare() {
	userLogin := this.GetSession("userLogin")
	if userLogin == nil {
		this.IsLogin = false
	} else {
		this.IsLogin = true
		tmp := strings.Split((this.GetSession("userLogin")).(string), "||")

		userid, _ := strconv.ParseInt(tmp[0], 10, 64)
		longid := int64(userid)
		this.Data["LoginUserid"] = longid
		this.Data["LoginUsername"] = tmp[1]
		this.Data["LoginAvatar"] = tmp[2]

		this.UserUserId = longid
		this.UserUsername = tmp[1]
		this.UserAvatar = tmp[2]

		msgcondArr := make(map[string]string)
		msgcondArr["touserid"] = fmt.Sprintf("%d", longid)
		msgcondArr["view"] = "1"
		countTopMessage := CountMessages(msgcondArr)
		_, _, topMessages := ListMessages(msgcondArr, 1, 10)
		this.Data["topMessages"] = topMessages
		this.Data["countTopMessage"] = countTopMessage

		_, _, leftNav := ListGroupsUserPermission(this.GetSession("userGroupid").(string))
		this.Data["leftNav"] = leftNav
	}
	this.Data["IsLogin"] = this.IsLogin
}
