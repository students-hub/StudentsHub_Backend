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
// @router /sign-up [post]
func (u *UserController) AddUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	o := orm.NewOrm()
	o.Begin()
	_, err := o.Insert(&user)
	if err != nil {
		logs.Info("添加失败，原因是:", err)
		u.Ctx.WriteString("添加失败，原因是:" + err.Error())
		o.Rollback()
		return
	} else {
		o.Commit()
	}
	u.Data["json"] = map[string]int{"UserID": user.UserID} //返回UserID, 给予用户提示
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	body		body 	models.User	true		"Old info of the user"
// @Param	NewPassword		body 	string	true		"New Password"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /updatepswd [put]
func (u *UserController) UpdatePassword() {
	var user models.User
	//var u models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	m := make(map[string]string)
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)
	o := orm.NewOrm()
	_user := &models.User{UserName: user.UserName}
	err := o.Read(_user, "UserName")
	if err == nil {

		//验证密码
		if m["Password"] != _user.Password {
			logs.Info("密码错误!")
			u.Data["json"] = "原密码错误!"
			u.ServeJSON()
			return
		}

		_user.Password = m["NewPassword"]
		if _, err := o.Update(_user, "Password", "UpdateAt"); err != nil {
			logs.Info(err.Error())
		}
	} else {
		logs.Info(err.Error())
	}
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
