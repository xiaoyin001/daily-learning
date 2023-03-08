package test

import (
    "WebServer/engine"
    "WebServer/engine/xiaoyin"
    "fmt"
    "testing"
)

func Test4_1(t *testing.T) {
    mEngine := xiaoyin.Create()
    
    mGroup1 := mEngine.AddRouteGroup("/R1")
    mGroup1.GET("/aa", tempHandlerFunc2)
    
    mGroup2 := mEngine.AddRouteGroup("/R2")
    mGroup2.GET("/aa", tempHandlerFunc2)
    
    mContext := engine.Context{
        ResponseWriter: nil,
        Request:        nil,
        Path:           "",
        Method:         "GET",
        Params:         nil,
        StatusCode:     0,
    }
    
    mContext.Path = "/R1/aa"
    mEngine.RouterMgr.ExecHandleFunc(&mContext)
    
    mContext.Path = "/R2/aa"
    mEngine.RouterMgr.ExecHandleFunc(&mContext)
}

func tempHandlerFunc2(aContext *engine.Context) {
    fmt.Printf("%s test URL: %s \n", "小印001", aContext.Path)
}
