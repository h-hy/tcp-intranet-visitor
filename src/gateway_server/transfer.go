package main

import (
	"bytes"
	"fmt"
	"libnet"
	"net"
	"protocol"
	"regexp"
	"sync"
	"time"
	log "user_log"
)

func init() {
	Host = make(map[string]string)
	index = 1
}

var indexLook sync.Mutex
var index int

type Transfer struct {
	server      *MsgServer
	listener    net.Listener
	Client_UUID string
	RemoteAddr  string
	RemotePort  string
	LocalAddr   string
	LocalPort   string
	IsHttp      bool
	Running     bool
}

func (self *MsgServer) checkClient() error {
	self.TransferListCheckMutex.Lock()
	defer self.TransferListCheckMutex.Unlock()
	for _, ms := range self.cfg.TransferList {
		if ms.server == nil {
			ms.server = self
		}
		if ms.Running == false {
			if self.sessions[ms.Client_UUID] != nil {
				log.Info("begin")
				log.Info("SET true")
				ms.Running = true
				log.Info(ms.LocalPort, ms.Running)
				go func(ms *Transfer) {
					go ms.Begin()
				}(ms)
			}
		} else {
			if self.sessions[ms.Client_UUID] == nil {
				log.Info("ms.Stop()")
				ms.Stop()
				//				ms.Running = true

			}
		}
	}

	return nil
}
func (self *Transfer) Begin() error {
	log.Info(self.LocalPort)
	var err error
	self.listener, err = net.Listen("tcp", ":"+self.LocalPort)
	if err != nil {
		log.Error(err.Error())
		self.Running = false
		return err
	}

	log.Infof("启用服务：本地端口：%s，远程地址：%s:%s，HTTP模式：%t", self.LocalPort, self.RemoteAddr, self.RemotePort, self.IsHttp)
	for {
		conn, err := self.listener.Accept()
		if err != nil {
			log.Error(err.Error())
			return err
		}
		go self.handleConnection(conn)
	}
	log.Info(self.LocalPort, self.Running)
	return nil
}
func (self *Transfer) Stop() error {

	self.listener.Close()
	self.Running = false
	return nil
}

var Host map[string]string

func (self *Transfer) handleConnection(connWAN net.Conn) error {
	if Host[self.RemoteAddr] == "" {
		ip, err := net.LookupIP(self.RemoteAddr)
		log.Infof("获取IP地址：", ip)
		if err != nil {
			log.Error(err.Error())
			return err
		}
		Host[self.RemoteAddr] = ip[0].String()
	}
	indexLook.Lock()
	index++
	var this_index string
	this_index = fmt.Sprintf("%d", index)
	log.Infof("建立通道中：" + this_index)
	indexLook.Unlock()
	receive_connLAN := make(chan []byte)
	if self.server.sessions[self.Client_UUID] != nil {
		cmd := protocol.NewCmdSimple(protocol.MAKE_REQUEST_CMD)
		self.server.clientChannels[self.Client_UUID+"_"+this_index] = receive_connLAN
		cmd.AddArg(this_index)
		cmd.AddArg(Host[self.RemoteAddr])
		cmd.AddArg(self.RemotePort)
		self.server.sessions[self.Client_UUID].Send(libnet.Json(cmd))
		// log.Info(this_index + "链接：发送MAKE_REQUEST指令")
	} else {
		// log.Info(this_index + "链接：发送MAKE_REQUEST指令失败")
	}

	receive_connWAN := make(chan []byte)
	go func(conn net.Conn) {
		for {
			received := make([]byte, 2048000)
			i, err := conn.Read(received)
			if err != nil {
				log.Infof("关闭通道：" + this_index)
				// log.Info(this_index + "链接：访问者Read返回异常")
				if self.server.sessions[self.Client_UUID+"_"+this_index] != nil {
					self.server.sessions[self.Client_UUID+"_"+this_index].Close()
					// log.Info(this_index + "链接：关闭对应受访者链接")
				}
				conn.Close()
				//				log.Error(err.Error())
				break
			} else {
				// log.Info(this_index + "链接：接收访问者数据正确")
				if self.IsHttp {
					received_string := string(received[0:i])
					// log.Info(this_index+"链接：接收访问者数据：", received_string)
					re, _ := regexp.Compile("Host:(.*)")
					received_string = re.ReplaceAllString(received_string, "Host: "+self.RemoteAddr)

					receive_connWAN <- []byte(received_string)
				} else {

					receive_connWAN <- []byte(received[0:i])
				}
			}
		}
	}(connWAN)
	first := true
	for {
		select {
		case x := <-receive_connWAN:
			var times int = 1
			if first == true {
				times = 16
			}
			for i := 0; i < times; i++ {
				if self.server.sessions[self.Client_UUID+"_"+this_index] != nil {
					self.server.sessions[self.Client_UUID+"_"+this_index].Conn().Write(x)
					break
				} else {
					if first == true {
						time.Sleep(time.Millisecond * 100)
					} else {
						// log.Info(this_index + "链接收到访问者请求，但找不到受访者")
					}
				}
			}
			// log.Info(this_index + "链接收到访问者请求，成功找到受访者")
			first = false
		case y := <-receive_connLAN:
			if bytes.Equal(y, []byte("close!!!!!!!!!!")) {
				// log.Info(this_index + "链接：受访者关闭链接，准备关闭对应访问者链接")
				connWAN.Close()
			} else {
				// log.Info(this_index + "链接收到受访者返回数据")

				if self.IsHttp {
					received_string := string(y)
					//					re, _ := regexp.Compile("<a HREF=\"http://" + self.RemoteAddr + "(.*)>")
					//					received_string = re.ReplaceAllString(received_string, "<a HREF=\"http://127.0.0.1$1>")
					//					re, _ = regexp.Compile("<a href=\"http://" + self.RemoteAddr + "(.*)>")
					//					received_string = re.ReplaceAllString(received_string, "<a href=\"http://127.0.0.1$1>")
					//					re, _ = regexp.Compile("<A href=\"http://" + self.RemoteAddr + "(.*)>")
					//					received_string = re.ReplaceAllString(received_string, "<a href=\"http://127.0.0.1$1>")
					re, _ := regexp.Compile("Location: http://" + self.RemoteAddr + "(.*)")
					received_string = re.ReplaceAllString(received_string, "Location: http://"+self.LocalAddr+":"+self.LocalPort+"$1")
					re, _ = regexp.Compile("Set-Cookie: (.*)domain=.*?;(.*)")
					received_string = re.ReplaceAllString(received_string, "Set-Cookie: $1 $2")

					connWAN.Write([]byte(received_string))
				} else {

					connWAN.Write(y)
				}
			}
		}
	}
	return nil
}
