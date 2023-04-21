package calcregular

import (
    "fmt"
    "reflect"
    "strconv"
    "strings"
)

// 运算符
const (
    symbolTypeNone = 0
    symbolTypeAdd  = 1 // 加
    symbolTypeDel  = 2 // 减
    symbolTypeRide = 3 // 乘
    symbolTypeDiv  = 4 // 除
)

// 具体的值
type calcValue struct {
    value float64 // 具体计算值
}

// 根据接口获得
type calcGetValue struct {
    funcValue reflect.Value // 方法
    param     string        // 参数
    paramUnit []any         // 参数套娃【单元 or 实际参数】
}

// 运算符
type calcSymbol struct {
    symbolType int // 运算符
}

// 单元（同级运算组）
type calcUnit struct {
    calcUnitSlice []any // 存放计算单元组
}

// =================================================================

// 初始化
func (cu *calcUnit) init() {
    cu.calcUnitSlice = make([]any, 0)
}

// 表达式解析
func (cu *calcUnit) parse(aContent string) bool {
    // 去掉内容中的空格
    mContent := strings.Replace(aContent, " ", "", -1)
    
    // 开始检查括号
    mLeftNum, mRightNum := 0, 0
    for i := 0; i < len(mContent); i++ {
        if mContent[i] == '(' {
            mLeftNum++
        } else if mContent[i] == ')' {
            mRightNum++
        }
        
        if mLeftNum < mRightNum {
            panic("括号不对劲！！！")
            return false
        }
    }
    
    if mLeftNum != mRightNum {
        panic("括号数量不匹配呀！！！")
        return false
    }
    
    // 开始检查运算符
    mHasSymbol := false
    for i := 0; i < len(mContent); i++ {
        if mContent[i] == '+' || mContent[i] == '-' || mContent[i] == '*' || mContent[i] == '/' {
            if i == 0 {
                panic("第一个字符就是运算符，你不对劲呀！！！！")
                return false
            }
            
            if i == len(mContent)-1 {
                panic("最后一个字符是运算符，不对劲不对劲！！！")
                return false
            }
            
            if mHasSymbol {
                panic("连续2个运算符，不行不行！！！")
                return false
            } else {
                mHasSymbol = true
            }
        } else {
            mHasSymbol = false
        }
    }
    
    cu.doParse(mContent, len(mContent), 0, 1)
    return true
}

// 开始解析表达式
func (cu *calcUnit) doParse(aContent string, aLen int, aLeft int, aRight int) int {
    if aContent == "" || aLen <= 0 || aLeft < 0 || aLeft > aRight || aRight > aLen {
        return -1
    }
    
    mLeft := aLeft
    mRight := aRight
    for mRight <= aLen {
        if aContent[mRight-1:mRight] == "(" {
            mCU := calcUnit{}
            mCU.init()
            // 这里开始解析子单元数据
            mRight = mCU.doParse(aContent, aLen, mLeft+1, mRight+1)
            cu.calcUnitSlice = append(cu.calcUnitSlice, mCU)
            if mRight < 0 {
                return -1
            }
            
            // 能到这里说明本组后族还有需要进行解析的，跳过子单元已经解析的部分，继续解析后续内容
            mRight++
            mLeft = mRight
        } else if aContent[mRight-1:mRight] == "+" || aContent[mRight-1:mRight] == "-" || aContent[mRight-1:mRight] == "*" || aContent[mRight-1:mRight] == "/" {
            mCS := calcSymbol{}
            switch aContent[mRight-1 : mRight] {
            case "+":
                mCS.symbolType = symbolTypeAdd
            case "-":
                mCS.symbolType = symbolTypeDel
            case "*":
                mCS.symbolType = symbolTypeRide
            case "/":
                mCS.symbolType = symbolTypeDiv
            default:
            }
            
            if mLeft+1 < mRight {
                // 这里是将 运算符左边未解析部分的内容进行解析 例如：(10+5) 这里将10解析出来
                mValue := aContent[mLeft : mRight-1]
                mValueF64, mErr := strconv.ParseFloat(mValue, 64)
                if mErr != nil {
                    // 这里不是数字，是根据接口进行解析的
                    mCGV := calcGetValue{}
                    mCGV.parseParam(mValue)
                    
                    if mCGV.funcValue.IsValid() {
                        cu.calcUnitSlice = append(cu.calcUnitSlice, mCGV)
                    }
                } else {
                    // 这里是直接的数值
                    mCV := calcValue{}
                    mCV.value = mValueF64
                    
                    cu.calcUnitSlice = append(cu.calcUnitSlice, mCV)
                }
                
                // 内容解析完毕后将标记位置改变
                mLeft = mRight
            }
            
            // 然后在将运算符加入到运算组中
            cu.calcUnitSlice = append(cu.calcUnitSlice, mCS)
            mRight++
        } else if aContent[mRight-1:mRight] == ")" {
            if mLeft+1 < mRight {
                // 这里是将 运算符左边未解析部分的内容进行解析 例如：(10+5) 这里将5解析出来
                mValue := aContent[mLeft : mRight-1]
                mValueF64, mErr := strconv.ParseFloat(mValue, 64)
                if mErr != nil {
                    // 这里不是数字，是根据接口进行解析的
                    mCGV := calcGetValue{}
                    mCGV.parseParam(mValue)
                    
                    if mCGV.funcValue.IsValid() {
                        cu.calcUnitSlice = append(cu.calcUnitSlice, mCGV)
                    }
                } else {
                    // 这里是直接的数值
                    mCV := calcValue{}
                    mCV.value = mValueF64
                    
                    cu.calcUnitSlice = append(cu.calcUnitSlice, mCV)
                }
            }
            
            return mRight
        } else {
            mRight++
        }
    }
    
    if mRight > aLen {
        mRight = aLen
    }
    
    if mLeft+1 < mRight {
        // 这里是将最后未解析的部分进行解析 例如：(10+5)+6 这里将6解析出来
        mValue := aContent[mLeft:mRight]
        mValueF64, mErr := strconv.ParseFloat(mValue, 64)
        if mErr != nil {
            // 这里不是数字，是根据接口进行解析的
            mCGV := calcGetValue{}
            mCGV.parseParam(mValue)
            
            if mCGV.funcValue.IsValid() {
                cu.calcUnitSlice = append(cu.calcUnitSlice, mCGV)
            }
        } else {
            // 这里是直接的数值
            mCV := calcValue{}
            mCV.value = mValueF64
            
            cu.calcUnitSlice = append(cu.calcUnitSlice, mCV)
        }
    }
    
    return mRight
}

// 计算乘除
func calcRideOrDiv(aCV calcValue, aSymbol int, aSlice *[]any) {
    if aSlice == nil {
        return
    }
    
    if aSymbol != symbolTypeNone {
        if len(*aSlice) == 0 {
            *aSlice = append(*aSlice, aCV)
        } else {
            mCV, mOK := (*aSlice)[len(*aSlice)-1].(calcValue)
            if mOK {
                if aSymbol == symbolTypeRide {
                    mCV.value = mCV.value * aCV.value
                } else if aSymbol == symbolTypeDiv {
                    mCV.value = mCV.value / aCV.value
                }
                
                (*aSlice)[len(*aSlice)-1] = mCV
            }
        }
    } else {
        *aSlice = append(*aSlice, aCV)
    }
}

// 开始计算
func (cu *calcUnit) doCalc(aSelfObj any, aTargetObj any) float64 {
    mTempSlice := make([]any, 0)
    mCurrSymbol := symbolTypeNone
    for i := 0; i < len(cu.calcUnitSlice); i++ {
        switch cu.calcUnitSlice[i].(type) {
        case calcValue:
            mCV := cu.calcUnitSlice[i].(calcValue)
            calcRideOrDiv(mCV, mCurrSymbol, &mTempSlice)
        case calcGetValue:
            mCGV := cu.calcUnitSlice[i].(calcGetValue)
            mTempCV := calcValue{}
            mTempCV.value = mCGV.getValueCalc(aSelfObj, aTargetObj)
            calcRideOrDiv(mTempCV, mCurrSymbol, &mTempSlice)
        case calcSymbol:
            mCS := cu.calcUnitSlice[i].(calcSymbol)
            if mCS.symbolType == symbolTypeRide || mCS.symbolType == symbolTypeDiv {
                // 如果是 乘 or 除 修改当前运算符标记
                mCurrSymbol = mCS.symbolType
            } else {
                // 如果是 加 or 减 直接将内容加入到mTempSlice中
                mTempSlice = append(mTempSlice, mCS)
                mCurrSymbol = symbolTypeNone
            }
        case calcUnit:
            mCU := cu.calcUnitSlice[i].(calcUnit)
            mTempCV := calcValue{}
            mTempCV.value = mCU.doCalc(aSelfObj, aTargetObj)
            calcRideOrDiv(mTempCV, mCurrSymbol, &mTempSlice)
            mCurrSymbol = symbolTypeNone
        default:
        }
    }
    
    mCurrSymbol = symbolTypeNone
    var mNum float64
    
    // 计算剩下未进行与计算的加减
    for i := 0; i < len(mTempSlice); i++ {
        switch mTempSlice[i].(type) {
        case calcValue:
            mCV := mTempSlice[i].(calcValue)
            if mCurrSymbol == symbolTypeNone {
                mNum = mCV.value
            } else {
                if mCurrSymbol == symbolTypeAdd {
                    mNum = mNum + mCV.value
                } else if mCurrSymbol == symbolTypeDel {
                    mNum = mNum - mCV.value
                }
            }
        case calcSymbol:
            mCS := mTempSlice[i].(calcSymbol)
            mCurrSymbol = mCS.symbolType
        default:
        }
    }
    
    return mNum
}

// =================================================================

// 解析参数并填充
func (cgv *calcGetValue) parseParam(aParam string) {
    cgv.paramUnit = make([]any, 0)
    
    if aParam == "" {
        return
    }
    
    mSlice := strings.Split(aParam, "_")
    cgv.funcValue, _ = uCalcMgr.getBindFunc(mSlice[0])
    if len(mSlice) == 2 {
        mParamSlice := strings.Split(mSlice[1], "|")
        mIsAddParam := true
        for i := 0; i < len(mParamSlice); i++ {
            if ("正则" + mSlice[0]) == mParamSlice[i] {
                panic("套娃不能自己套自己")
            }
            
            mIdx := strings.Index(mParamSlice[i], "正则")
            if mIdx == 0 {
                // 套娃的参数单元
                mUnitName := mParamSlice[i][len("正则"):len(mParamSlice[i])]
                mUnitValue, mOk := uCalcMgr.getBindUnit(mUnitName)
                if !mOk {
                    panic("套娃参数不对劲呀")
                }
                cgv.paramUnit = append(cgv.paramUnit, mUnitValue)
                mIsAddParam = false
            } else {
                // 正常的参数
                cgv.paramUnit = append(cgv.paramUnit, mParamSlice[i])
            }
        }
        
        if mIsAddParam {
            cgv.param = mSlice[1]
        }
    } else {
        cgv.param = ""
    }
}

// 接口值计算
func (cgv *calcGetValue) getValueCalc(aSelfObj any, aTargetObj any) float64 {
    if aSelfObj == nil || aTargetObj == nil {
        // 这里的参数不要给nil，在掉计算方法的时候会抛异常
        return 0
    }
    
    mParam := make([]reflect.Value, 0)
    mParam = append(mParam, reflect.ValueOf(aSelfObj))
    mParam = append(mParam, reflect.ValueOf(aTargetObj))
    
    if cgv.param != "" {
        if cgv.funcValue.IsValid() {
            mParam = append(mParam, reflect.ValueOf(cgv.param))
        }
    } else {
        mParamStr := ""
        for i := 0; i < len(cgv.paramUnit); i++ {
            switch cgv.paramUnit[i].(type) {
            case calcUnit:
                mCU := cgv.paramUnit[i].(calcUnit)
                mNum := mCU.doCalc(aSelfObj, aTargetObj)
                if mParamStr == "" {
                    mParamStr = fmt.Sprintf("%f", mNum)
                } else {
                    mParamStr = mParamStr + "|" + fmt.Sprintf("%f", mNum)
                }
            case string:
                if mParamStr == "" {
                    mParamStr = cgv.paramUnit[i].(string)
                } else {
                    mParamStr = mParamStr + "|" + cgv.paramUnit[i].(string)
                }
            default:
            }
        }
        
        mParam = append(mParam, reflect.ValueOf(mParamStr))
    }
    
    mV := cgv.funcValue.Call(mParam)
    if len(mV) == 2 {
        if mV[1].Bool() {
            return mV[0].Float()
        }
    }
    
    return 0
}
