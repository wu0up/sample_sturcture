package models

import (
	// h "sample/pkg/helpers"
	"sample/pkg/module/sensorinfo/dto"
)

type SensorConfigModel struct {
	dto.GetSensorConfigDto
}

// func (m *MemberModel) TableName() string {
// 	return "member"
// }

// func (m *MemberModel) Check() (bool, string) {
// 	if !h.CheckUuid(m.MemberId) {
// 		return false, "invalid uuid"
// 	}
// 	if len(m.Name) == 0 {
// 		return false, "invalid name"
// 	}
// 	if !h.IsValidEmail(m.Email) {
// 		return false, "invalid email"
// 	}
// 	return true, ""
// }
