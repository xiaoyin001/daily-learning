package engine

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Context 上下文对象
type Context struct {
    ResponseWriter http.ResponseWriter // 响应写入对象
    Request        *http.Request       // 请求对象
    Path           string              // 请求路径
    Method         string              // 请求方式
    Params         map[string]string   // 模糊匹配参数
    StatusCode     int                 // 返回状态码
}

// NewContext 创建上下文
func NewContext(aRspW http.ResponseWriter, aReq *http.Request) *Context {
    return &Context{
        ResponseWriter: aRspW,
        Request:        aReq,
        Path:           aReq.URL.Path,
        Method:         aReq.Method,
        StatusCode:     0,
    }
}

// PostForm 获取Post表单数据ByKey
func (c *Context) PostForm(aKey string) string {
    return c.Request.PostFormValue(aKey)
}

// Query 获取URL中键值对的值
func (c Context) Query(aKey string) string {
    return c.Request.URL.Query().Get(aKey)
}

// 设置状态码
func (c Context) setStatus(aCode int) {
    c.StatusCode = aCode
    c.ResponseWriter.WriteHeader(aCode)
}

// 添加响应返回头
func (c *Context) addRspHeader(aKey string, aValue string) {
    c.ResponseWriter.Header().Set(aKey, aValue)
}

// RspString 以文本形式返回
func (c Context) RspString(aCode int, aFormat string, aValue ...interface{}) {
    c.addRspHeader("Content-Type", "text/plain")
    c.setStatus(aCode)
    mContent := []byte(fmt.Sprintf(aFormat, aValue))
    _, err := c.ResponseWriter.Write(mContent)
    if err != nil {
        fmt.Println("Err:", err.Error())
        return
    }
}

// RspJson 以JSON格式返回
func (c Context) RspJson(aCode int, aJsonObj interface{}) {
    c.addRspHeader("Content-Type", "application/json")
    c.setStatus(aCode)
    mEncoder := json.NewEncoder(c.ResponseWriter)
    err := mEncoder.Encode(aJsonObj)
    if err != nil {
        http.Error(c.ResponseWriter, err.Error(), http.StatusInternalServerError)
    }
}

// RspData 原始数据返回
func (c *Context) RspData(aCode int, aData []byte) {
    c.setStatus(aCode)
    _, err := c.ResponseWriter.Write(aData)
    if err != nil {
        fmt.Println("Err:", err.Error())
        return
    }
}

// RspHTML 以HTML形式返回
func (c *Context) RspHTML(aCode int, aHTMLStr string) {
    c.addRspHeader("Content-Type", "text/html")
    c.setStatus(aCode)
    _, err := c.ResponseWriter.Write([]byte(aHTMLStr))
    if err != nil {
        fmt.Println("Err:", err.Error())
        return
    }
}
