package mysqlclient

import (
	"reflect"
	"time"

	. "github.com/luckyweiwei/base/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MysqlClientConfig struct {
	Name         string // 名称 master slave memory
	URL          string // dsn
	IdleSize     int
	MaxSize      int
	MaxLifeTime  int64
	InnoModels   []interface{}
	MyisamModels []interface{}
	MemoryModels []interface{}
}

func NewMysqlClient(c MysqlClientConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.URL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		Log.Errorf("connect to mysql failed, err = %v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		Log.Error(err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(c.IdleSize)
	sqlDB.SetMaxOpenConns(c.MaxSize)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifeTime) * time.Second)

	// migrate
	memoryModels := c.MemoryModels
	if len(memoryModels) > 0 {
		// 创建表时添加后缀
		db = db.Set("gorm:table_options", "ENGINE=MEMORY CHARSET=utf8mb4")

		for _, m := range memoryModels {
			if !db.Migrator().HasTable(m) {
				err := db.Migrator().CreateTable(m)
				if err != nil {
					Log.Errorf("m = %v, err = %v", reflect.TypeOf(m), err)
					continue
				}
			}
		}
		err := db.AutoMigrate(memoryModels...)
		if err != nil {
			Log.Error(err)
		}
	}

	innoModels := c.InnoModels
	if innoModels != nil {
		if len(innoModels) > 0 {
			db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

		}
	}

	//db = db.Set("gorm:table_options", "ENGINE=MyISAM CHARSET=utf8mb4")
	//
	return nil, nil
}

func migrate(db *gorm.DB, engine string, models []interface{}) {

}

// func MysqlClientRegisterModels(models ...interface{}) {
// 	mysqlClientManager.models = models
// }
