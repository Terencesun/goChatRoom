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
			str := <- ch
			connMap := tcpWorker.GetConnMap()
			for k, v := range connMap{
				var sendStr string
				if k != str.ConnIndex {
					sendStr = fmt.Sprintf("【用户%v】%v\n", str.ConnIndex, string(str.Msg))
				} else {
					sendStr = fmt.Sprintf("【你】%v\n", string(str.Msg))
				}
				if _, err := v.Write([]byte(sendStr)); err != nil {
					fmt.Printf("用户%v，信息发送失败\n", k)
				}
			}
		}
	}()
	tcpWorker.Accept()
}
