package xiaoyin

import (
    "WebServer/engine"
    "net/http"
)

// Engine 实现HTTP服务接口的管理对象
type Engine struct {
    routerMgr *router // 路由管理
}

// Create 创建XiaoYinEngine服务管理对象
func Create() *Engine {
    return &Engine{routerMgr: newRouter()}
}

// 添加路由
func (e *Engine) addRouter(aMethod string, aPath string, aHandler HandlerFunc) {
    e.routerMgr.addRouter(aMethod, aPath, aHandler)
}

// GET 添加Get请求路由
func (e *Engine) GET(aPath string, aHandler HandlerFunc) {
    e.addRouter("GET", aPath, aHandler)
}

// POST 添加Post请求路由
func (e *Engine) POST(aPath string, aHandler HandlerFunc) {
    e.addRouter("POST", aPath, aHandler)
}

// StartServer 启动服务
func (e *Engine) StartServer(aAddr string) (err error) {
    return http.ListenAndServe(aAddr, e)
}

// 实现 ServeHTTP 方法，请求都会经过这里
func (e *Engine) ServeHTTP(aRspW http.ResponseWriter, aReq *http.Request) {
    mContext := engine.NewContext(aRspW, aReq)
    e.routerMgr.ExecHandleFunc(mContext)
}
