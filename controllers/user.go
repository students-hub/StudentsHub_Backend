package controllers

import (
	"StudentsHub_Backend/models"
	"encoding/json"
	"errors"
	"strings"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.UserID
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) AddUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		logs.Info("添加失败，原因是:", err)
		u.Ctx.WriteString("添加失败，原因是:" + err.Error())
		return
	}
	u.Data["json"] = map[string]int{"UserID": user.UserID}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]

func (u *UserController) Put() {

}

// @Title Login
// @Description Logs user into the system
// @Param	UserID		query 	int		true		"The username for login"
// @Param	Password	query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	UserID, err := u.GetInt("UserID")
	if err != nil {
		logs.Info("获取用户ID失败，原因是:", err)
		u.Ctx.WriteString("获取用户ID失败，原因是:" + err.Error())
		return
	}
	Password := u.GetString("Password")
	isLogin, _ := Authenticate(UserID, Password)
	if isLogin {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "user not exist or password invalid"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

func Authenticate(UserID int, Password string) (bool, error) {
	o := orm.NewOrm()
	user := models.User{UserID: UserID}
	err := o.QueryTable("user").Filter("UserID", UserID).Limit(1).One(&user)
	if err != nil {
		logs.Info("查询失败，原因是:", err)
		return false, errors.New(strings.Join([]string{"查询失败，原因是", err.Error()}, ""))
	}
	if user.Password == Password {
		return true, nil
	} else {
		return false, errors.New("密码错误")
	}
}
