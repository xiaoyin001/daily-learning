package test

import (
    "WebServer/engine/xiaoyin"
    "fmt"
    "strings"
    "testing"
)

func SplitPath(aPath string) []string {
    vs := strings.Split(aPath, "/")
    
    parts := make([]string, 0)
    for _, item := range vs {
        if item != "" {
            parts = append(parts, item)
            if item[0] == '*' {
                break
            }
        }
    }
    return parts
}

func Test3_1(t *testing.T) {
    mRoot := xiaoyin.TrieNode{
        Path:      "",
        Part:      "",
        ChildNode: nil,
        IsDim:     false,
    }
    
    mPath := "/"
    mParts := SplitPath(mPath)
    mRoot.Insert(mPath, mParts, 0)
    
    mPath = "/xiaoyin"
    mParts = SplitPath(mPath)
    mRoot.Insert(mPath, mParts, 0)
    
    mPath = "/xiaoyin02"
    mParts = SplitPath(mPath)
    mRoot.Insert(mPath, mParts, 0)
    
    mPath = "/xiaoyin/aa"
    mParts = SplitPath(mPath)
    mRoot.Insert(mPath, mParts, 0)
    
    mPath = "/xiaoyin/bb"
    mParts = SplitPath(mPath)
    mRoot.Insert(mPath, mParts, 0)
    
    mPath = "/xiaoyin/bb/:name"
    mParts = SplitPath(mPath)
    mRoot.Insert(mPath, mParts, 0)
    
    mPath = "/xiaoyin/bb/name/userid"
    mParts = SplitPath(mPath)
    mRoot.Insert(mPath, mParts, 0)
    
    mPath = "/xiaoyin/bb/name/age"
    mParts = SplitPath(mPath)
    mRoot.Insert(mPath, mParts, 0)
    
    mTempNode := mRoot.FindNode(mParts, 0)
    fmt.Println(mTempNode.String())
    
    mPath = "/xiaoyin/bb/name/test"
    mParts = SplitPath(mPath)
    mTempNode = mRoot.FindNode(mParts, 0)
    if mTempNode == nil {
        fmt.Println(mTempNode.String())
    }
    
}
