package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	calc "gin111/the-way-to-go/14-unit-test"
	"io/ioutil"
	"log"
	"net/rpc"
)

type Result struct {
	Num, Ans int
}

// func main() {
// 	// 使用 rpc.DialHTTP 创建了 HTTP 客户端 client，并且创建了与 localhost:1234 的链接，1234 恰好是 RPC 服务监听的端口
// 	client, _ := rpc.DialHTTP("tcp", "localhost:1234")
//
// 	var result Result
// 	// 使用 rpc.Call 调用远程方法，第1个参数是方法名 Cal.Square，后两个参数与 Cal.Square 的定义的参数相对应。
// 	/*if err := client.Call("Cal.Square", 12, &result); err != nil {
// 		log.Fatal("Failed to call Cal.Square", err)
// 	}*/
//
// 	// 异步调用
// 	// client.Call 是同步调用的方式，会阻塞当前的程序，直到结果返回。如果有异步调用的需求，可以考虑使用client.Go，如下
// 	asyncCall := client.Go("Cal.Square", 12, &result, nil)
// 	// 2021/10/16 17:28:02 0^2 = 0
// 	log.Printf("%d^2 = %d", result.Num, result.Ans)
//
// 	<-asyncCall.Done
// 	// 2021/10/16 17:28:02 12^2 = 144
// 	log.Printf("%d^2 = %d", result.Num, result.Ans)
// }

// func main() {
// 	config := &tls.Config{InsecureSkipVerify: true}
//
// 	conn, _ := tls.Dial("tcp", "localhost:1234", config)
// 	defer conn.Close()
// 	client := rpc.NewClient(conn)
//
// 	var result Result
// 	// 使用 rpc.Call 调用远程方法，第1个参数是方法名 Cal.Square，后两个参数与 Cal.Square 的定义的参数相对应。
// 	if err := client.Call("Cal.Square", 12, &result); err != nil {
// 		log.Fatal("Failed to call Cal.Square", err)
// 	}
//
// 	log.Printf("%d^2 = %d", result.Num, result.Ans)
// }

func main() {
	sum := calc.Add(1, 2)
	fmt.Printf("calc.add sum= %#v \n", sum)
	certPool := x509.NewCertPool()
	// certBytes, err := ioutil.ReadFile("./server.crt")
	// certBytes, err := ioutil.ReadFile("gin111/studygo/rpc/02-client/server.crt")
	certBytes, err := ioutil.ReadFile("./studygo/rpc/02-client/server.crt")
	if err != nil {
		log.Fatal("Failed to read server.crt")
	}
	certPool.AppendCertsFromPEM(certBytes)

	config := &tls.Config{
		RootCAs: certPool,
	}

	conn, _ := tls.Dial("tcp", "localhost:1234", config)
	defer func(conn *tls.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	client := rpc.NewClient(conn)

	var result Result
	if err := client.Call("Cal.Square", 12, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}

	log.Printf("%d^2 = %d", result.Num, result.Ans)
}
