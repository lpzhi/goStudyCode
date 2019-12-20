package handler

import (
	"FILESTORE-SERVER/meta"
	"FILESTORE-SERVER/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter,r *http.Request)  {
	if r.Method=="GET" {
		data , err := ioutil.ReadFile("./static/view/index.html")

		if err !=nil {
			log.Fatal(err)
			io.WriteString(w,"inter error")
			return
		}

		io.WriteString(w,string(data))
	} else if r.Method=="POST"{
		//接收文件流及存储到本地目录
		file,head,err := r.FormFile("file")

		if err != nil {
			fmt.Printf("file upload fail ,err:%s\n",err.Error())
			return
		}

		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:"./upload/"+head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile,err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("file create fail ,err:%s\n",err.Error())
			return
		}

		defer newFile.Close()

		fileMeta.FileSize,err = io.Copy(newFile,file)

		if err != nil {
			fmt.Printf("file copy fail ,err:%s\n",err.Error())
			return
		}

		//计算文件hash值
		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)

		fmt.Println(fileMeta.FileSha1)
		//meta.UpdateFileMeta(fileMeta)
		//存储到mysql中
		_ = meta.UpdateFileMetaDb(fileMeta)
		//跳转到成功页面
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
		//io.WriteString(w,"upload success!!!")
	}


}

//UploadFinish:上传完成
func UploadFinish(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"upload finish!")
}

func GetFileMetaHandler(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	filehash := r.Form["filehash"][0]

	//fMeta := meta.GetFileMeta(filehash)
	fMeta,err := meta.GetFileMetaDb(filehash)
	if err !=nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err.Error())
		return
	}

	 data, err := json.Marshal(fMeta)

	 if err !=nil{
	 	w.WriteHeader(http.StatusInternalServerError)
		 return
	 }

	 w.Write(data)
}


func DownLoadHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()

	fsha1 := r.Form.Get("filehash")

	fm := meta.GetFileMeta(fsha1)

	f,err := os.Open(fm.Location)
	defer f.Close()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data,err := ioutil.ReadAll(f)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//设置头信息

	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("content-disposition","attachment;filename="+fm.FileName)
	w.Write(data)

}


//修改文件名
func FileUpdateMetaHandler(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	fileHash := r.Form.Get("fileHash")
	newFileName := r.Form.Get("fileName")

	//获取文件信息
	metaFile := meta.GetFileMeta(fileHash)
	metaFile.FileName = newFileName

	//更新文件信息
	meta.UpdateFileMeta(metaFile)

	w.WriteHeader(http.StatusOK)
}


//删除文件

func FileDeleteHandler(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()

	fileHash := r.Form.Get("fileHash")

	metaFile := meta.GetFileMeta(fileHash)

	os.Remove(metaFile.Location)

	meta.RemoveFileMetaHandler(fileHash)

	w.WriteHeader(http.StatusOK)
}


