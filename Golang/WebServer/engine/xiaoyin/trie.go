package xiaoyin

import (
    "fmt"
    "strings"
)

// TrieNode 字典树结点
type TrieNode struct {
    Path      string      // 路由(请求路径)，例如：/xiaoyin/aa/:name
    Part      string      // 路由的一部分，例如：xiaoyin、aa、:name
    ChildNode []*TrieNode // 子节点，例如：当前结点Part=aa，那么其子节点有 :name
    IsDim     bool        // 是否模糊匹配，Part 含有 : 或 * 时为true
}

// 结构体转String
func (tn *TrieNode) String() string {
    return fmt.Sprintf("TrieNode{Path=%s, Part=%s, IsDim=%t}", tn.Path, tn.Part, tn.IsDim)
}

// Insert 插入节点
func (tn *TrieNode) Insert(aPath string, aParts []string, aIdx int) {
    if len(aParts) == aIdx {
        // TODO 这种情况会出现后面添加的路由覆盖前面的路由，关系是否允许后期覆盖或者需要怎么处理，这里可以添加提示或者直接让服务起不来，具体根据实际情况而定
        // 当前路由最后一段，那么改结点对应的是当前路由
        tn.Path = aPath
        return
    }
    
    // 获取当前匹配的路径段
    mPart := aParts[aIdx]
    // 获取匹配到的第一个子路由
    mChild := tn.MatchFirstChild(mPart)
    // 如果不存在这个子节点
    if mChild == nil {
        // 创建这个结点
        mChild = &TrieNode{Part: mPart, IsDim: mPart[0] == ':' || mPart[0] == '*'}
        // 然后将擦行间的结点加入到当前结点的子节点中
        tn.ChildNode = append(tn.ChildNode, mChild)
    }
    // 然后开始地柜，弄下一个，直到最后一个结点，然后给其设置一下路径示例，然后就没了
    mChild.Insert(aPath, aParts, aIdx+1)
}

// FindNode 查找结点
func (tn *TrieNode) FindNode(aParts []string, aIdx int) *TrieNode {
    // 数量层数 == 检查深度  或  当前结点是以”*”开头的
    if len(aParts) == aIdx || strings.HasPrefix(tn.Part, "*") {
        if tn.Path == "" {
            return nil
        }
        
        // 返回内容不为空表示找到了对应的结点
        return tn
    }
    
    // 获取当前查询的结点
    mPart := aParts[aIdx]
    // 获取所有满足条件的子节点
    mAllChildNode := tn.MatchAllChildren(mPart)
    
    for _, child := range mAllChildNode {
        // 遍历所有的子节点开始匹配
        mResult := child.FindNode(aParts, aIdx+1)
        // 返回内容不为空表示找到了
        if mResult != nil {
            return mResult
        }
    }
    
    return nil
}

// MatchFirstChild 匹配第一个子路由
func (tn *TrieNode) MatchFirstChild(aPart string) *TrieNode {
    // 遍历所有的子节点
    for _, mChild := range tn.ChildNode {
        // 如果这个子路由的地址相同  或  这个子节点是模糊匹配的
        if mChild.Part == aPart || mChild.IsDim {
            return mChild
        }
    }
    return nil
}

// MatchAllChildren 匹配所有子路由
func (tn *TrieNode) MatchAllChildren(aPart string) []*TrieNode {
    // 先创建一个切片存用来存所有的子节点
    mNodes := make([]*TrieNode, 0)
    // 遍历所有的子节点
    for _, mChild := range tn.ChildNode {
        // 只要条件满足的都加入到切片中
        if mChild.Part == aPart || mChild.IsDim {
            mNodes = append(mNodes, mChild)
        }
    }
    
    return mNodes
}
