package mysql

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB


func init()  {
	db,_=sql.Open("mysql","root:123456,@tcp(127.0.0.1:3306)/test?charset=utf8")

	db.SetMaxOpenConns(1000)
	if err := db.Ping();err != nil{
		fmt.Println("fial to conn db err:"+err.Error())
		os.Exit(1)
	}
}

func DbConn() * sql.DB {
	return db
}