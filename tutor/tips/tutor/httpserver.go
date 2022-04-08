package tutor

import (
	"fmt"
	"net/http"
)

// *handlers* 是 `net/http` 服务器里面的一个基本概念。
// handler 对象实现了 `http.Handler` 接口。
// 编写 handler 的常见方法是，在具有适当签名的函数上使用 `http.HandlerFunc` 适配。
func hello(w http.ResponseWriter, req *http.Request) {

	// handler 函数有两个参数，`http.ResponseWriter` 和 `http.Request`。
	// response writer 被用于写入 HTTP 响应数据，这里我们简单的返回 "hello\n"。
	_, err := fmt.Fprintf(w, "hello\n")
	if err != nil {
		return
	}
}

func headers(w http.ResponseWriter, req *http.Request) {

	// 这个 handler 稍微复杂一点，
	// 我们需要读取的 HTTP 请求 header 中的所有内容，并将他们输出至 response body。
	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				return
			}
		}
	}
}

func main() {

	// 使用 `http.HandleFunc` 函数，可以方便的将我们的 handler 注册到服务器路由。
	// 它是 `net/http` 包中的默认路由，接受一个函数作为参数。
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 最后，我们用端口和处理程序调用 `ListenAndServe`。`nil` 告诉它使用我们刚刚设置的默认路由器。
	http.ListenAndServe(":8090", nil)
}
