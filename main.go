package main

import (
	"FILESTORE-SERVER/handler"
	"fmt"
	"net/http"
)

func main()  {
	//文件
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/suc",handler.UploadFinish)
	http.HandleFunc("/file/meta",handler.GetFileMetaHandler)
	http.HandleFunc("/file/download",handler.DownLoadHandler)
	http.HandleFunc("/file/delete",handler.FileDeleteHandler)
	http.HandleFunc("/file/update",handler.FileUpdateMetaHandler)

	//用户
	http.HandleFunc("/user/register",handler.Register)
	http.HandleFunc("/user/signin",handler.SignIn)
	err := http.ListenAndServe(":9001",nil)

	if err!= nil{
		fmt.Println("start server error:",err.Error())
	}

	fmt.Println("stat")
}