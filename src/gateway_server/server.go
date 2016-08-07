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
	"encoding/json"
	"flag"
	"sync"
	"time"
	// "fmt"

	"base"
	// "github.com/oikomi/FishChatServer/common"
	"libnet"
	"protocol"
	log "user_log"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type MsgServer struct {
	cfg                    *MsgServerConfig
	sessions               base.SessionMap
	server                 *libnet.Server
	scanSessionMutex       sync.Mutex
	readMutex              sync.Mutex // multi client session may ask for REDIS at the same time
	clientChannels         map[string]chan []byte
	transferMap            map[string]Transfer
	TransferListCheckMutex sync.Mutex
}

func NewMsgServer(cfg *MsgServerConfig) *MsgServer {
	return &MsgServer{
		cfg:            cfg,
		sessions:       make(base.SessionMap),
		server:         new(libnet.Server),
		clientChannels: make(map[string]chan []byte),
		transferMap:    make(map[string]Transfer),
	}
}

func (self *MsgServer) scanDeadSession() {
	timer := time.NewTicker(self.cfg.ScanDeadSessionTimeout * time.Second)
	ttl := time.After(self.cfg.Expire * time.Second)
	for {
		select {
		case <-timer.C:
			go func() {
				for id, s := range self.sessions {
					self.scanSessionMutex.Lock()
					if (s.State).(*base.SessionState).Alive == false {
						log.Infof("服务器端关闭链接：", id)
						self.procOffline(id)
					} else {
						s.State.(*base.SessionState).Alive = false
					}
					self.scanSessionMutex.Unlock()
				}
			}()
		case <-ttl:
			break
		}
	}
}

func (self *MsgServer) procOnline(ID string) {
}

func (self *MsgServer) procOffline(ID string) {
	// load all the topic list of this user
	if self.sessions[ID] != nil {
		self.sessions[ID].Close()
		delete(self.sessions, ID)
	}
}

func (self *MsgServer) parseProtocol(cmd []byte, session *libnet.Session) error {
	if session.State != nil {
		self.scanSessionMutex.Lock()
		session.State.(*base.SessionState).Alive = true
		self.scanSessionMutex.Unlock()
		if session.State.(*base.SessionState).Index != "0" {
			this_index := session.State.(*base.SessionState).Index
			UUID := session.State.(*base.SessionState).UUID
			temp := make([]byte, len(cmd))
			copy(temp[0:], cmd[0:])
			self.clientChannels[UUID+"_"+this_index] <- temp
			return nil
		}
	}

	var c protocol.CmdSimple
	err := json.Unmarshal(cmd, &c)
	if err != nil {
		log.Error("error:", err)
		return err
	}
	pp := NewProtoProc(self)

	switch c.GetCmdName() {
	case protocol.DEFINE_LINK_CMD: //客户端连接声明
		err = pp.procDefineLink(&c, session)
		if err != nil {
			log.Error("error:", err)
			return err
		}
	case protocol.HEARTBEAT_CMD: //服务器要求客户端发起链接
		err = pp.procHeartbeat(&c, session)
		if err != nil {
			log.Error("error:", err)
			return err
		}
	}

	return err
}
