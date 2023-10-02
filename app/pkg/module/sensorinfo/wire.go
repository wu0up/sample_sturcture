//go:build wireinject
// +build wireinject

//go:generate echo Running go generate for wire.go
//go:generate wire

package sensorinfo

import (
	"github.com/google/wire"
)

// var SensorConfigRepo = wire.NewSet(implement.NewMemberRepo, wire.Bind(new(interfaces.MemberRepo), new(*implement.MemberRepo)))

func InitSensorConfigController() *SensorConfigController {
	wire.Build(
		NewSensorConfigController,
		// token.NewTokenService,
		NewSensorConfigService,
		SensorConfigRepo,
		// helpers.NewSqlSession,
	)
	return &SensorConfigController{}
}
