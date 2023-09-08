package test

import (
	"encoding/json"
	"fmt"
	"supermancell/src/model"
	"testing"
)

func TestVoToJson(t *testing.T) {

	instSet := []string{"BTC-USDT-SWAP", "ETH-USDT-SWAP"}

	op := model.OpVo{Op: "subscribe"}

	for _, instId := range instSet {
		args := model.OpArgs{Channel: "candle1H", InstId: instId}
		op.Args = append(op.Args, args)
	}

	opJson, err := json.Marshal(op)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(opJson))

}
