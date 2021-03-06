/*
 * Copyright 2018 The ThunderDB Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package route provides DHT routing functions
package route

import (
	"net"

	log "github.com/sirupsen/logrus"
	"github.com/thunderdb/ThunderDB/rpc"
)

// InitDHTServer install DHTService payload to RPC server, also set listener
func InitDHTServer(l net.Listener) (server *rpc.Server, err error) {
	server, err = rpc.NewServerWithService(rpc.ServiceMap{"DHT": NewDHTService()})
	if err != nil {
		log.Fatal(err)
		return
	}
	server.SetListener(l)
	return
}
