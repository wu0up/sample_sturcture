package implement

import (
	"sample/pkg/repos/models"

	"gorm.io/gorm"
)

type SensorConfigRepo struct {
	// base.Repository
	DBConn *gorm.DB
}

func NewSensorConfigRepo(db *gorm.DB) *MemberRepo {
	return &MemberRepo{DBConn: db}
}

func (r *SensorConfigRepo) GetSensorList() (member []models.MemberModel, err error) {
	var model []models.MemberModel
	error := r.DBConn.Find(&model).Error
	if error != nil {
		return model, r.SystemError("member_repo_error :" + error.Error())
	}

	return model, error
}
