package model

type DatabaseConfig struct {
	Url      string `json:"url"`
	Db       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
}
