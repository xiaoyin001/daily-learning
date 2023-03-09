package xiaoyin

import (
    "net/http"
    "strings"
)

// RouterGroup 路由组
type RouterGroup struct {
    basePath    string        // 当前组的URL前缀
    middlewares []HandlerFunc // 中间件绑定方法
    engine      *Engine       // 所有地方都用是一个
}

// AddRouteGroup 添加路由组
func (group *RouterGroup) AddRouteGroup(aBasePath string) *RouterGroup {
    mNewGroup := &RouterGroup{
        basePath:    group.basePath + aBasePath,
        middlewares: make([]HandlerFunc, 0),
        engine:      group.engine,
    }
    
    group.engine.AllGroups = append(group.engine.AllGroups, mNewGroup)
    
    return mNewGroup
}

// AddMiddlewares 添加中间件绑定的方法
func (group *RouterGroup) AddMiddlewares(aMiddlewares ...HandlerFunc) {
    group.middlewares = append(group.middlewares, aMiddlewares...)
}

// GET 添加Get请求路由
func (group *RouterGroup) GET(aPath string, aHandler HandlerFunc) {
    mPath := group.basePath + aPath
    
    group.engine.RouterMgr.addRouter("GET", mPath, aHandler)
}

// POST 添加Post请求路由
func (group *RouterGroup) POST(aPath string, aHandler HandlerFunc) {
    mPath := group.basePath + aPath
    
    group.engine.RouterMgr.addRouter("POST", mPath, aHandler)
}

/**********************************************************************************************************************/

// Engine 实现HTTP服务接口的管理对象
type Engine struct {
    *RouterGroup                // 可以理解为继承 RouterGroup
    AllGroups    []*RouterGroup // 所有的路由组
    RouterMgr    *RouterMgr     // 路由管理
}

// Create 创建XiaoYinEngine服务管理对象
func Create() *Engine {
    mEngine := &Engine{
        RouterGroup: &RouterGroup{},
        AllGroups:   make([]*RouterGroup, 0),
        RouterMgr:   NewRouter(),
    }
    
    mEngine.middlewares = make([]HandlerFunc, 0)
    mEngine.engine = mEngine
    
    return mEngine
}

// GET 添加Get请求路由
func (e *Engine) GET(aPath string, aHandler HandlerFunc) {
    e.RouterMgr.addRouter("GET", aPath, aHandler)
}

// POST 添加Post请求路由
func (e *Engine) POST(aPath string, aHandler HandlerFunc) {
    e.RouterMgr.addRouter("POST", aPath, aHandler)
}

// StartServer 启动服务
func (e *Engine) StartServer(aAddr string) (err error) {
    return http.ListenAndServe(aAddr, e)
}

// 实现 ServeHTTP 方法，请求都会经过这里
func (e *Engine) ServeHTTP(aRspW http.ResponseWriter, aReq *http.Request) {
    mContext := NewContext(aRspW, aReq)
    
    // 先将匹配的中间件方法加入到切片中
    for _, mGroup := range e.AllGroups {
        if strings.HasPrefix(aReq.URL.Path, mGroup.basePath) {
            mContext.Handlers = append(mContext.Handlers, mGroup.middlewares...)
        }
    }
    
    // 然后再执行路由方法
    e.RouterMgr.ExecHandleFunc(mContext)
}
