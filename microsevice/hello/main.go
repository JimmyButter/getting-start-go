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
	"fmt"
	"os"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/registry/nacos"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {

	addr := "127.0.0.1:8888"

	for i, v := range os.Args {
		fmt.Println(i, v)
		if strings.Contains(v, "--port=") {
			strArr := strings.Split(v, "=")
			addr = strings.Replace(addr, "8888", strArr[1], 1)
		}
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

	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	r := nacos.NewNacosRegistry(cli)

	// server.Default() creates a Hertz with recovery middleware.
	// If you need a pure hertz, you can use server.New()
	h := server.Default(
		server.WithHostPorts(addr),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "hertz.test.hello",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}),
	)

	h.GET("/hello", func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "Hello hertz!")
	})

	h.GET("/hello/hertz/:version", func(ctx context.Context, c *app.RequestContext) {
		// 获取路由参数
		// curl http://localhost:8888/hello/hertz/v1.0
		version := c.Param("version")
		c.String(consts.StatusOK, "Hello hertz! %s", version)
	})

	h.POST("/hello/form", func(c context.Context, ctx *app.RequestContext) {
		// post 表单示例
		// curl --location 'http://localhost:8888/hello/form' --form 'age="10"'
		age, _ := ctx.GetPostForm("age")
		ctx.JSON(consts.StatusOK, utils.H{"age": age})
	})

	h.GET("/hello/query", func(c context.Context, ctx *app.RequestContext) {
		// 获取QueryString
		// curl http://localhost:8888/hello/query\?name\=mike
		name := ctx.Query("name")
		ctx.JSON(consts.StatusOK, utils.H{"name": name})

	})

	h.POST("/hello/body", func(c context.Context, ctx *app.RequestContext) {

		/**
		curl --location 'http://localhost:8888/hello/body' \
		--header 'Content-Type: application/json' \
		--data '{
			"age":10,
			"name":"zhangsan"
		}'
		**/
		type Person struct {
			Age  int    `json:"age"`
			Name string `json:"name"`
		}

		var p Person
		err := ctx.Bind(&p)
		if err != nil {
			ctx.JSON(consts.StatusBadRequest, utils.H{"msg": err.Error()})
		}

		ctx.JSON(consts.StatusOK, utils.H{"person": p})
	})

	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.Spin()
}
