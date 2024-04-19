package nacos

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
)

func GetNacosConfigInstance() {
	clientconfig := constant.ClientConfig{
		NamespaceId:         "d165c4e0-c76b-42cc-81bb-b173ffa5ad3d",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "./tmp/nacos/log",
		CacheDir:            "./tmp/nacos/cache",
		LogLevel:            "info",
	}
	serverConfig := []constant.ServerConfig{
		{
			IpAddr:      "127.0.0.1",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig":  clientconfig,
		"serverConfigs": serverConfig,
	})
	if err != nil {
		fmt.Println("连接nacos")
	}
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "okex",
		Group:  "DEFAULT_GROUP",
	})

	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("err:", err)
	fmt.Println("content")
	getConfig(content)

	//增加监听配置
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "hello",
		Group:  "DEFAULT_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ",dataId:" + dataId + ",data:" + data)
			getConfig(data)
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	//time.Sleep(time.Second * 60)

}

func getConfig(date string) {

	mapConfig := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(date), &mapConfig)
	if err != nil {
		fmt.Printf("Error unmarshalling YAML: %v\n", err)
		return // or handle the error as appropriate for your application
	}

	mapstructure.Decode(mapConfig, &OkexConfigYAML)
	fmt.Println("OKEX Config:", OkexConfigYAML)
}
