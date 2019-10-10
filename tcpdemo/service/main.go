package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听....")
	//1 tcp 表示使用的网络协议是tcp
	//2 0.0.0.0:8888 表示本地监听8888端口
	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=",err)
		return 
	}
	//延时关闭listen
	defer listen.Close()
	//循环等待客户端来连接我	
	for {
		fmt.Println("等待客户端连接...")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err=",err)
		} else {
			fmt.Printf("Accept() success =%v\n",conn)
		}
		//这里起一个协程 为客户端服务
	}



	//fmt.Printf("listen suc=%v\n",listen)
}