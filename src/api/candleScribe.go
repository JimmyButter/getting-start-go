package api

import (
	"encoding/json"
	"log"
	"net/http"
	"supermancell/src/model"
	"supermancell/src/service"

	"github.com/gin-gonic/gin"
)

func CandleScribe(ctx *gin.Context) {
	instSet := []string{"BTC-USDT-SWAP", "ETH-USDT-SWAP"}
	op := model.OpVo{Op: "subscribe"}

	for _, instId := range instSet {
		args := model.OpArgs{Channel: "candle1H", InstId: instId}
		op.Args = append(op.Args, args)
	}

	opJson, err := json.Marshal(op)
	if err != nil {
		log.Println(err)
	}

	service.ChBusinessSend(string(opJson))

	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
