package version_one

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)



//被调用的方法
func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put(w, r)
		return
	}
	if m == http.MethodGet {
		get(w, r)
		return
	}
}

//Put方法
func put(w http.ResponseWriter, r *http.Request) {
	//根据url路径创建文件
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	//检验错误
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//延时执行关闭操作
	defer f.Close()
	//把文件写入所写的路径
	io.Copy(f, r.Body)
}

//Get方法
func get(w http.ResponseWriter, r *http.Request) {
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}
