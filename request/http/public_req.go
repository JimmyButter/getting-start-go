package http

import (
	"context"
	"hertz_demo/middleware/nacos"
	"hertz_demo/middleware/nacos/model"
	"log"

	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
	"github.com/amir-the-h/okex/models/publicdata"
	publicReq "github.com/amir-the-h/okex/requests/rest/public"
)

func GetInstruments() []*publicdata.Instrument {
	strConfig := nacos.GetConfig("okex", "DEFAULT_GROUP")
	var config model.OkexConfig
	nacos.GetDecode(strConfig, &config)
	dest := okex.NormalServer // The main API server
	ctx := context.Background()
	client, err := api.NewClient(ctx, config.ApiKey, config.SecretKey, config.Passphrase, dest)
	if err != nil {
		log.Fatalln(err)
	}

	reqInstruments := &publicReq.GetInstruments{
		InstType: "SWAP",
	}

	rspInstrument, err := client.Rest.PublicData.GetInstruments(*reqInstruments)
	if err != nil {
		log.Fatalln(err)
	}

	return rspInstrument.Instruments
}
