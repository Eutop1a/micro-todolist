package dao

import (
	"context"
	"fmt"
	"todo_list/config"

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
	user := config.DbUser
	database := config.DbName
	password := config.DbPassword
	charset := config.Charset
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, password, host, port, database, charset)
	fmt.Println(dsn)
	err := Database(dsn)
	if err != nil {
		fmt.Println("database init err:", err)
	}

}

func Database(connString string) error {
	var ormLogger logger.Interface = logger.Default.LogMode(logger.Info)
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

	_db = db

	migration()

	return err
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
