// Copyright 2021 CloudWeGo Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/hertz-contrib/registry/nacos/v2"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/hertz-contrib/reverseproxy"
	"github.com/cloudwego/hertz/pkg/protocol"
)


func main() {
	cli, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	sc := []constant.ServerConfig{
		*constant.NewServerConfig("127.0.0.1", 8848),
	}
	cc := constant.ClientConfig{
		NamespaceId:         "d165c4e0-c76b-42cc-81bb-b173ffa5ad3d",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "info",
	}

	nacosCli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	r := nacos.NewNacosResolver(nacosCli)
	cli.Use(sd.Discovery(r))

	h := server.New(server.WithHostPorts(":8000"))
	proxy, _ := reverseproxy.NewSingleHostReverseProxy("http://hertz.test.hello")
	proxy.SetClient(cli)
	proxy.SetDirector(func(req *protocol.Request) {
		req.SetRequestURI(string(reverseproxy.JoinURLPath(req, proxy.Target)))
		req.Header.SetHostBytes(req.URI().Host())
		req.Options().Apply([]config.RequestOption{config.WithSD(true)})
	})
	h.GET("/hello", proxy.ServeHTTP)
	h.Spin()
	// for i := 0; i < 10; i++ {
	// 	status, body, err := cli.Get(context.Background(), nil, "http://hertz.test.hello/ping", config.WithSD(true))
	// 	if err != nil {
	// 		hlog.Fatal(err)
	// 	}
	// 	hlog.Infof("code=%d,body=%s", status, string(body))
	// }
}
