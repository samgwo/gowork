package routers

import (
	"gowork/controllers/groups"
	"gowork/controllers/knowledges"
	"gowork/controllers/messages"
	"gowork/controllers/projects"
	"gowork/controllers/users"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &users.MainController{})

	//用户
	beego.Router("/user/manage", &users.ManageUserController{})
	beego.Router("/user/ajax/status", &users.AjaxStatusUserController{})
	beego.Router("/user/edit/:id", &users.EditUserController{})
	beego.Router("/user/add", &users.AddUserController{})
	beego.Router("/user/avatar", &users.AvatarUserController{})
	beego.Router("/user/ajax/search", &users.AjaxSearchUserController{})
	beego.Router("/user/show/:id", &users.ShowUserController{})
	beego.Router("/my/manage", &users.ShowUserController{})
	beego.Router("/user/profile", &users.EditUserProfileController{})
	beego.Router("/user/password", &users.EditUserPasswordController{})

	beego.Router("/user/permission/:id", &users.PermissionController{})

	beego.Router("/login", &users.LoginUserController{})
	beego.Router("/logout", &users.LogoutUserController{})

	//部门
	beego.Router("/department/manage", &users.ManageDepartmentController{})
	beego.Router("/department/ajax/status", &users.AjaxStatusDepartmentController{})
	beego.Router("/department/edit/:id", &users.EditDepartmentController{})
	beego.Router("/department/add", &users.AddDepartmentController{})

	//职位
	beego.Router("/position/manage", &users.ManagePositionController{})
	beego.Router("/position/ajax/status", &users.AjaxStatusPositionController{})
	beego.Router("/position/edit/:id", &users.EditPositionController{})
	beego.Router("/position/add", &users.AddPositionController{})

	//公告
	beego.Router("/notice/manage", &users.ManageNoticeController{})
	beego.Router("/notice/ajax/status", &users.AjaxStatusNoticeController{})
	beego.Router("/notice/ajax/delete", &users.AjaxDeleteNoticeController{})
	beego.Router("/notice/edit/:id", &users.EditNoticeController{})
	beego.Router("/notice/add", &users.AddNoticeController{})

	//项目
	beego.Router("/project/manage", &projects.ManageProjectController{})
	beego.Router("/project/ajax/status", &projects.AjaxStatusProjectController{})
	beego.Router("/project/edit/:id", &projects.EditProjectController{})
	beego.Router("/project/add", &projects.AddProjectController{})
	beego.Router("/project/:id", &projects.ShowProjectController{})

	beego.Router("/my/project", &projects.MyProjectController{})
	beego.Router("/project/chart/:id", &projects.ChartProjectController{})

	//项目成员
	beego.Router("/project/team/:id", &projects.TeamProjectController{})
	beego.Router("/team/ajax/delete", &projects.AjaxDeleteTeamProjectController{})
	beego.Router("/team/add/:id", &projects.AddTeamProjectController{})

	//项目需求
	beego.Router("/project/need/:id", &projects.NeedsProjectController{})
	beego.Router("/need/edit/:id", &projects.EditNeedsProjectController{})
	beego.Router("/need/add/:id", &projects.AddNeedsProjectController{})
	beego.Router("/need/show/:id", &projects.ShowNeedsProjectController{})
	beego.Router("/need/ajax/status", &projects.AjaxStatusNeedProjectController{})

	beego.Router("/my/need", &projects.MyNeedProjectController{})

	//项目任务
	beego.Router("/project/task/:id", &projects.TaskProjectController{})
	beego.Router("/task/edit/:id", &projects.EditTaskProjectController{})
	beego.Router("/task/add/:id", &projects.AddTaskProjectController{})
	beego.Router("/task/ajax/accept", &projects.AjaxAcceptTaskController{})
	beego.Router("/task/ajax/status", &projects.AjaxStatusTaskController{})
	beego.Router("/task/ajax/delete", &projects.DeleteTaskProjectController{})
	beego.Router("/task/show/:id", &projects.ShowTaskProjectController{})
	beego.Router("/project/taskbatch/:id", &projects.TaskBatchProjectController{})
	beego.Router("/task/clone/:id", &projects.CloneTaskProjectController{})

	beego.Router("/my/task", &projects.MyTaskProjectController{})

	//项目测试Bug
	beego.Router("/project/test/:id", &projects.TestProjectController{})
	beego.Router("/test/edit/:id", &projects.EditTestProjectController{})
	beego.Router("/test/add/:id", &projects.AddTestProjectController{})
	beego.Router("/test/ajax/accept", &projects.AjaxAcceptTestController{})
	beego.Router("/test/ajax/status", &projects.AjaxStatusTestController{})
	beego.Router("/test/ajax/delete", &projects.DeleteTestProjectController{})
	beego.Router("/test/show/:id", &projects.ShowTestProjectController{})

	beego.Router("/my/test", &projects.MyTestProjectController{})

	//项目文档
	beego.Router("/project/doc/:id", &projects.DocProjectController{})
	beego.Router("/doc/ajax/delete", &projects.AjaxDeleteDocPorjectController{})
	beego.Router("/doc/add/:id", &projects.FormDocProjectController{})
	beego.Router("/doc/edit/:id", &projects.FormDocProjectController{})
	beego.Router("/doc/show/:id", &projects.ShowDocProjectController{})

	//项目版本
	beego.Router("/project/version/:id", &projects.VersionProjectController{})
	beego.Router("/version/ajax/delete", &projects.AjaxDeleteVersionPorjectController{})
	beego.Router("/version/add/:id", &projects.FormVersionProjectController{})
	beego.Router("/version/edit/:id", &projects.FormVersionProjectController{})
	beego.Router("/version/show/:id", &projects.ShowVersionProjectController{})

	//知识分享
	beego.Router("/knowledge/manage", &knowledges.ManageKnowledgeController{})
	beego.Router("/knowledge/add", &knowledges.AddKnowledgeController{})
	beego.Router("/knowledge/edit/:id", &knowledges.EditKnowledgeController{})
	beego.Router("/knowledge/:id", &knowledges.ShowKnowledgeController{})
	beego.Router("/knowledge/comment/add", &knowledges.AddCommentController{})
	beego.Router("/knowledge/ajax/laud", &knowledges.AjaxLaudController{})
	beego.Router("/knowledge/ajax/delete", &knowledges.AjaxDeleteKnowledgeController{})

	//beego.Router("/task/ajax/status", &projects.AjaxAcceptTaskController{}, "*:AddPost")

	//消息
	beego.Router("/message/manage", &messages.ManageMessageController{})
	beego.Router("/message/ajax/delete", &messages.AjaxDeleteMessageController{})
	beego.Router("/message/ajax/status", &messages.AjaxStatusMessageController{})

	//组
	beego.Router("/group/manage", &groups.ManageGroupController{})
	beego.Router("/group/ajax/delete", &groups.AjaxDeleteGroupController{})
	beego.Router("/group/add", &groups.FormGroupController{})
	beego.Router("/group/edit/:id", &groups.FormGroupController{})
	//组成员
	beego.Router("/group/user/:id", &groups.ManageGroupUserController{})
	beego.Router("/group/user/add/:id", &groups.FormGroupUserController{})
	beego.Router("/group/user/ajax/delete", &groups.AjaxDeleteGroupUserController{})
	//组权限
	beego.Router("/group/permission/:id", &groups.ManageGroupPermissionController{})
	//beego.Router("/group/permission/add/:id", &groups.FormGroupPermissionController{})
	beego.Router("/group/permission/ajax/delete", &groups.AjaxDeleteGroupPermissionController{})

	//权限
	beego.Router("/permission/manage", &groups.ManagePermissionController{})
	beego.Router("/permission/ajax/delete", &groups.AjaxDeletePermissionController{})
	beego.Router("/permission/add", &groups.FormPermissionController{})
	beego.Router("/permission/edit/:id", &groups.FormPermissionController{})

}
