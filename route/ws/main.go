// Copyright 2023 CloudWeGo Authors
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
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/reverseproxy"
)

var (
	backendURL = "ws://127.0.0.1:7001/echo"
	proxyAddr  = "127.0.0.1:7000"
)

func main() {
	// websocket reverse proxy
	wsrp := reverseproxy.NewWSReverseProxy(backendURL)
	ps := server.Default(server.WithHostPorts(proxyAddr))
	ps.GET("/ws", wsrp.ServeHTTP)
	ps.Spin()
}
