package controllers

import (
	"net/http"

	"github.com/pangxianfei/framework/config"
	"github.com/pangxianfei/framework/helpers/m"
	"github.com/pangxianfei/framework/helpers/toto"
	"github.com/pangxianfei/framework/http/controller"
	"github.com/pangxianfei/framework/http/middleware"
	"github.com/pangxianfei/framework/policy"
	"github.com/pangxianfei/framework/request"

	"tmaic/app/models"
	"tmaic/app/policies"
)

type User struct {
	controller.BaseController
}

func (*User) LogOut(c request.Context) {
	if err := middleware.Revoke(c, config.GetString("auth.sign_key")); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, toto.V{})
	return
}

func (u *User) Info(c request.Context) {
	if c.ScanUserWithJSON() {
		return
	}
	user := c.User().Value().(*models.User)

	if permit, _ := u.Authorize(c, policies.NewUserPolicy(), policy.ActionView); !permit {
		c.JSON(http.StatusForbidden, toto.V{"error": policy.UserNotPermitError{}.Error()})
		return
	}

	c.JSON(http.StatusOK, toto.V{"data": user})
	return
}

func (*User) AllUser(c request.Context) {

	return
}

func (*User) PaginateUser(c request.Context) {

	return
}

func (*User) Update(c request.Context) {
	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	if err := m.H().First(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.Error()})
		return
	}

	name := "t2222es123t"
	modifyUser := models.User{
		Name: &name,
	}
	if err := m.H().Save(&user, modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, toto.V{"data": user})
	return

	// m.Transaction(func() {
	// 	fmt.Println(id)
	// 	panic(123)
	// }, 3)
}
func (*User) Delete(c request.Context) {
	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	if err := m.H().Delete(&user, false); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, toto.V{"data": true})
	return
}
func (*User) DeleteTransaction(c request.Context) {
	defer func() { // handle transaction error
		if err := recover(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err.(error).Error()})
			return
		}
	}()

	var id uint
	id = 14
	user := models.User{
		ID: &id,
	}
	m.Transaction(func(h *m.Helper) {
		user.SetTX(h.DB()) // important
		if err := h.Delete(&user, false); err != nil {
			panic(err)
		}
	}, 1)

	c.JSON(http.StatusOK, toto.V{"data": true})
	return
}
func (*User) Restore(c request.Context) {
	var id uint
	id = 14
	modifyUser := models.User{
		ID: &id,
	}

	if err := m.H().Restore(&modifyUser); err != nil {
		c.JSON(http.StatusUnprocessableEntity, toto.V{"error": err})
		return
	}
	c.JSON(http.StatusOK, toto.V{"data": true})
	return
}