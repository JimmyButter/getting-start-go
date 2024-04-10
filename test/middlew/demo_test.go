package middlew_test
//包名规范 <you module>

import (
	"hertz_demo/middleware"
	"testing"
)

/**
* 这个示例展示了golang如何使用模块调用
*/
func TestDemo(t *testing.T) {
	//方法名规范：Test<your name>(t *testing.T)
	middleware.Demo()
}
