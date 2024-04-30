package middlew_test

import (
	"fmt"
	"hertz_demo/middleware/nacos"
	"hertz_demo/middleware/nacos/model"
	"testing"
)

func TestNacosOkex(t *testing.T) {
	context := nacos.GetConfig("okex", "DEFAULT_GROUP")
	fmt.Println(context)
	var config = &model.OkexConfig{}
	nacos.GetDecode(context, config)
	fmt.Println(config)
}

func TestNacosDatabase(t *testing.T) {
	context := nacos.GetConfig("database", "DEFAULT_GROUP")
	fmt.Println(context)
	var config = &model.DatabaseConfig{}
	nacos.GetDecode(context, config)
	fmt.Println(config)
}
