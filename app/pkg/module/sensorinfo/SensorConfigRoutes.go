package sensorinfo

import (
	"github.com/gin-gonic/gin"
	_ "sample/docs"
)

func SetRoute(g *gin.Engine) { //baseRoute ==api
	sensorConfig := g.Group("/api/sensorConfig")
	// controller := InitSensorConfigController()
	// sensorConfigService := NewSensorConfigService()
	controller := NewSensorConfigController()
	// jwtMiddleWare := middleWare.InitJwtMiddleWare()  //因為api沒有使用驗證

	sensorConfig.GET("/serviceType", controller.GetServiceType)
	// sensorConfig.GET("/sensorList", controller.GetSensorList)

	// // GetBuildingTopoList godoc
	// //
	// //	@Summary		Get building topo
	// //	@Description	Get building topo
	// //	@Tags			sensorConfig
	// //	@Accept			json
	// //	@Produce		json
	// //	@Success		200	{array}
	// //	@Failure		400	{object}	httputil.HTTPError
	// //	@Failure		404	{object}	httputil.HTTPError
	// //	@Failure		500	{object}	httputil.HTTPError
	// //	@Router			/api/sensorConfig/buildingTopoList [get]
	// sensorConfig.GET("/buildingTopoList", controller.GetBuildingList)

	// // GetSensorConfig godoc
	// //
	// //	@Summary		Get sensor config
	// //	@Description	Get sensor config
	// //	@Tags			sensorConfig
	// //	@Accept			json
	// //	@Produce		json
	// //	@Success		200	{array}
	// //	@Failure		400	{object}	httputil.HTTPError
	// //	@Failure		404	{object}	httputil.HTTPError
	// //	@Failure		500	{object}	httputil.HTTPError
	// //	@Router			/api/sensorConfig/Config [get]
	// sensorConfig.GET("/Config", controller.GetSensorConfig)

	// // CreateServiceType godoc
	// //
	// //	@Summary		Create sensor config
	// //	@Description	Create sensor config
	// //	@Tags			sensorConfig
	// //	@Accept			json
	// //	@Produce		json
	// //	@Success		200	{object}
	// //	@Failure		400	{object}	httputil.HTTPError
	// //	@Failure		404	{object}	httputil.HTTPError
	// //	@Failure		500	{object}	httputil.HTTPError
	// //	@Router			/api/sensorConfig/Config [post]
	// sensorConfig.POST("/Config", controller.CreateSensorConfig)

	// // PutSensorConfig godoc
	// //
	// //	@Summary		Put sensor config
	// //	@Description	Put sensor config
	// //	@Tags			sensorConfig
	// //	@Accept			json
	// //	@Produce		json
	// //	@Success		200	{object}
	// //	@Failure		400	{object}	httputil.HTTPError
	// //	@Failure		404	{object}	httputil.HTTPError
	// //	@Failure		500	{object}	httputil.HTTPError
	// //	@Router			/api/sensorConfig/Config [put]
	// sensorConfig.PUT("/Config", controller.DeleteSensorConfig)

	// // DeleteSensorConfig godoc
	// //
	// //	@Summary		Delete sensor config
	// //	@Description	Delete sensor config
	// //	@Tags			sensorConfig
	// //	@Accept			json
	// //	@Produce		json
	// //	@Success		200	{object}
	// //	@Failure		400	{object}	httputil.HTTPError
	// //	@Failure		404	{object}	httputil.HTTPError
	// //	@Failure		500	{object}	httputil.HTTPError
	// //	@Router			/api/sensorConfig/Config [delete]
	// sensorConfig.DELETE("/Config", controller.UpdateSensorConfig)
}
