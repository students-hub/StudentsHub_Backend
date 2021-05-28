package controllers

import (
	"StudentsHub_Backend/models"

	"github.com/beego/beego/v2/adapter/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type CourseController struct {
	beego.Controller
}

// @Title Add Course
// @Description add course
// @Param	UserName		query 	string	true		"Course Creator"
// @Param	CourseName		query 	string	true		"Course Name"
// @Success 200 {string} add succeeded
// @Failure 403 add failed
// @router /add [put]
func (c *CourseController) AddCourse() {
	UserName := c.GetString("UserName")
	access := AccessQuery(UserName)
	if access[0] == '1' {
		CourseName := c.GetString("CourseName")
		var course models.Course
		course.CourseName = CourseName
		course.TeacherName = UserName
		o := orm.NewOrm()
		o.Begin()
		_, err := o.Insert(&course)
		if err != nil {
			logs.Info("添加失败，原因是:", err)
			c.Ctx.WriteString("添加失败，原因是:" + err.Error())
			o.Rollback() //回滚
			return
		} else {
			o.Commit() //加入
		}
		c.Data["json"] = map[string]int{"CourseID": course.CourseID} //返回CourseID, 给予用户提示
		c.ServeJSON()
		return
	} else {
		logs.Info("您没有创建班级的权限!")
		c.Data["json"] = "您没有创建班级的权限!"
		c.ServeJSON()
		return
	}
}

func AccessQuery(UserName string) string {
	var access string
	o := orm.NewOrm()
	err := o.Raw(`SELECT access FROM access NATURAL JOIN user WHERE user_name = ?`).SetArgs(UserName).QueryRow(&access)

	if err != nil {
		logs.Info("查询失败!")
	}
	return access
}
