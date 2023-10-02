package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sample/config/dto"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

var (
	// TaipeiTimeZone, _ = time.LoadLocation("Asia/Taipei")
	UTCTIMEZONE, _ = time.LoadLocation("UTC")
	// AppName           = "datafabric"
	// Debug             bool

	// 初始化環境變數
	SKIP_TLS            bool
	VIEWLINKSIMULATOR   bool
	GRAFANA_URL         string
	III_GARFANAUSER     string
	III_GARFANAPASSWORD string

	// WISEPAAS  dto.WisePaasStruct
	MONGODB  dto.MongoDBStruct
	INFLUXDB dto.InfluxDBStruct
	POSTGRES dto.PostgresStruct
	REDIS    dto.RedisStruct

// 	Ifp      = &IfpStruct{}
// 	Etcd     = &EtcdStruct{}
// 	Minio    = &MinioStruct{}
// 	Superset = &SupersetStruct{}
// 	Datahub  = &DatahubStruct{}
// 	Dis      = &DisStruct{}
)

func init() {
	envErr := godotenv.Load("config/env/.env") //針對dev時，讀取.env的環境變數，在production環境，可以註解
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}
	// if appName := os.Getenv("APP_NAME"); appName != "" {
	// 	AppName = appName
	// }
	// if debug := os.Getenv("DEBUG"); debug == "true" {
	// 	Debug = true
	// }

	fmt.Printf("---------- Set Env Start ----------\n")

	// WisePaas.SetEnv()
	MONGODB.SetEnv()
	INFLUXDB.SetEnv()
	POSTGRES.SetEnv()
	REDIS.SetEnv()
	// Minio.SetEnv()
	// Superset.SetEnv()
	// Datahub.SetEnv()
	// Dis.SetEnv()
	fmt.Printf("---------- Set Env Done ----------\n")
}

func (m *dto.MongoDBStruct) SetEnv() {
	m.URL = os.Getenv("MONGODB_URL")
	m.Database = os.Getenv("MONGODB_DATABASE")
	m.Username = os.Getenv("MONGODB_USERNAME")
	m.PasswordFile = os.Getenv("MONGODB_PASSWORD_FILE")
	m.Password = os.Getenv("MONGODB_PASSWORD")
	m.AuthSource = os.Getenv("MONGODB_AUTH_SOURCE")
	if m.PasswordFile != "" {
		password, err := os.ReadFile(m.PasswordFile)
		if err != nil {
			log.Error("Read MongodbPasswordFile error:", err)
		} else {
			m.Password = strings.Trim(string(password), "\n")
		}
	}

	if m.AuthSource == "" {
		m.AuthSource = m.Database
	}
	m.URI = fmt.Sprintf("mongodb://%s:%s@%s/%s", m.Username, m.Password, m.URL, m.AuthSource)

	fmt.Printf("Env Mongodb: %+v\n\n", m)

}

func (i *dto.InfluxDBStruct) SetEnv() {
	i.InfluxdbURL = os.Getenv("INFLUX_URL")
	i.InfluxdbDatabase = os.Getenv("INFLUX_DATABASE")
	i.InfluxdbUsername = os.Getenv("INFLUX_USERNAME")
	i.InfluxdbPasswordFile = os.Getenv("INFLUX_PASSWORD_FILE")
	i.InfluxdbPassword = os.Getenv("INFLUX_PASSWORD")
}

func (r *dto.RedisStruct) SetEnv() {
	r.RedisURL = os.Getenv("REDIS_URL")
	r.RedisPassword = os.Getenv("REDIS_PASSWORD")
	r.RedisPort = os.Getenv("REDIS_PORT")
	r.RedisPasswordFile = os.Getenv("REDIS_PASSWORD_FILE")
	r.RedisURLPort = fmt.Sprintf("%s:%s", r.RedisURL, r.RedisPort)

}

func (p *dto.PostgresStruct) SetEnv() {
	p.PostgresURL = os.Getenv("Postgres_URL")
	p.PostgresPassword = os.Getenv("Postgres_PASSWORD")
	p.PostgresDatabase = os.Getenv("Postgres_DATABASE")
	p.PostgresPasswordFile = os.Getenv("Postgres_PASSWORD_FILE")
	p.PostgresUsername = os.Getenv("Postgres_USERNAME")

}
