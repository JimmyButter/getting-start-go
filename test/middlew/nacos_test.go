package middlew_test

import (
	"fmt"
	"hertz_demo/middleware/nacos"
	"testing"
)


func TestNacos(t *testing.T) {
	nacos.GetNacosConfigInstance()
	fmt.Println(nacos.OkexConfigYAML)
}