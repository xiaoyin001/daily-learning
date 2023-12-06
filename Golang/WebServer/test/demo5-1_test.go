package test

import (
    "WebServer/engine/xiaoyin"
    "fmt"
    "math/rand"
    "testing"
    "time"
)

func Test5_1(t *testing.T) {
    
    rand.Seed(time.Now().UnixNano())
    
    mEngine := xiaoyin.Create()
    
    mGroup0 := mEngine.AddRouteGroup("/")
    mGroup0.AddMiddlewares(tempHandlerFunc4)
    
    mGroup1 := mEngine.AddRouteGroup("/R1")
    mGroup1.AddMiddlewares(tempHandlerFunc4)
    mGroup1.GET("/aa", tempHandlerFunc3)
    mGroup1.GET("/aa/bb", tempHandlerFunc3)
    
    mGroup2 := mEngine.AddRouteGroup("/R2")
    mGroup2.GET("/aa", tempHandlerFunc3)
    
    mContext := xiaoyin.Context{
        ResponseWriter: nil,
        Request:        nil,
        Path:           "",
        Method:         "GET",
        Params:         nil,
        StatusCode:     0,
        Handlers:       make([]xiaoyin.HandlerFunc, 0),
        Idx:            -1,
    }
    
    mContext.Path = "/R1/aa"
    mEngine.RouterMgr.ExecHandleFunc(&mContext)
    mContext.Idx = -1
    mContext.Path = "/R1/aa/bb"
    mEngine.RouterMgr.ExecHandleFunc(&mContext)
    
    mContext.Idx = -1
    mContext.Path = "/R2/aa"
    mEngine.RouterMgr.ExecHandleFunc(&mContext)
}

func tempHandlerFunc3(aContext *xiaoyin.Context) {
    fmt.Printf("%s test URL: %s \n", "小印001", aContext.Path)
}

func tempHandlerFunc4(aContext *xiaoyin.Context) {
    fmt.Printf("%s test中间件的执行方法 URL: %s \n", "小印001", aContext.Path)
}
