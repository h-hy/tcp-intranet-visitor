//
// Copyright 2014 Hong Miao (miaohong@miaohong.org). All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"net"
	//	"regexp"
	//	"strconv"
	//	"time"
	// "bytes"

	//	"base"
	// "github.com/oikomi/FishChatServer/common"
	"libnet"
	"protocol"
	log "user_log"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type ProtoProc struct {
	gatewayClient *GatewayClient
}

func NewProtoProc(gatewayClient *GatewayClient) *ProtoProc {
	return &ProtoProc{
		gatewayClient: gatewayClient,
	}
}
func connectRemote(remoteIp, remotePort string) (net.Conn, error) {
	conn, err := net.Dial("tcp", remoteIp+":"+remotePort)
	if err != nil {
		log.Error(err)
		return conn, err
	}
	return conn, nil
}

var Host map[string]string

func (self *ProtoProc) procHeartbeat(cmd protocol.Cmd, session *libnet.Session, ms string) error {
	self.gatewayClient.gatewayServerClientMap[ms].Alive = true
	return nil
}

func (self *ProtoProc) procMakeRequest(cmd protocol.Cmd, session *libnet.Session) error {
	//	var msgServer string
	//	IMEI := cmd.GetInfos()["IMEI"]
	index := cmd.GetArgs()[0]
	remoteIp := cmd.GetArgs()[1]
	remotePort := cmd.GetArgs()[2]
	log.Info(index + "链接：收到发起链接指令")

	connLan, err := connectRemote(remoteIp, remotePort)
	if err != nil {
		log.Fatal(err)
		return err
	}

	msgServer := session.Conn().RemoteAddr().String()
	log.Info(index + "准备连接服务器" + session.Conn().RemoteAddr().String())
	gateWayClient, err := connectGatewayServer(msgServer) //发起连接
	if err != nil {
		log.Error(err.Error())
		return err
	}
	//连接建立成功，开始发送通道订阅
	log.Info(index + "链接：连接服务器成功，发送链接说明")
	cmd2 := protocol.NewCmdSimple(protocol.DEFINE_LINK_CMD)
	cmd2.AddArg(protocol.DEFINE_LINK_DATA_CMD)
	cmd2.AddArg(self.gatewayClient.cfg.UUID)
	cmd2.AddArg(index)

	err = gateWayClient.Send(libnet.Json(cmd2))
	if err != nil {
		log.Error(err.Error())
		return err
	}
	go func(connWan net.Conn) {
		for {
			received := make([]byte, 20480)
			i, err := connWan.Read(received)
			if err != nil {
				log.Info(index + "链接：访问者关闭链接，准备关闭受访者链接")
				connLan.Close()
				break
			} else {
				log.Info(index + "：收到访问者数据")
				connLan.Write(received[0:i])
			}
		}
	}(gateWayClient.Conn())

	go func(connLan net.Conn) {
		for {
			received := make([]byte, 204800)
			i, err := connLan.Read(received)

			if err != nil {
				log.Info(index + "链接：受访者关闭链接，准备关闭访问者链接")
				gateWayClient.Close()
				break
			} else {
				log.Info(index + "链接：收到受访者数据")
				//				log.Info(index+"链接：收到受访者数据", string(received[0:i]))
				gateWayClient.Send(libnet.Bytes(received[0:i]))
			}
		}
	}(connLan)

	return nil
}
func connectGatewayServer(ms string) (*libnet.Session, error) {
	client, err := libnet.Dial("tcp", ms)
	if err != nil {
		log.Error(err.Error())
		// panic(err)
	}

	return client, err
}
func init() {
	Host = make(map[string]string)
}
