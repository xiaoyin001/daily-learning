package test

import (
	"fmt"
	"net/http"
	"testing"
)

/*
我们直接将1-1中最重要的一行弄过来看看
http.ListenAndServe(":9999", nil)

是不是要思考一下，为啥第二个参数给的是 nil，如果我们给一个不是 nil 的参数会出现什么效果
我点进去看看源码，第二个参数我们需要给什么
type Handler interface {
	// 有没有 ResponseWriter 和 Request 红红的，不知道在那个包里面的，动动你们的小手，找找看，嘿嘿嘿
	ServeHTTP(ResponseWriter, *Request)
}

发现是一个只有一个方法的接口，到这里额悟了，这不就是让俺弄个结构体，然后实现这个方法不就可以了嘛
那就试试看？
*/

// 小印引擎？小印发动机？哈哈哈哈哈
type TXiaoYinEngine struct {
}

// 这就是我们自己实现的 ServeHTTP 方法
func (x TXiaoYinEngine) ServeHTTP(aRspW http.ResponseWriter, aReq *http.Request) {
	fmt.Println("访问 URL=", aReq.URL.Path, "开始处理")

	switch aReq.URL.Path {
	case "/":
		mNum, err := fmt.Fprintf(aRspW, "当前请求路径=%q\n", aReq.URL.Path)
		if err != nil {
			return
		}

		fmt.Println("返回数据大小为", mNum, "字节")

	case "/xiaoyin":
		for k, v := range aReq.Header {
			_, err := fmt.Fprintf(aRspW, "[%q] = %q\n", k, v)
			if err != nil {
				return
			}
		}

	default:
		_, err := fmt.Fprintf(aRspW, "404 无效路径: %s\n", aReq.URL)
		if err != nil {
			return
		}
	}
}

func Test1_2(t *testing.T) {
	mXiaoYinEngine := new(TXiaoYinEngine)

	err := http.ListenAndServe(":8888", mXiaoYinEngine)
	if err != nil {
		fmt.Println("ListenAndServe --> Err:", err.Error())
		return
	}
}

/*
经过上面的测试，感觉如何，是不是发现了好东西

第二个参数给了，就相当于单身时候赚的钱都冲了游戏，突然有一天脱单了，经济大权由女朋友控制了
后面再想充游戏就得经过女朋友同意才可以进行下去，嗯...差不多就是这个意思，懂的都懂

既然发现了，继续改进改进
*/
