package main

import (
	"fmt"
	"server/tcpholder"
)

func main() {
	// 识别指令
	// 使用usermgr
	// 成功后，建立tcp
	tcpWorker := tcpholder.InitTcp(4000)
	ch := tcpWorker.GetChan()
	go func() {
		for {
			fmt.Println(<- ch)
		}
	}()
	tcpWorker.Accept()
}
