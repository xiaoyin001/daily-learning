package main

import (
    "WeChatAPI/config"
    "WeChatAPI/global"
    "fmt"
    "time"
)

func init() {
    fmt.Println("================ 开始初始化程序 ================")
    
    config.InitCfg()
    
    fmt.Println("================ 初始化程序完毕 ================")
}

func main() {
    
    fmt.Println("小印丶")
    
    global.UpdateAccessToken()
    time.Sleep(time.Second * 5)
    global.UpdateAccessToken()
    
}
