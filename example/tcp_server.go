package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Start server...")

	listen, err := net.Listen("tcp", ":5000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			return
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		str := string(buf[:n])
		fmt.Printf("recv from client, data:%v\n", str)
	}
}