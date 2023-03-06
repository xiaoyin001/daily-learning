package xiaoyin

import (
    "WebServer/engine"
    "fmt"
    "net/http"
)

// HandlerFunc 定义路由映射的方法类型
type HandlerFunc func(*engine.Context)

type router struct {
    handlers map[string]HandlerFunc // 存放路由绑定方法
}

// 创建路由管理
func newRouter() *router {
    return &router{handlers: make(map[string]HandlerFunc)}
}

func (r router) addRouter(aMethod string, aPath string, aHandler HandlerFunc) {
    fmt.Println("添加路由 --> ", aMethod, "-", aPath)
    
    mKey := aMethod + "-" + aPath
    r.handlers[mKey] = aHandler
}

func (r router) ExecHandleFunc(aContext *engine.Context) {
    if aContext == nil {
        fmt.Println("ExecHandleFunc: aContext = nil")
        
        http.Error(aContext.ResponseWriter, "ExecHandleFunc: aContext = nil", http.StatusInternalServerError)
        return
    }
    
    fmt.Println("开始处理请求 --> ", aContext.Method, "-", aContext.Path)
    
    mKey := aContext.Method + "-" + aContext.Path
    mHandlerFunc, mOk := r.handlers[mKey]
    if mOk {
        mHandlerFunc(aContext)
    } else {
        aContext.RspString(http.StatusNotFound, "404 接口不存在: %s\n", aContext.Path)
    }
}
