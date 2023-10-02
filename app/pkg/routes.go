package pkg

import (
	"github.com/gin-gonic/gin"
	// "sample/pkg/module/member"
	"github.com/rs/zerolog/log"
	"sample/pkg/module/sensorinfo"
	"sample/pkg/parser"
)

// func SetRouter(g *gin.Engine) {
// 	const baseGroup string = "api"
// 	member.SetRoute(g, baseGroup)
// }

func SetRouter(g *gin.Engine) {
	// const baseGroup string = "api"
	log.Info().Msg("in pkg route.go")
	sensorinfo.SetRoute(g)
	parser.SetRoute(g)
}
