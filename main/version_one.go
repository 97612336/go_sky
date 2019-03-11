package main

import (
	"go_sky/version_one"
	"log"
	"net/http"
)

//设置主方法
func main() {
	http.HandleFunc("/objects/", version_one.Handler)
	log.Fatal(http.ListenAndServe(":12345", nil))
}
