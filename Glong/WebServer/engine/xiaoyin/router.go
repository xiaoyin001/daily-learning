package xiaoyin

import (
    "WebServer/engine"
    "fmt"
    "net/http"
    "strings"
)

// HandlerFunc 定义路由映射的方法类型
type HandlerFunc func(*engine.Context)

type RouterMgr struct {
    roots    map[string]*TrieNode   // 路由字典树根节点
    handlers map[string]HandlerFunc // 存放路由绑定方法
}

// NewRouter 创建路由管理
func NewRouter() *RouterMgr {
    return &RouterMgr{
        roots:    make(map[string]*TrieNode),
        handlers: make(map[string]HandlerFunc),
    }
}

// 解析路径(只允许存在一个*)
func parsePath(aPath string) []string {
    // 将地址拆分为多个部分
    mPathSplit := strings.Split(aPath, "/")
    
    mParts := make([]string, 0)
    // 开始筛选拆分地址
    for _, mPart := range mPathSplit {
        if mPart != "" {
            mParts = append(mParts, mPart)
            
            if mPart[0] == '*' {
                break
            }
        }
    }
    return mParts
}

// AddRouter 添加路由
func (r RouterMgr) addRouter(aMethod string, aPath string, aHandler HandlerFunc) {
    fmt.Println("添加路由 --> ", aMethod, "-", aPath)
    
    // 获取某种请求类型的根节点
    _, mOk := r.roots[aMethod]
    if !mOk {
        r.roots[aMethod] = &TrieNode{}
    }
    
    // 拆分完整路径
    mParts := parsePath(aPath)
    // 将该地址（路由）插入到字典树中
    r.roots[aMethod].Insert(aPath, mParts, 0)
    
    // 为该路由绑定执行方法
    mKey := aMethod + "-" + aPath
    r.handlers[mKey] = aHandler
}

// 获取路由
func (r RouterMgr) getRoute(aMethod, aPath string) (*TrieNode, map[string]string) {
    // 获取请求类型的根节点
    mRoot, mOk := r.roots[aMethod]
    if !mOk {
        return nil, nil
    }
    
    // 拆分当前请求的路径
    mParts := parsePath(aPath)
    // 从字典树中找到对应路由的结点数据
    mNode := mRoot.FindNode(mParts, 0)
    
    if mNode != nil {
        mParam := make(map[string]string)
        // 拆分结点中添加的路由数据
        mRoteParts := parsePath(mNode.Path)
        // 遍历原始路径，有满足填充的参数就将其进行填充
        for i, mPart := range mRoteParts {
            if mPart[0] == ':' {
                mParam[mPart[1:]] = mParts[i]
            }
            
            if mPart[0] == '*' && len(mPart) > 1 {
                mParam[mPart[1:]] = strings.Join(mParts[i:], "/")
                break
            }
        }
        return mNode, mParam
    }
    
    return nil, nil
}

// ExecHandleFunc 执行路由方法
func (r RouterMgr) ExecHandleFunc(aContext *engine.Context) {
    if aContext == nil {
        fmt.Println("ExecHandleFunc: aContext = nil")
        
        http.Error(aContext.ResponseWriter, "ExecHandleFunc: aContext = nil", http.StatusInternalServerError)
        return
    }
    
    fmt.Println("开始处理请求 --> ", aContext.Method, "-", aContext.Path)
    
    mNode, mParams := r.getRoute(aContext.Method, aContext.Path)
    if mNode != nil {
        aContext.Params = mParams
        
        mKey := aContext.Method + "-" + mNode.Path
        // 在添加路由的时候就都是同时添加的，结点判断了这里就将判断省略了
        mHandlerFunc, _ := r.handlers[mKey]
        mHandlerFunc(aContext)
    } else {
        aContext.RspString(http.StatusNotFound, "404 接口不存在: %s\n", aContext.Path)
    }
}
