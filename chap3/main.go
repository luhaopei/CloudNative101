package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/version", GetVersion)
	http.HandleFunc("/healthz", Healthz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request) {
	// 接收客户端 request，并将 request 中带的 header 写入 response header
	println("IP:", r.Host, " HTTP Status Code:", 200)
	
	for key, values := range r.Header {
		for _, value := range values {
			w.Header().Set(key, value)
		}
	}
	w.WriteHeader(200)
}
func Healthz(w http.ResponseWriter, r *http.Request) {
	// 当访问 localhost/healthz 时，应返回200
	println("IP:", r.Host, " HTTP Status Code:", 200)
	w.WriteHeader(200)
}
func GetVersion(w http.ResponseWriter, r *http.Request) {
	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	println("IP:", r.Host, " HTTP Status Code:", 200)
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	w.WriteHeader(200)
}
