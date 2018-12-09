package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func sayHelloName(response http.ResponseWriter, request *http.Request) {
	// 解析参数，默认是不会解析的
	request.ParseForm()
	// 从客户端获取到的信息
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_log"])

	for k, v := range request.Form {
		fmt.Println("k: ", k)
		fmt.Println("v: ", strings.Join(v, ""))
	}
	//从服务器端写出到客户端
	fmt.Fprintf(response, "Hello Xdhuxc")
}

func main() {
	// 设置访问的路由
	http.HandleFunc("/", sayHelloName)
	// 设置监听的端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}