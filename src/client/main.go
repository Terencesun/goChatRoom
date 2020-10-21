package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	errorCode "server/errorcode"
	"strings"
)

func main() {
	if conn, err := net.Dial("tcp", "127.0.0.1:4000"); err != nil {
		fmt.Println("connect error", err)
		return
	} else {
		defer func() {
			if err := conn.Close(); err != nil {
				fmt.Println("connect error", err)
				return
			}
		}()
		go func() {
			inputReader := bufio.NewReader(os.Stdin)
			for {
				input, _ := inputReader.ReadString('\n')
				trimmedInput := strings.Trim(input, "\r\n")
				_, err := conn.Write([]byte(trimmedInput))
				if err != nil {
					return
				}
			}
		}()

		for {
			buf := make([]byte, 1024)
			if n, err := conn.Read(buf); err != nil {
				if find := strings.Contains(err.Error(), "forcibly closed"); find {
					// 远程退出
					fmt.Printf("服务端关闭\n")
					return
				} else {
					panic(errorCode.TCPREADERROR)
				}
			} else {
				fmt.Printf("%v", string(buf[0:n]))
			}
		}
	}
}
