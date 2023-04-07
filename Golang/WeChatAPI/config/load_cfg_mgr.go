package config

import (
    "WeChatAPI/global"
    "fmt"
    "os"
)

// 配置加载管理单元

// InitCfg 初始化配置入口
func InitCfg() {
    fmt.Println("================ 开始初始化配置 ================")
    
    global.GAppBasePath = getAppPath()
    
    loadSetupCfg()
    
    fmt.Println("================ 初始化配置完毕 ================")
}

// 获取项目根路径
func getAppPath() string {
    mPath, mErr := os.Getwd()
    if mErr != nil {
        fmt.Println("[Err]", mErr.Error())
        return ""
    }
    
    return mPath
}
