package main

import (
	"WebServer/engine/xiaoyin"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("哈哈哈，是时候开始了")

	mEngine := xiaoyin.Create()

	mGroup0 := mEngine.AddRouteGroup("/")
	mGroup0.AddMiddlewares(testFunc02)
	mGroup0.GET("", testFunc01)

	mGroup1 := mEngine.AddRouteGroup("/xiaoyin01")
	mGroup1.AddMiddlewares(testFunc02)
	mGroup1.GET("/aa", testFunc01)
	mGroup1.GET("/aa/bb", testFunc01)

	mGroup2 := mEngine.AddRouteGroup("/xiaoyin02")
	mGroup2.AddMiddlewares(testFunc02)
	mGroup2.GET("/aa", testFunc01)
	mGroup2.GET("/aa/bb", testFunc01)

	err := mEngine.StartServer(":8888")
	if err != nil {
		fmt.Println("服务启动失败 --> Err:", err.Error())
	}
}

func testFunc01(aContext *xiaoyin.Context) {
	mContent := fmt.Sprintf(" 测试返回内容 URL= %s \n", aContext.Path)
	aContext.RspString(http.StatusOK, "%s"+mContent, "小印01")

	fmt.Printf("%s test路由的执行方法 URL: %s  Idx=%d \n", "小印01", aContext.Path, aContext.Idx)
}

func testFunc02(aContext *xiaoyin.Context) {
	fmt.Printf("%s test中间件的执行方法 URL: %s  Idx=%d \n", "小印02", aContext.Path, aContext.Idx)
}
