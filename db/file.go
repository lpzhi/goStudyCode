package db

import (
	"FILESTORE-SERVER/db/mysql"
	"database/sql"
	"fmt"
	"time"
)

func OnFileUploadFinished(filehash,filename string,filesize int64,fileaddr string) bool  {
	createAt := time.Now().Format("2006-01-02 15:04:05")

	prepareSql := "insert ignore into tbl_file (`file_sha1`,`file_name`,`file_size`,`file_addr`,`create_at`,`status`) value(?,?,?,?,?,1)"
	stmt,err := mysql.DbConn().Prepare(prepareSql)

	if err != nil {
		fmt.Println("Failed to prepare statement, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret,err := stmt.Exec(filehash,filename,filesize,fileaddr,createAt)
	fmt.Println(createAt)

	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Printf("File with hash:%s has been uploaded before", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName string
	FileSize int64
	FileAddr string
}

func GetFileMeta(filehash string) (*TableFile,error){
	stmt,err := mysql.DbConn().Prepare("select file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1=?")


	if err !=nil {
		fmt.Println("query file fial err:"+err.Error())
		return nil,err
	}

	defer stmt.Close()

	tableFile :=  TableFile{}
	if err := stmt.QueryRow(filehash).Scan(&tableFile.FileHash,&tableFile.FileName,&tableFile.FileSize,&tableFile.FileAddr);err!=nil{
		if err ==sql.ErrNoRows {
			return nil,nil
		}else{
			fmt.Println(err.Error())
			return nil,err
		}
	}
	return &tableFile,nil
}
