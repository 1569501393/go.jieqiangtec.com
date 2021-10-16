package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

type Cal int

/*func (cal *Cal) Square(num int) *Result {
	return &Result{
		Num: num,
		Ans: num * num,
	}
}*/

// Square net/rpc 规范
/**
RPC 需要满足什么条件
虽然说，远程过程调用并不需要我们关心如何编解码，如何通信，但是最基本的，如果一个方法需要支持远程过程调用，需要满足一定的约束和规范。不同 RPC 框架的约束和规范是不同的，如果使用 Golang 的标准库 net/rpc，方法需要长这个样子：

func (t *T) MethodName(argType T1, replyType *T2) error
即需要满足以下 5 个条件：

方法类型（T）是导出的（首字母大写）
方法名（MethodName）是导出的
方法有2个参数(argType T1, replyType *T2)，均为导出/内置类型
方法的第2个参数一个指针(replyType *T2)
方法的返回值类型是 error
*/
func (cal *Cal) Square(num int, result *Result) error {
	result.Num = num
	result.Ans = num * num
	return nil

}

/*func main() {
	cal := new(Cal)
	result := cal.Square(12)
	// 2021/10/16 16:58:08 12^2 = 144
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	// 12^2 = 144
	fmt.Printf("%d^2 = %d", result.Num, result.Ans)
}*/

/*func main() {
	cal := new(Cal)

	var result Result
	cal.Square(12, &result)
	// 2021/10/16 16:58:08 12^2 = 144
	log.Printf("%d^2 = %d", result.Num, result.Ans)

	// 12^2 = 144
	fmt.Printf("%d^2 = %d", result.Num, result.Ans)
}*/

/*func main() {
	// 使用 rpc.Register，发布 Cal 中满足 RPC 注册条件的方法（Cal.Square）
	rpc.Register(new(Cal))
	// 使用 rpc.HandleHTTP 注册用于处理 RPC 消息的 HTTP Handler
	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 1234)
	// 使用 http.ListenAndServe 监听 1234 端口，等待 RPC 请求。
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error servering:", err)
	}
}*/

// 客户端对服务器端鉴权
// func main() {
// 	// 使用 rpc.Register，发布 Cal 中满足 RPC 注册条件的方法（Cal.Square）
// 	rpc.Register(new(Cal))
//
// 	// HTTP 协议默认是不加密的，我们可以使用证书来保证通信过程的安全。
// 	// 生成私钥和自签名的证书，并将 server.key 权限设置为只读，保证私钥的安全。
// 	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
// 	config := &tls.Config{
// 		Certificates: []tls.Certificate{cert},
// 	}
// 	listener, _ := tls.Listen("tcp", ":1234", config)
// 	log.Printf("Serving RPC server on port %d", 1234)
//
// 	for {
// 		conn, _ := listener.Accept()
// 		defer conn.Close()
// 		go rpc.ServeConn(conn)
// 	}
// }

func main() {
	err := rpc.Register(new(Cal))
	if err != nil {
		return
	}
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)

	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
