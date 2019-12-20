package db

import (
	"FILESTORE-SERVER/db/mysql"
	"fmt"
)

func SingIn()  {

}

//注册
func Register(username,paasword string) bool {

	//预处理
	stmt,err := mysql.DbConn().Prepare("insert ignore into tbl_user (`user_name`,`user_pwd`) value (?,?)")

	if err != nil {
		fmt.Println("Register fail err :"+err.Error())
		return false
	}

	defer stmt.Close()
	
	
	ret,err := stmt.Exec(username,paasword)

	if err != nil {
		fmt.Println("register Exec errr :"+err.Error())
		return false
	}

	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}

	return false
}