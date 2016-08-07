//
// Copyright 2014 Hong Miao. All Rights Reserved.
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

package base

import (
	"libnet"
)

type SessionMap map[string]*libnet.Session

type AckMap map[string]map[string]string

type SessionState struct {
	Alive bool
	UUID  string
	Index string
}

func NewSessionState(UUID string, Index string) *SessionState {
	return &SessionState{
		Alive: true,
		UUID:  UUID,
		Index: Index,
	}
}

type Config interface {
	LoadConfig(configfile string) (*Config, error)
}
