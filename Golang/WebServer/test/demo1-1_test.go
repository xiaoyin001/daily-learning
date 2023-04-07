package test

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func Test1_1(t *testing.T) {
	// 设置路由（请求路径与指定方法进行绑定）
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/xiaoyin", xiaoYinHandler)
	// 启动服务配置监听的端口号
	log.Fatal(http.ListenAndServe(":8888", nil))

	// log.Fatal() 函数在控制台屏幕上打印带有时间戳的指定消息。
	// log.Fatal() 类似于 log.Print() 函数，后跟调用 os.Exit(1) 函数
	// 意思就是打印后就退出了，懂得都懂
}

// 访问 URL=”/” 的处理方法
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("问 URL=”/” 开始处理")

	// 返回的参数1：返回的字节数
	mNum, err := fmt.Fprintf(w, "URL.Path = %q", req.URL.Path)
	if err != nil {
		return
	}

	fmt.Println("这个返回的是写入 w 中多少字节 mNum=", mNum)
}

// 访问 URL=”/xiaoyin” 的处理方法
func xiaoYinHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("访问 URL=”/xiaoyin” 开始处理")

	// 打印请求头的数据
	for k, v := range req.Header {
		_, err := fmt.Fprintf(w, "[%q] = %q\n", k, v)
		if err != nil {
			return
		}
	}
}

/*
用官方给的方式弄提个http服务那是”赶赶单单”，1行足以

http.ListenAndServe(":8888", nil)

至于上面为啥要写不止1行，开动小脑袋，我知道你会明白的
*/
