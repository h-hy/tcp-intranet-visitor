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
	//	"fmt"
	//	"time"

	//	"base"

	log "user_log"
)

func init() {
	flag.Set("alsologtostderr", "false")
	flag.Set("log_dir", "false")
}

var InputConfFile = flag.String("conf_file", "gateway_client.json", "input conf file name")

func main() {
	flag.Parse()
	cfg := NewGatewayClientConfig(*InputConfFile)
	err := cfg.LoadConfig()
	if err != nil {
		log.Error(err.Error())
		return
	}

	ms := NewGatewayClient(cfg)

	ms.subscribeChannels() //连接到中转服务器

	ms.scanDeadSession()

}
