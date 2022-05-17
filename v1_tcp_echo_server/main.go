package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080") //监听端口
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client) //启动一个子进程
	}
	//接受一个请求，返回一个连接
}

//处理这个连接
func process(conn net.Conn) {
	defer conn.Close()              //退出函数，断开连接，根据生命周期
	reader := bufio.NewReader(conn) //创建一个只读，带缓冲流
	for {
		b, err := reader.ReadByte() //读一个字节
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b}) //写入字节，slice 类型转换
		if err != nil {
			break
		}
	}
}
