/*
 * Copyright 2022 CloudWeGo Authors
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

package main

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"

	nacosUtil "hertz_demo/middleware/nacos"
)

func main() {

	cli := nacosUtil.GetNewNamingClient()

	addr := "127.0.0.1:8001"

	// sc := []constant.ServerConfig{
	// 	*constant.NewServerConfig("127.0.0.1", 8848),
	// }

	// cc := constant.ClientConfig{
	// 	NamespaceId:         "d165c4e0-c76b-42cc-81bb-b173ffa5ad3d",
	// 	TimeoutMs:           5000,
	// 	NotLoadCacheAtStart: true,
	// 	LogDir:              "./tmp/nacos/log",
	// 	CacheDir:            "./tmp/nacos/cache",
	// 	LogLevel:            "info",
	// }

	// cli, err := clients.NewNamingClient(
	// 	vo.NacosClientParam{
	// 		ClientConfig:  &cc,
	// 		ServerConfigs: sc,
	// 	},
	// )

	// if err != nil {
	// 	panic(err)
	// }

	r := nacos.NewNacosRegistry(cli)

	// server.Default() creates a Hertz with recovery middleware.
	// If you need a pure hertz, you can use server.New()
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "supermancell.intergration",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello supermancell.intergration!")
	})

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.Spin()

}
