package handler

import (
	"MyGo/basic"
	"fmt"
	"net/http"
)

func Router() {
	http.HandleFunc("/sum", mySum)
	http.HandleFunc("/lissajous", basic.MyLisssajousHttp) // 页面会展示生成的gif动态图片
}

func mySum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sum Leen")
}
