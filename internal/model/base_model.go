package model

import (
	"fmt"
	"ginchat/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func init() {
	Db = Connect()
}

func Connect() *gorm.DB {
	c := config.NewConfig()
	viper := c.Viper
	username := viper.Get("database.username") // 账号
	password := viper.Get("database.password") // 密码
	host := viper.Get("database.host")         // 地址
	port := viper.Get("database.port")         // 端口
	database := viper.Get("database.database") // 数据库名称

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, database)
	logfile, _ := os.Create("runtime/logs/sql/sql.log")
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		Logger: logger.New(log.New(logfile, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			}),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Println(err)
		panic("failed to connect mysql.")
	}
	migrateErr := db.AutoMigrate(&Admin{})
	if migrateErr != nil {
		panic("failed to connect migrate.")
	}

	return db
}
