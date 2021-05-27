package controllers

import (
	"StudentsHub_Backend/models"
	"encoding/json"
	"errors"
	"strings"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/prometheus/common/log"
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

// @Title Update Password
// @Description update the user on password
// @Param	body			body 	models.User	true		"Old info of the user"
// @Param	NewPassword		body 	string	true			"New Password"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /updatepswd [put]
func (u *UserController) UpdatePassword() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	m := make(map[string]string)
	json.Unmarshal(u.Ctx.Input.RequestBody, &m)
	o := orm.NewOrm()
	_user := &models.User{UserName: user.UserName}
	err := o.Read(_user, "UserName")
	if err == nil {

		//验证密码
		if user.Password != _user.Password {
			logs.Info("原密码错误!")
			u.Data["json"] = "原密码错误!"
			u.ServeJSON()
			return
		}

		//如果新、旧密码相同
		if _user.Password == m["NewPassword"] {
			logs.Info("新、旧密码不能相同!")
			u.Data["json"] = "新、旧密码不能相同!"
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

// @Title Update UserName
// @Description update the user on username
// @Param	OldName			query 	string	true		"Old name of the user"
// @Param	NewName			query 	string	true		"New name"
// @Success 200 {object} models.User
// @Failure 403 NewName has been used
// @router /updatename [put]
func (u *UserController) UpdateUsername() {
	OldName := u.GetString("OldName")
	NewName := u.GetString("NewName")
	/*
		if OldName == NewName {
			logs.Info("新、旧用户名不能相同!")
			u.Data["json"] = "新、旧用户名不能相同!"
			u.ServeJSON()
			return
		}
	*/
	o := orm.NewOrm()
	user := &models.User{UserName: OldName}  //原用户
	o.Read(user, "UserName")                 //err应该为nil
	_user := &models.User{UserName: NewName} //新用户
	err := o.Read(_user, "UserName")         //通过查询，判断有无重复，如果err不为nil，说明新用户名查不出来，那么就可以更改
	if err != nil {                          //没有重复
		user.UserName = NewName
		if _, err := o.Update(user, "UserName", "UpdateAt"); err != nil {
			logs.Info(err.Error())
		}
	} else {
		logs.Info("用户名已存在!")
		u.Data["json"] = "用户名已存在!"
		u.ServeJSON()
		return
	}
}

// @Title Update Role
// @Description update the user on Role
// @Param	UserName		query 	string	true		"Name of the user"
// @Param	NewRole			query 	string	true		"New Role of the user"
// @Success 200 {object} models.User
// @Failure 403 Update failed
// @router /updaterole [put]
func (u *UserController) UpdateRole() {
	UserName := u.GetString("UserName")
	NewRole := u.GetString("NewRole") //1 teacher; 2 assistant; 3 student

	o := orm.NewOrm()
	user := &models.User{UserName: UserName}
	err := o.Read(user, "UserName")
	if err == nil {

		if user.Role == NewRole {
			log.Info("新、旧角色不能相同!")
			u.Data["json"] = "新、旧角色不能相同!"
			u.ServeJSON()
			return
		}

		user.Role = NewRole
		if _, err := o.Update(user, "Role", "UpdateAt"); err != nil {
			logs.Info(err.Error())
		}
	} else {
		logs.Info("更新失败!")
		u.Data["json"] = "更新失败!"
		u.ServeJSON()
		return
	}
}

// @Title Login
// @Description Logs user into the system
// @Param	UserName	query 	string	true		"The username for login"
// @Param	Password	query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist or password invalid
// @router /login [get]
func (u *UserController) Login() {
	UserName := u.GetString("UserName")
	Password := u.GetString("Password")
	isLogin, _ := Authenticate(UserName, Password)
	if isLogin {
		u.Data["json"] = map[string]string{"Info": "login succeeded", "UserName": UserName} //返回UserName, 供以后改名、密码等等之用
	} else {
		u.Data["json"] = "user not exist or password invalid"
	}
	u.ServeJSON()
}

// @Title Logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout succeeded"
	u.ServeJSON()
}

// @Title Delete
// @Description Delete user
// @Param	body	body 	models.User	true		"User Info"
// @Success 200 {string} delete success
// @router /delete [put]
func (u *UserController) Delete() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	o := orm.NewOrm()
	_user := &models.User{UserName: user.UserName}
	err := o.Read(_user, "UserName")
	if err == nil {

		//验证密码
		if user.Password != _user.Password {
			logs.Info("密码错误!")
			u.Data["json"] = "密码错误!"
			u.ServeJSON()
			return
		}

		if _, err := o.Delete(_user); err != nil {
			logs.Info(err.Error())
			u.Data["json"] = "delete failed"
		} else {
			u.Data["json"] = "delete succeeded"
		}
		u.ServeJSON()
		return
	} else {
		logs.Info(err.Error())
		u.Data["json"] = "delete failed"
		u.ServeJSON()
		return
	}

}

func Authenticate(UserName string, Password string) (bool, error) {
	o := orm.NewOrm()
	user := models.User{UserName: UserName}
	err := o.QueryTable("user").Filter("UserName", UserName).Limit(1).One(&user)
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
