package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type WisePaasStruct struct {
	EnsaasServices string
	Datacenter     string
	Workspace      string
	Cluster        string
	Namespace      string
	SsoApiUrl      string
	External       string
	IsPrivateCloud bool //是否私有雲
	IsCloud        bool //是否雲端，或地端
	// #may deprecated
	AppID       string
	ClientName  string
	ServiceName string
	SkipTLS     string
}

type MongoDBStruct struct {
	URL          string `validate:"required"`
	Database     string `validate:"required"`
	Username     string `validate:"required"`
	Password     string `validate:"required"`
	PasswordFile string
	AuthSource   string
	URI          string
	//---------- dev mongoDB(iiipivot) ----------
	DevURL      string
	DevDatabase string
	DevUsername string
	DevPassword string
	DevURI      string
}

type InfluxDBStruct struct {
	InfluxdbURL          string
	InfluxdbDatabase     string
	InfluxdbUsername     string
	InfluxdbPassword     string
	InfluxdbPasswordFile string
}

type PostgresStruct struct {
	PostgresURL          string
	PostgresDatabase     string
	PostgresUsername     string
	PostgresPassword     string
	PostgresPasswordFile string
}

type RedisStruct struct {
	RedisURL          string
	RedisPassword     string
	RedisPort         string
	RedisPasswordFile string
	RedisURLPort      string
}

type IfpStruct struct {
	URL           string `validate:"required"`
	Token         string
	AdminUsername string
	AdminPassword string
	AppSecretFile string
	OutboundURL   string
	Status        string
}

type EtcdStruct struct {
	EtcdEnable   string
	EtcdHost     string `validate:"required_if=EtcdEnable true"`
	EtcdPort     string `validate:"required_if=EtcdEnable true"`
	EtcdUsername string
	EtcdPassword string
	APIUrl       string
	UIUrl        string
}

type MinioStruct struct {
	Enable    string `json:"STORAGE_ENABLE"`
	Port      string `validate:"required_if=Enable true" json:"STORAGE_API_PORT"`
	Endpoint  string `validate:"required_if=Enable true" json:"STORAGE_API_URL"`
	Accesskey string `validate:"required_if=Enable true" json:"STORAGE_ACCESS_KEY"`
	Secretkey string `validate:"required_if=Enable true" json:"STORAGE_SECRET_KEY"`
	UseV4     string `validate:"required_if=Enable true" json:"STORAGE_USE_V4"`
}
