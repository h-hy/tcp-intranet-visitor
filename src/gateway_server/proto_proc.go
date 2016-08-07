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
	"libnet"
	"protocol"
	log "user_log"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type ProtoProc struct {
	msgServer *MsgServer
}

func checkService() {

}

func NewProtoProc(msgServer *MsgServer) *ProtoProc {
	return &ProtoProc{
		msgServer: msgServer,
	}
}

func (self *ProtoProc) procHeartbeat(cmd protocol.Cmd, session *libnet.Session) error {
	session.State.(*base.SessionState).Alive = true
	resp := protocol.NewCmdSimple(protocol.HEARTBEAT_CMD)
	session.Send(libnet.Json(resp))

	return nil
}

func (self *ProtoProc) procDefineLink(cmd protocol.Cmd, session *libnet.Session) error {

	Type := cmd.GetArgs()[0]
	UUID := cmd.GetArgs()[1]
	self.msgServer.scanSessionMutex.Lock()
	if Type == protocol.DEFINE_LINK_COMMAND_CMD {
		self.msgServer.sessions[UUID] = session
		self.msgServer.sessions[UUID].State = base.NewSessionState(UUID, "0")
		log.Infof("网关客户端 %s 已注册", UUID)
		go self.msgServer.checkClient()
	} else if Type == protocol.DEFINE_LINK_DATA_CMD {
		log.Infof("通道已建立：" + cmd.GetArgs()[2])
		self.msgServer.sessions[UUID+"_"+cmd.GetArgs()[2]] = session
		self.msgServer.sessions[UUID+"_"+cmd.GetArgs()[2]].State = base.NewSessionState(UUID, cmd.GetArgs()[2])
	}
	self.msgServer.scanSessionMutex.Unlock()
	return nil
}
