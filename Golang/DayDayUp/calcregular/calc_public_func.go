package calcregular

// DoCalcUnit 外部调用单元计算
func DoCalcUnit(aSelfObj any, aTargetObj any, aUnitName string) (float64, bool) {
    mCU, mOK := uCalcMgr.getBindUnit(aUnitName)
    if mOK {
        return mCU.doCalc(aSelfObj, aTargetObj), mOK
    } else {
        return 0, false
    }
}
