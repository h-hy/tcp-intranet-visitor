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
	"base"
	"flag"
	//	"fmt"
	//	"time"

	//	"base"
	"libnet"
	log "user_log"
)

func init() {
	flag.Set("alsologtostderr", "false")
	flag.Set("log_dir", "false")
}

var InputConfFile = flag.String("conf_file", "gateway_server.json", "input conf file name")

func handleSession(ms *MsgServer, session *libnet.Session) {
	session.Process(func(msg *libnet.InBuffer) error {
		err := ms.parseProtocol(msg.Data, session)
		if err != nil {
			log.Error(err.Error())
		}

		return nil
	})
}

func main() {
	flag.Parse()
	cfg := NewMsgServerConfig(*InputConfFile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}

	ms := NewMsgServer(cfg)

	ms.server, err = libnet.Listen(cfg.TransportProtocols, cfg.Listen)
	if err != nil {
		panic(err)
	}
	log.Infof("网关服务器在 %s 监听中，等待网关客户端连接...", ms.server.Listener().Addr().String())

	go ms.scanDeadSession()
	go ms.checkClient()
	//	remoteAddr := "127.0.0.1"
	//	remotePort := "8081"

	ms.server.Serve(func(session *libnet.Session) {
		session.AddCloseCallback(ms, func() {
			if session.State != nil {
				if session.State.(*base.SessionState).Index != "0" {
					this_index := session.State.(*base.SessionState).Index
					//					log.Info(this_index + "链接：受访者关闭链接")
					UUID := session.State.(*base.SessionState).UUID
					ms.clientChannels[UUID+"_"+this_index] <- []byte("close!!!!!!!!!!")
				} else {
					UUID := session.State.(*base.SessionState).UUID
					log.Infof("网关客户端 %s 已断开", UUID)
					ms.procOffline(UUID)
					ms.checkClient()
				}
			}

		})
		go handleSession(ms, session)
	})
}
