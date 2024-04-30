package request_test

import (
	"context"
	"hertz_demo/middleware/nacos"
	"hertz_demo/middleware/nacos/model"
	"hertz_demo/request/http"
	"log"
	"testing"

	"github.com/amir-the-h/okex"
	"github.com/amir-the-h/okex/api"
	accountReq "github.com/amir-the-h/okex/requests/rest/account"
)

/**
 * Tests the API with GetBalance
 */
func TestAPIGetBalance(t *testing.T) {
	strConfig := nacos.GetConfig("okex", "DEFAULT_GROUP")
	var config model.OkexConfig
	nacos.GetDecode(strConfig, &config)
	dest := okex.NormalServer // The main API server
	ctx := context.Background()
	client, err := api.NewClient(ctx, config.ApiKey, config.SecretKey, config.Passphrase, dest)
	if err != nil {
		log.Fatalln(err)
	}

	reqBalance := &accountReq.GetBalance{
		Ccy: []string{"USDT"},
	}

	rspBalance, err := client.Rest.Account.GetBalance(*reqBalance)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(rspBalance.Balances[0].TotalEq)
}

func TestAPIGetInst(t *testing.T) {
	instruments := http.GetInstruments()
	for _, inst := range instruments {
		log.Println(inst.InstID)
	}

}
