package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"ginrequest/global"
)

func InitMySQL() {

	mysqlconfig := &global.GAppConfig.MySQL
	var err error
	//用户名：密码@tcp(地址:端口)/数据库名
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		mysqlconfig.User,
		mysqlconfig.Password,
		mysqlconfig.Host,
		mysqlconfig.Port,
		mysqlconfig.DBName,
	)

	fmt.Println("sql db:", url)

	db, err := sql.Open("mysql", url)
	global.GMySQL = db
	//database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	if err != nil {
		fmt.Println("---InitMySQL err:---", err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("---InitMySQL Ping err:---", err)
	}
}
