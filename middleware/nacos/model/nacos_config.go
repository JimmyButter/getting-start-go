package model

type NacosConfig struct {
	Ip        string
	Path      string
	Port      int
	Scheme    string
	Namespace string
	Timeout   int
	Cache     bool
}
