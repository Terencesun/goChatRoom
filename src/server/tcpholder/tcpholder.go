package tcpholder

import (
	"fmt"
	"net"
	errorCode "server/errorcode"
	"strconv"
	"strings"
)

type Msg struct {
	ConnIndex int
	Msg []byte
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

func (p *TcpWorker) GetConnMap() map[int]net.Conn {
	return p.connMap
}

func (p *TcpWorker) removeConn(index int) {
	delete(p.connMap, index)
}

func (p *TcpWorker) process(index int, conn net.Conn)  {
	defer func() {
		if err := conn.Close(); err != nil {
			panic(errorCode.TCPCLOSEERROR)
		}
	}()
	for {
		buf := make([]byte, 1024)
		if n, err := conn.Read(buf); err != nil {
			if find := strings.Contains(err.Error(), "forcibly closed"); find {
				// 远程退出
				fmt.Printf("用户%v退出\n", index)
				p.removeConn(index)
				return
			} else {
				panic(errorCode.TCPREADERROR)
			}
		} else {
			p.publicChan <- Msg{
				ConnIndex: index,
				Msg: buf[0:n],
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
