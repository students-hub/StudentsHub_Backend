package models

import (
	"strings"
	"time"

	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 连接数据库
const (
	DB_USER_NAME = "root"
	DB_PASSWORD  = "10086KevinXu."
	DB_IP        = "localhost"
	DB_PORT      = "3306"
	DB_NAME      = "StudentsHub"
)

type User struct {
	UserID   int       `orm:"pk;auto"`   // UNIQUE, PRIMARY KEY, AUTO_INCREMENT
	Password string    `orm:"size(255)"` // NOT NULL
	UserName string    `orm:"size(255)"` // NOT NULL
	CreateAt time.Time `orm:"auto_now"`
	UpdateAt time.Time `orm:"auto_now"`
	Role     string    `orm:"size(1)"`
}

func init() {
	// 构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8mb4"
	dbInfo := strings.Join([]string{DB_USER_NAME, ":", DB_PASSWORD, "@tcp(", DB_IP, ":", DB_PORT, ")/", DB_NAME, "?charset=utf8mb4"}, "")
	orm.RegisterDataBase("default", "mysql", dbInfo)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
