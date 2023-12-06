package main

import (
    "WeChatAPI/config"
    "WeChatAPI/global"
    "fmt"
    "math/rand"
    "time"
)

// TODO haha
// lala
func init() {
    fmt.Println("================ 开始初始化程序 ================")
    
    config.InitCfg()
    
    fmt.Println("================ 初始化程序完毕 ================")
}

// TODO 这是2
// FIXME 嘿嘿
func main() {
    
    rand.Seed(time.Now().UnixNano())
    
    fmt.Println("小印丶")
    
    global.UpdateAccessToken()
    time.Sleep(time.Second * 5)
    global.UpdateAccessToken()
    
}
