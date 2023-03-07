package test

import (
    "WebServer/engine"
    "WebServer/engine/xiaoyin"
    "fmt"
    "testing"
)

func Test3_2(t *testing.T) {
    mRoute := createTempRoute()
    
    mContext := engine.Context{
        ResponseWriter: nil,
        Request:        nil,
        Path:           "",
        Method:         "Get",
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

func tempHandlerFunc(aContext *engine.Context) {
    fmt.Printf("%s test URL: %s \n", "小印001", aContext.Path)
    if len(aContext.Params) > 0 {
    
    }
}

// 创建临时测试的路由
func createTempRoute() *xiaoyin.RouterMgr {
    mRoute := xiaoyin.NewRouter()
    
    mRoute.AddRouter("Get", "/", tempHandlerFunc)
    mRoute.AddRouter("Get", "/aa", tempHandlerFunc)
    mRoute.AddRouter("Get", "/aa/:name", tempHandlerFunc)
    mRoute.AddRouter("Get", "/aa/bb/cc", tempHandlerFunc)
    
    return mRoute
}
