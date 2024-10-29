package helper

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

type GormDb struct {
	DataSource string
	Db         *gorm.DB
}

func RegisterGormMysql() GormDb {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, //慢sql阈值
		LogLevel:      logger.Info, //log等级
		Colorful:      true,        //禁止彩色打印
	})
	mysqlUrl := os.Getenv("MYSQL_USERNAME") + ":" +
		os.Getenv("MYSQL_PASSWORD") + "@tcp(" +
		os.Getenv("MYSQL_HOST") + ":" +
		os.Getenv("MYSQL_PORT") + ")/" +
		os.Getenv("MYSQL_DATABASE") +
		"?charset=utf8mb4&parseTime=true&collation=utf8mb4_0900_ai_ci&loc=Local"
	db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:          newLogger,
		CreateBatchSize: 2000,
	})

	if err != nil {
		panic("Gorm failed to connect database,err:" + err.Error())
	}

	return GormDb{DataSource: mysqlUrl, Db: db}
}

func InitGormDbByDataSourceName(dbName string) GormDb {
	noNameConnStr := os.Getenv("MYSQL_NO_DB_Name")
	connStr := fmt.Sprintf(noNameConnStr, dbName)
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second, //慢sql阈值
		LogLevel:      logger.Info, //log等级
		Colorful:      true,        //禁止彩色打印
	})

	db, err := gorm.Open(mysql.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger:          newLogger,
		CreateBatchSize: 2000,
	})

	if err != nil {
		panic("Gorm failed to connect database,err:" + err.Error())
	}

	return GormDb{DataSource: connStr, Db: db}
}
