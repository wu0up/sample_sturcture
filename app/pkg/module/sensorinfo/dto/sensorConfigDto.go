package dto

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type CreateSensorConfigDto struct {
	Host        string     `validate:"required"`
	ServiceName string     `validate:"required"`
	ServiceType string     `validate:"required"`
	UserName    string     `validate:"required"`
	Password    string     `validate:"required"`
	CreatedAt   *time.Time `json:"omitempty"`
	UpdatedAt   *time.Time
}

type GetSensorConfigDto struct {
	Host        string     `gorm:"column:Host"`
	ServiceName string     `gorm:"column:ServiceName"`
	ServiceType string     `gorm:"column:ServiceType"`
	UserName    string     `gorm:"column:UserName"`
	Password    string     `json:"-"` //要求是想要JSON格式返回用户数据，并且不希望其中包含有Password字段
	CreatedAt   *time.Time `gorm:"column:CreatedAt"`
	UpdatedAt   *time.Time `gorm:"column:UpdatedAt"`
}

type GetServiceTypeDto struct {
	ServiceType string
}

// ValidateGetSensorConfig 用于验证 GetSensorConfig 结构体
func ValidateGetSensorConfig(CreateSensorConfigDto CreateSensorConfigDto) error {
	validate := validator.New()
	return validate.Struct(CreateSensorConfigDto)
}
