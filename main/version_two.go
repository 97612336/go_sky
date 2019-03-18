package main

import (
	"go_sky/heartbeat"
	"go_sky/locate"
	"go_sky/version_two"
	"log"
	"net/http"
	"os"
)

//设置主方法
func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", version_two.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
