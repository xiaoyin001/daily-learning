package main

import (
    "DayDayUp/calcregular"
    "fmt"
    "runtime"
)

func main() {
    fmt.Println(runtime.Version())
    
    mNum, _ := calcregular.DoCalcUnit(1, 2, "测试计算1")
    fmt.Println("计算结果=", mNum)
    
    mNum, _ = calcregular.DoCalcUnit(1, 2, "测试计算2")
    fmt.Println("计算结果=", mNum)
}
