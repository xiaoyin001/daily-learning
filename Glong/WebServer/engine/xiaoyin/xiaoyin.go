package xiaoyin

import (
	"fmt"
	"net/http"
)

// HandlerFunc 定义路由映射的方法类型
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 实现HTTP服务接口的管理对象
type Engine struct {
	router map[string]HandlerFunc // 存放路由方法
}

// Create 创建XiaoYinEngine服务管理对象
func Create() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由
func (engine *Engine) addRoute(aMethod string, aPath string, aHandler HandlerFunc) {
	key := aMethod + "-" + aPath
	engine.router[key] = aHandler
}

// GET 添加Get请求路由
func (engine *Engine) GET(aPath string, aHandler HandlerFunc) {
	engine.addRoute("GET", aPath, aHandler)
}

// POST 添加Post请求路由
func (engine *Engine) POST(aPath string, aHandler HandlerFunc) {
	engine.addRoute("POST", aPath, aHandler)
}

// StartServer 启动服务器
func (engine *Engine) StartServer(aAddr string) (err error) {
	return http.ListenAndServe(aAddr, engine)
}

// 实现 ServeHTTP 方法，请求都会经过这里
func (engine *Engine) ServeHTTP(aRspW http.ResponseWriter, aReq *http.Request) {
	// 获取请求路径
	key := aReq.Method + "-" + aReq.URL.Path

	// 从路由的 map 中找到指定的方法进行处理
	mHandlerFunc, mOk := engine.router[key]
	if mOk {
		mHandlerFunc(aRspW, aReq)
	} else {
		_, err := fmt.Fprintf(aRspW, "404 无效接口: %s\n", aReq.URL)
		if err != nil {
			return
		}
	}
}
