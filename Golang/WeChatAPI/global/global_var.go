package global

import "time"

// 全局变量

var GAppBasePath string // 项目根目录

var GAppID string             // 开发者ID
var GAppSecret string         // 开发者密码
var GAccessToken string       // 访问Token
var GTokenNewTime time.Time   // Token创建时间
var GTokenExpiresTick float64 // Token到期Tick
