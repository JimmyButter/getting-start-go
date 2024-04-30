package nacos

import (
	"os"
	"sync"

	"hertz_demo/middleware/nacos/model"

	"github.com/mitchellh/mapstructure"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v2"
)

var instance config_client.IConfigClient
var namingInstance naming_client.INamingClient
var once sync.Once

func GetInstance() config_client.IConfigClient {
	if instance == nil {
		once.Do(func() {

			yamlBytes, err := os.ReadFile("../../config/nacos.yaml")
			if err != nil {
				panic(err)
			}

			nacosConfig := model.NacosConfig{}
			err = yaml.Unmarshal(yamlBytes, &nacosConfig)

			if err != nil {
				panic(err)
			}
			//GetDecode(string(yamlBytes), nacosConfig)

			clientconfig := constant.ClientConfig{
				NamespaceId:         nacosConfig.Namespace,
				TimeoutMs:           uint64(nacosConfig.Timeout),
				NotLoadCacheAtStart: nacosConfig.Cache,
				LogDir:              "./tmp/nacos/log",
				CacheDir:            "./tmp/nacos/cache",
				LogLevel:            "info",
			}
			serverConfig := []constant.ServerConfig{
				{
					IpAddr:      nacosConfig.Ip,
					ContextPath: nacosConfig.Path,
					Port:        uint64(nacosConfig.Port),
					Scheme:      nacosConfig.Scheme,
				},
			}

			configClient, err := clients.CreateConfigClient(map[string]interface{}{
				"clientConfig":  clientconfig,
				"serverConfigs": serverConfig,
			})

			if err != nil {
				panic(err)
			}

			instance = configClient
		})
	}

	return instance
}

func GetNewNamingClient() naming_client.INamingClient {
	if namingInstance == nil {
		once.Do(func() {

			yamlBytes, err := os.ReadFile("../../config/nacos.yaml")
			if err != nil {
				panic(err)
			}

			nacosConfig := model.NacosConfig{}
			err = yaml.Unmarshal(yamlBytes, &nacosConfig)

			if err != nil {
				panic(err)
			}
			//GetDecode(string(yamlBytes), nacosConfig)

			clientconfig := constant.ClientConfig{
				NamespaceId:         nacosConfig.Namespace,
				TimeoutMs:           uint64(nacosConfig.Timeout),
				NotLoadCacheAtStart: nacosConfig.Cache,
				LogDir:              "./tmp/nacos/log",
				CacheDir:            "./tmp/nacos/cache",
				LogLevel:            "info",
			}
			serverConfig := []constant.ServerConfig{
				{
					IpAddr:      nacosConfig.Ip,
					ContextPath: nacosConfig.Path,
					Port:        uint64(nacosConfig.Port),
					Scheme:      nacosConfig.Scheme,
				},
			}

			configClient, err := clients.CreateNamingClient(map[string]interface{}{
				"clientConfig":  clientconfig,
				"serverConfigs": serverConfig,
			})

			if err != nil {
				panic(err)
			}

			namingInstance = configClient
		})
	}

	return namingInstance
}

func GetConfig(dataId string, group string) string {
	client := GetInstance()
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	if err != nil {
		panic(err)
	}

	return content
}

func GetDecode(date string, output interface{}) interface{} {

	mapConfig := make(map[interface{}]interface{})

	err := yaml.Unmarshal([]byte(date), &mapConfig)
	if err != nil {
		panic(err)
	}

	mapstructure.Decode(mapConfig, output)
	return output
}

//增加监听配置
// err = configClient.ListenConfig(vo.ConfigParam{
// 	DataId: "hello",
// 	Group:  "DEFAULT_GROUP",
// 	OnChange: func(namespace, group, dataId, data string) {
// 		fmt.Println("group:" + group + ",dataId:" + dataId + ",data:" + data)
// 		getConfig(data)
// 	},
// })
// if err != nil {
// 	fmt.Println(err)
// 	return
// }
//time.Sleep(time.Second * 60)
