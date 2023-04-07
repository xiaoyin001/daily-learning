package global

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "time"
)

// 全局方法

// UpdateAccessToken 更新Access token【微信公众号】
func UpdateAccessToken() {
    // Token未失效前禁止更新
    mNum := time.Since(GTokenNewTime)
    if mNum.Seconds() < GTokenExpiresTick {
        return
    }
    
    mURL := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
    mURL = mURL + "&appid=" + GAppID
    mURL = mURL + "&secret=" + GAppSecret
    
    mResp, mErr := http.Get(mURL)
    if mErr != nil {
        fmt.Println("[Err]", mErr.Error())
        return
    }
    
    // 一定要关闭返回的response中的body
    defer func(Body io.ReadCloser) {
        err := Body.Close()
        if err != nil {
            fmt.Println("[Err]", err.Error())
        }
    }(mResp.Body)
    
    // 得到返回结果
    mBody, mErr := io.ReadAll(mResp.Body)
    if mErr != nil {
        fmt.Println("[Err]", mErr.Error())
        return
    }
    
    // 对返回的JSON做解析
    mDataMap := make(map[string]interface{}, 0)
    mErr = json.Unmarshal(mBody, &mDataMap)
    if mErr != nil {
        fmt.Println("[Err]", mErr.Error())
        return
    }
    
    for mKey, mValue := range mDataMap {
        if mKey == "access_token" {
            GAccessToken = mValue.(string)
            GTokenNewTime = time.Now()
        } else if mKey == "expires_in" {
            GTokenExpiresTick = mValue.(float64)
        } else if mKey == "errcode" {
            fmt.Println(mKey, mValue)
        } else if mKey == "errmsg" {
            fmt.Println(mKey, mValue)
        }
    }
}
