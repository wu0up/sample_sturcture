package sensorinfo

import (
	// "sample/pkg/base"
	"sample/pkg/helpers"

	"github.com/gin-gonic/gin"
	// "net/http"
)

type SensorConfigController struct {
	SensorConfigService SensorConfigService
	// base.Controller
}

func NewSensorConfigController() *SensorConfigController {
	return &SensorConfigController{
		// SensorConfigService: sensorConfigService,
	}
}

// GetServiceType godoc
//
//	@Summary		Get service type
//	@Description	Get service type
//	@Tags			sensorConfig
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}	dto.GetServiceTypeDto	"success"
//	@Router			/api/sensorConfig/serviceType [get]
func (c *SensorConfigController) GetServiceType(g *gin.Context) {
	response, err := c.SensorConfigService.GetServiceType()
	helpers.HandleError(g, response, err)
}

// // GetSensorList godoc
// //
// //	@Summary		Get sensor list
// //	@Description	Get sensor list
// //	@Tags			sensorConfig
// //	@Accept			json
// //	@Produce		json
// //	@Success		200	{array}
// //	@Failure		400	{object}	httputil.HTTPError
// //	@Failure		404	{object}	httputil.HTTPError
// //	@Failure		500	{object}	httputil.HTTPError
// //	@Router			/api/sensorConfig/sensorList [get]
// func (c *SensorConfigController) GetSensorList(g *gin.Context) {
// 	response, err := c.SensorConfigService.GetSensorList()
// 	helpers.HandleError(g, response, err)
// }

// func (c *SensorConfigController) GetBuildingList(g *gin.Context) {
// 	response, err := c.SensorConfigService.GetBuildingList()
// 	helpers.HandleError(g, response, err)
// }

// func (c *SensorConfigController) GetSensorConfig(g *gin.Context) {
// 	response, err := c.SensorConfigService.GetSensorConfig()
// 	helpers.HandleError(g, response, err)
// }

// func (c *SensorConfigController) CreateSensorConfig(g *gin.Context) {

// 	var sensorConfig dto.CreateSensorConfigDto
// 	// g.Bind(&SensorConfig) //因為在swaggo中有註解了

// 	response, err := c.SensorConfigService.CreateSensorConfig(&sensorConfig)
// 	helpers.HandleError(g, response, err)
// }

// func (c *SensorConfigController) DeleteSensorConfig(g *gin.Context) {

// 	id := g.Query("id")

// 	response, err := c.SensorConfigService.DeleteSensorConfig(&id)
// 	helpers.HandleError(g, response, err)
// }

// func (c *SensorConfigController) UpdateSensorConfig(g *gin.Context) {

// 	id := g.Query("id")

// 	response, err := c.SensorConfigService.UpdateSensorConfig(&id)
// 	helpers.HandleError(g, response, err)
// }
