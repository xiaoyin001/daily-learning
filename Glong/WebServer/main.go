package main

import (
    "WebServer/engine"
    "WebServer/engine/xiaoyin"
    "fmt"
    "net/http"
)

func main() {
    fmt.Println("哈哈哈，是时候开始了")
    
    mEngine := xiaoyin.Create()
    
    mEngine.GET("/", testFunc01)
    mEngine.POST("/xiaoyin", testFunc02)
    mEngine.GET("/json", testFunc02)
    mEngine.GET("/data", testFunc03)
    
    err := mEngine.StartServer(":8888")
    if err != nil {
        fmt.Println("服务启动失败 --> Err:", err.Error())
    }
}

// 这些是临时处理，后面会优化掉的哦

func testFunc01(aContext *engine.Context) {
    aContext.RspString(http.StatusOK, "这是以字符串的形式返回的 %s", "小印")
    
    // mNum, err := fmt.Fprintf(aContext.ResponseWriter, "当前请求路径=%q\n", aContext.Path)
    // if err != nil {
    // 	return
    // }
    //
    // fmt.Println("返回数据大小为", mNum, "字节")
}

type Temp struct {
    Name string
    Age  int
}

func testFunc02(aContext *engine.Context) {
    mTemp := Temp{
        Name: "小印6688",
        Age:  18,
    }
    aContext.RspJson(http.StatusOK, mTemp)
    
    // for k, v := range aContext.Request.Header {
    // 	_, err := fmt.Fprintf(aContext.ResponseWriter, "[%q] = %q\n", k, v)
    // 	if err != nil {
    // 		return
    // 	}
    // }
}

func testFunc03(aContext *engine.Context) {
    mData := []byte("这是Data的返回内容")
    aContext.RspData(http.StatusOK, mData)
    
    // for k, v := range aContext.Request.Header {
    // 	_, err := fmt.Fprintf(aContext.ResponseWriter, "[%q] = %q\n", k, v)
    // 	if err != nil {
    // 		return
    // 	}
    // }
}
