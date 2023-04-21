package calcregular

import "reflect"

var uCalcMgr calcMgr

func init() {
    uCalcMgr.init()
    uCalcMgr.loadCalcFunc()
    uCalcMgr.loadCalcUnit()
}

// 计算管理
type calcMgr struct {
    calcUnitBindMap map[string]calcUnit      // 当前支持的的计算单元
    calcFuncBindMap map[string]reflect.Value // 当前支持的接口方法
}

func (cm *calcMgr) init() {
    cm.calcUnitBindMap = make(map[string]calcUnit)
    cm.calcFuncBindMap = make(map[string]reflect.Value)
}

// 获取绑定的运算单元
func (cm *calcMgr) getBindUnit(aUnitName string) (calcUnit, bool) {
    mCU, mOK := cm.calcUnitBindMap[aUnitName]
    return mCU, mOK
}

// 获取绑定的接口方法
func (cm *calcMgr) getBindFunc(aFuncName string) (reflect.Value, bool) {
    mValue, mOk := cm.calcFuncBindMap[aFuncName]
    return mValue, mOk
}

// 根据方法名查找 calcInterface{} 中的方法
func (cm *calcMgr) getFuncValueByFuncNameStr(aFuncName string) reflect.Value {
    mValue := reflect.ValueOf(&calcInterface{})
    mFunc := mValue.MethodByName(aFuncName)
    return mFunc
}

// 加载接口方法
func (cm *calcMgr) loadCalcFunc() {
    // 这里的加载方式，根据个人需求来
    
    mValue := cm.getFuncValueByFuncNameStr("GetRandom")
    if mValue.IsValid() {
        cm.calcFuncBindMap["随机数"] = mValue
    }
}

// 加载表达式单元
func (cm *calcMgr) loadCalcUnit() {
    // 这里的加载方式，根据个人需求来
    
    mCU := calcUnit{}
    mCU.init()
    if mCU.parse("(1+2) *3+(4+5+6)+7-随机数_10|30+100") {
        cm.calcUnitBindMap["测试计算1"] = mCU
    }
    
    mCU = calcUnit{}
    mCU.init()
    if mCU.parse("200-随机数_正则测试计算1|200") {
        cm.calcUnitBindMap["测试计算2"] = mCU
    }
}
