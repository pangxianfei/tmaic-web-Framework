package groups

import (
	"github.com/pangxianfei/framework/route"
	"tmaic/app/http/controllers"
)

type UserGroup struct {
	UserController controllers.User
}

func (ug *UserGroup) Group(group route.Grouper) {
	//group.GET("/info/userId", ug.UserController.Info).Can(policies.NewUserPolicy(), policy.ActionView)
	group.GET("/info/:userId", ug.UserController.Info).Name("uesrinfo")




	group.GET("/update", ug.UserController.Update)
	group.GET("/delete", ug.UserController.Delete)
	group.GET("/delete-transaction", ug.UserController.DeleteTransaction)
	group.GET("/logout", ug.UserController.LogOut)
	group.GET("/restore", ug.UserController.Restore)
}
