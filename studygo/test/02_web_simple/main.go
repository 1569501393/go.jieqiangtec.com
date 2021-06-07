package main

import (
	"fmt"
	"log"
	"net/http"
	_ "strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数
	fmt.Println(r.Form)

	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("val", v)
	}

	fmt.Fprintf(w, "Hello This is Simple web.")
}

func main() {
	http.HandleFunc("/", sayHelloName) //设置路由

	err := http.ListenAndServe(":9090", nil) //设置监听端口
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
