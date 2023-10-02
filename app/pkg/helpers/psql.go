package helpers

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var SqlSession *gorm.DB

func NewSqlSession() *gorm.DB {
	return SqlSession
}

// func InitPsql() (err error, db *gorm.DB) {
// 	envErr := godotenv.Load(".env")
// 	if envErr != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	host := GetEnvStr("DB_HOST")
// 	port := GetEnvStr("DB_PORT")
// 	username := GetEnvStr("DB_USER")
// 	password := GetEnvStr("DB_PASSWORD")
// 	dbName := GetEnvStr("DB_NAME")
// 	maxIdleConn, _ := GetEnvInt("DB_MAX_IDLE_CONN")
// 	maxPoolConn, _ := GetEnvInt("DB_MAX_POOL_CONN")

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
// 		host,
// 		username,
// 		password,
// 		dbName,
// 		port,
// 	)
// 	SqlSession, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	sqlDB, err := SqlSession.DB()
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	sqlDB.SetMaxIdleConns(maxIdleConn)
// 	sqlDB.SetMaxOpenConns(maxPoolConn)
// 	return err, SqlSession
// }

func InitPsql() (err error, db *gorm.DB) {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		return envErr, nil // 返回.env文件加载错误，但不中断程序
	}

	host := GetEnvStr("DB_HOST")
	port := GetEnvStr("DB_PORT")
	username := GetEnvStr("DB_USER")
	password := GetEnvStr("DB_PASSWORD")
	dbName := GetEnvStr("DB_NAME")
	maxIdleConn, _ := GetEnvInt("DB_MAX_IDLE_CONN")
	maxPoolConn, _ := GetEnvInt("DB_MAX_POOL_CONN")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		host,
		username,
		password,
		dbName,
		port,
	)
	SqlSession, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err, nil // 返回数据库连接错误，但不中断程序
	}
	sqlDB, err := SqlSession.DB()
	if err != nil {
		return err, nil // 返回数据库连接错误，但不中断程序
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxPoolConn)
	return nil, SqlSession
}
