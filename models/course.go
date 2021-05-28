package models

import (
	"time"

	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Course struct {
	CourseID    int       `orm:"auto;pk"`   // PRIMARY KEY, AUTO_INCREMENT
	CourseName  string    `orm:"size(255)"` // PRIMARY KEY
	TeacherName string    `orm:"size(255)"` // PRIMARY KEY
	CreateAt    time.Time `orm:"auto_now"`
	UpdateAt    time.Time `orm:"auto_now"`
}

type User_Course struct {
	UserName string `orm:"size(255);pk"`
	CourseID int    `orm:"size(255)"`
}

type Access struct {
	Role   string `orm:"size(1);pk"` // PRIMARY KEY
	Access string `orm:"size(255)"`  // PRIMARY KEY
}

func init() {
	orm.RegisterModel(new(Course), new(User_Course), new(Access))
}
