package parser

import (
	// "./xintop"
	"github.com/gin-gonic/gin"
)

func SetRoute(g *gin.Engine) { //baseRoute ==api
	parserRoute := g.Group("/api/parser")
	// controller := InitSensorConfigController()
	// sensorConfigService := NewSensorConfigService()
	// parser := NewParserController()
	// jwtMiddleWare := middleWare.InitJwtMiddleWare()  //因為api沒有使用驗證

	parserRoute.POST("/XINTOP", parserPayload)
}
