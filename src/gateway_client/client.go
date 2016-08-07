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
	"errors"
	// "github.com/oikomi/FishChatServer/common"
	"libnet"
	"protocol"
	log "user_log"
)

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

type GatewayClient struct {
	cfg                    *GatewayClientConfig
	sessions               base.SessionMap
	scanSessionMutex       sync.Mutex
	msgServerClientMutex   sync.Mutex
	gatewayServerClientMap map[string]*gatewayClientState
	readMutex              sync.Mutex // multi client session may ask for REDIS at the same time
}

func NewGatewayClient(cfg *GatewayClientConfig) *GatewayClient {
	return &GatewayClient{
		cfg:                    cfg,
		sessions:               make(base.SessionMap),
		gatewayServerClientMap: make(map[string]*gatewayClientState),
	}
}

type gatewayClientState struct {
	Session *libnet.Session
	Alive   bool
}

/*
   用于反复检测没有连接成功的消息服务器，进行重连
*/
func (self *GatewayClient) subscribeChannels() error {
	log.Info("start to GatewayClient")
	self.msgServerClientMutex.Lock()
	defer self.msgServerClientMutex.Unlock()
	for _, ms := range self.cfg.GatewayServerList {
		if self.gatewayServerClientMap[ms] != nil {
			//已经创建过链接并且链接正常
			continue
		}
		gateWayClient, err := self.connectGatewayServer(ms) //发起连接
		if err != nil {
			log.Error(err.Error())
			go self.subscribeChannels()
			continue
		}
		//连接建立成功，开始发送通道订阅
		log.Info("connectGatewayServer ok ,sending..")
		cmd := protocol.NewCmdSimple(protocol.DEFINE_LINK_CMD)
		cmd.AddArg(protocol.DEFINE_LINK_COMMAND_CMD)
		cmd.AddArg(self.cfg.UUID)

		err = gateWayClient.Send(libnet.Json(cmd))
		if err != nil {
			log.Error(err.Error())
			go self.subscribeChannels()
			continue
		}
		log.Info("send ok")
		//通道订阅发送成功
		self.gatewayServerClientMap[ms] = new(gatewayClientState)
		self.gatewayServerClientMap[ms].Session = gateWayClient
		self.gatewayServerClientMap[ms].Alive = true

		//开始处理 消息服务器-> 接入服务器 的数据
		go func(ms string) {
			// go self.removeMsgServer(ms)
			err := self.handleGatewayServerClient(gateWayClient)
			log.Infof("err=%s", err)
			if err != nil {
				delete(self.gatewayServerClientMap, ms)
				log.Info("delete ok")
			}
			go self.subscribeChannels()
		}(ms)
	}
	return nil
}

func (self *GatewayClient) handleGatewayServerClient(msc *libnet.Session) error {
	err := msc.Process(func(msg *libnet.InBuffer) error {
		var c protocol.CmdSimple
		ms := msc.Conn().RemoteAddr().String()
		if self.gatewayServerClientMap[ms] == nil {
			log.Error(ms + " not exist")
			return errors.New(ms + " not exist")
		}
		err := json.Unmarshal(msg.Data, &c)
		if err != nil {
			log.Error("error:", err)
			return err
		}

		pp := NewProtoProc(self)
		log.Infof("c.GetCmdName()=%s\n\n", c.GetCmdName())
		switch c.GetCmdName() {
		case protocol.MAKE_REQUEST_CMD: //服务器要求客户端发起链接
			err = pp.procMakeRequest(&c, msc)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		case protocol.HEARTBEAT_CMD: //服务器要求客户端发起链接
			err = pp.procHeartbeat(&c, msc, ms)
			if err != nil {
				log.Error("error:", err)
				return err
			}
		}
		return nil
	})
	return err
}

func (self *GatewayClient) connectGatewayServer(ms string) (*libnet.Session, error) {
	client, err := libnet.Dial("tcp", ms)
	if err != nil {
		log.Error(err.Error())
		// panic(err)
	}

	return client, err
}

func (self *GatewayClient) scanDeadSession() {
	timer := time.NewTicker(self.cfg.ScanDeadSessionTimeout * time.Second)
	for {
		select {
		case <-timer.C:
			self.msgServerClientMutex.Lock()
			for ms, client := range self.gatewayServerClientMap {
				if client.Alive == false {
					self.procOffline(ms)
				} else {
					cmd := protocol.NewCmdSimple(protocol.HEARTBEAT_CMD)
					client.Session.Send(libnet.Json(cmd))
					client.Alive = false
				}
			}
			self.msgServerClientMutex.Unlock()
		}
	}
}

func (self *GatewayClient) procOnline(ID string) {
}

func (self *GatewayClient) procOffline(ms string) {
	// load all the topic list of this user
	if self.gatewayServerClientMap[ms] != nil {
		self.gatewayServerClientMap[ms].Session.Close()
		delete(self.gatewayServerClientMap, ms)
	}
}
