package test

import (
    "WebServer/engine/xiaoyin"
    "fmt"
    "testing"
)

func Test3_2(t *testing.T) {
    mRoute := createTempRoute()
    
    mContext := xiaoyin.Context{
        ResponseWriter: nil,
        Request:        nil,
        Path:           "",
        Method:         "GET",
        Params:         nil,
        StatusCode:     0,
    }
    
    mContext.Path = "/aa"
    mRoute.ExecHandleFunc(&mContext)
    
    mContext.Path = "/aa/xiaoyin"
    mRoute.ExecHandleFunc(&mContext)
    
    mContext.Path = "/aa/xiaoyin/cc"
    mRoute.ExecHandleFunc(&mContext)
    
    // mContext.Path = "/bb"
    // mRoute.ExecHandleFunc(&mContext)
}

func tempHandlerFunc(aContext *xiaoyin.Context) {
    fmt.Printf("%s test URL: %s \n", "小印001", aContext.Path)
    if len(aContext.Params) > 0 {
    
    }
}

// 创建临时测试的路由
func createTempRoute() *xiaoyin.RouterMgr {
    mRoute := xiaoyin.NewRouter()
    
    // mRoute.addRouter("GET", "/", tempHandlerFunc)
    // mRoute.addRouter("GET", "/aa", tempHandlerFunc)
    // mRoute.addRouter("GET", "/aa/:name", tempHandlerFunc)
    // mRoute.addRouter("GET", "/aa/bb/cc", tempHandlerFunc)
    
    return mRoute
}
