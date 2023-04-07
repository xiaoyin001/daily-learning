package config

import (
    "WeChatAPI/global"
    "gopkg.in/ini.v1"
)

// 加载 Setup.ini 配置
func loadSetupCfg() {
    mFileName := "D:/BBB_MyCode/DayDayUp/Golang/WeChatAPI/bin/config/Setup.ini"
    mCfg, err := ini.Load(mFileName)
    if err != nil {
        return
    }
    
    global.GAppID = mCfg.Section("base").Key("app_id").String()
    global.GAppSecret = mCfg.Section("base").Key("app_secret").String()
}
