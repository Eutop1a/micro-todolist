package dao

import (
	"context"
	"fmt"
	"strings"
	"time"
	"todo_list/config"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"

	"gorm.io/gorm"
)

// 定义对数据库 user model的CRUD操作

var _db *gorm.DB

func InitDB() {
	host := config.DbHost
	port := config.DbPort
	username := config.DbUser
	database := config.DbName
	password := config.DbPassword
	charset := config.Charset
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", user, password, host, port, database, charset)
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", database, "?charset=" + charset + "&parseTime=true"}, "")

	fmt.Println(dsn)
	err := Database(dsn)
	if err != nil {
		fmt.Println("database init err:", err)
	}

}

func Database(connString string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connString, // dsn data source name
		DefaultStringSize:         256,        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,       // 禁用datetime精度，是为了兼容mysql 5.6
		DontSupportRenameIndex:    true,       // 不支持重命名索引
		DontSupportRenameColumn:   true,       // 不支持重命名列
		SkipInitializeWithVersion: false,      // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // user --> users
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	migration()

	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
