package main

import (
	"WebServer/engine/xiaoyin"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("哈哈哈，是时候开始了")

	mEngine := xiaoyin.Create()

	mEngine.GET("/", testFunc01)
	mEngine.POST("/xiaoyin", testFunc02)

	err := mEngine.StartServer(":8888")
	if err != nil {
		fmt.Println("服务启动失败 --> Err:", err.Error())
	}
}

// 这些是临时处理，后面会优化掉的哦

func testFunc01(aRspW http.ResponseWriter, aReq *http.Request) {
	mNum, err := fmt.Fprintf(aRspW, "当前请求路径=%q\n", aReq.URL.Path)
	if err != nil {
		return
	}

	fmt.Println("返回数据大小为", mNum, "字节")
}

func testFunc02(aRspW http.ResponseWriter, aReq *http.Request) {
	for k, v := range aReq.Header {
		_, err := fmt.Fprintf(aRspW, "[%q] = %q\n", k, v)
		if err != nil {
			return
		}
	}
}
