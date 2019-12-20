package handler

import (
	"FILESTORE-SERVER/db"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SignIn(w http.ResponseWriter,r *http.Request)  {
	if r.Method == http.MethodGet{
		if data,err := ioutil.ReadFile("./static/view/signIn.html");err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}else {
			w.Write(data)
			return
		}
	}
}

func Register(w http.ResponseWriter,r *http.Request)  {



	if r.Method== http.MethodGet {
		//显示页面

		data , err:= ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}
		w.Write(data)
		return
	}else if (r.Method == http.MethodPost ){
		r.ParseForm()

		username := r.PostForm.Get("username")
		password := r.PostForm.Get("password")

		if ret := db.Register(username,password);ret{
			w.Write([]byte("SUCCESS"))
			return
		}
	}

	w.Write([]byte("fail"))




}
