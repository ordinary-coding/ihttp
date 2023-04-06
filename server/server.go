package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"runtime"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		readLen, err := reader.Read(buf[:])
		if err != nil {
			sysType := runtime.GOOS
			if sysType == "linux" {
				if err == io.EOF {
					fmt.Println("client close")
				} else {
					fmt.Println("read from client failed, err:", err)
				}
			} else {
				if strings.HasSuffix(err.Error(), "closed by the remote host.") {
					fmt.Println("client close")
				} else {
					fmt.Println("read from client failed, err:", err)
				}
			}
			break
		}
		fmt.Println("receive client msg:", string(buf[:readLen]))
	}
}

func RunServer() {

	listen, err := net.Listen("tcp", "0.0.0.0:9001")
	if err != nil {
		fmt.Println("listen failed, err:", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
