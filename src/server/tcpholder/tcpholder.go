package tcpholder

import (
	"fmt"
	"net"
	errorCode "server/errorcode"
	"strconv"
)

type Msg struct {
	connIndex int
	msg string
}

type TcpWorker struct {
	listener net.Listener
	connIndex int
	connMap map[int]net.Conn
	publicChan chan Msg
	handle func()
}

func (p *TcpWorker) startTcpListener(port int) {
	fmt.Println("start tcp server")
	p.connIndex = 0
	p.publicChan = make(chan Msg, 1024)
	p.connMap = make(map[int]net.Conn)
	if conn, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port)); err != nil {
		panic(err)
	} else {
		p.listener = conn
	}
}

func (p *TcpWorker) Accept()  {
	for {
		if conn, err := p.listener.Accept(); err != nil {
			fmt.Println(errorCode.TCPACCEPTERROR)
			continue
		} else {
			p.connIndex ++
			p.connMap[p.connIndex] = conn
			// todo 记录每位用户的conn
			fmt.Printf("用户%v进入\n", p.connIndex)
			go p.process(p.connIndex, conn)
		}
	}
}

func (p *TcpWorker) process(index int, conn net.Conn)  {
	defer func() {
		if err := conn.Close(); err != nil {
			panic(errorCode.TCPCLOSEERROR)
		}
	}()
	str := make([]byte, 0)
	for {
		buf := make([]byte, 1024)
		if n, err := conn.Read(buf); err != nil {
			panic(errorCode.TCPREADERROR)
		} else {
			if string(buf[0:n])=="\r\n"  {
				p.publicChan <- Msg{
					connIndex: index,
					msg: string(str),
				}
				str = str[0:0]
			} else {
				switch {
				case buf[n-1] == 8:
					str = str[0:len(str)-1]
				case buf[n-1] == 3:
					fmt.Printf("用户%v退出\n", index)
					return
				default:
					str = append(str, buf[0:n]...)
				}
			}
		}
	}
}

func (p *TcpWorker) GetChan() chan Msg {
	return p.publicChan
}

func InitTcp(port int) *TcpWorker {
	worker := new(TcpWorker)
	worker.startTcpListener(port)
	return worker
}
