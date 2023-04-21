package calcregular

import (
    "fmt"
    "math/rand"
    "strconv"
    "strings"
    "time"
)

type calcInterface struct{}

var uRandom *rand.Rand

func init() {
    uRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// 注意：方法参数和返回值 请按照规定设置！！！
// 注意：方法参数和返回值 请按照规定设置！！！
// 注意：方法参数和返回值 请按照规定设置！！！

// func (ci calcInterface) FuncName(aSelfObj any, aTargetObj any, aParam string) (float64, bool) {}

// 注意：方法参数和返回值 请按照规定设置！！！
// 注意：方法参数和返回值 请按照规定设置！！！
// 注意：方法参数和返回值 请按照规定设置！！！

// GetRandom 获取随机数【需要给随机区间，2个数之间用 | 分割】 例如：10|20  随机出来的整数包含 10和20
func (ci calcInterface) GetRandom(aSelfObj any, aTargetObj any, aParam string) (float64, bool) {
    mSlice := strings.Split(aParam, "|")
    if len(mSlice) != 2 {
        return 0, false
    }
    
    mMin, mErr := strconv.ParseFloat(mSlice[0], 64)
    if mErr != nil || mMin < 0 {
        return 0, false
    }
    
    mMax, mErr := strconv.ParseFloat(mSlice[1], 64)
    if mErr != nil || mMax < mMin {
        return 0, false
    }
    
    mRandomNum := uRandom.Intn(int(mMax-mMin)) + int(mMin)
    fmt.Println("Min=", mMin, "Max=", mMax, "本次随机数=", mRandomNum)
    return float64(mRandomNum), true
}
