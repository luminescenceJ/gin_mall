package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
	"time"
)

var _db *gorm.DB

func Database(connRead, connWrite string) {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       connRead,
		DefaultStringSize:         256,  // string类型默认长度
		DisableDatetimePrecision:  true, //禁止datetime精度，mysql 5.6前不支持
		DontSupportRenameIndex:    true, //重命名索引，先删除索引再重建，mysql5.7前不支持
		DontSupportRenameColumn:   true, //用change重命名列 mysql8前不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池，空闲
	sqlDB.SetMaxOpenConns(100) // 打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	_ = _db.Use(dbresolver.
		Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(connWrite)},                      //写操作
			Replicas: []gorm.Dialector{mysql.Open(connRead), mysql.Open(connRead)}, //读操作 两个数据库
			Policy:   dbresolver.RandomPolicy{},
		}))
	migration() // 创建表
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}
