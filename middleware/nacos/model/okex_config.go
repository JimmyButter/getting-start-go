package model

type OkexConfig struct {
	ApiKey     string `json:"apiKey"`
	SecretKey  string `json:"secretKey"`
	Passphrase string `json:"passphrase"`
}
