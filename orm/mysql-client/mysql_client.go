package mysqlclient

import (
	"reflect"
	"strings"
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
	SqlDebug     int
	InnoModels   []interface{}
	MyisamModels []interface{}
	MemoryModels []interface{}
}

func NewMysqlClient(c MysqlClientConfig) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.URL), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: &SqlLogger{
			SqlDebug: c.SqlDebug,
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
	migrate(db, "Memory", c.MemoryModels)
	migrate(db, "MyISAM", c.MyisamModels)
	migrate(db, "InnoDb", c.InnoModels)

	return db, nil
}

func migrate(db *gorm.DB, engine string, models []interface{}) {
	if len(models) <= 0 {
		return
	}

	// 创建表时添加后缀
	setTableOption := ""
	if strings.EqualFold(engine, "InnoDb") {
		setTableOption = "ENGINE=InnoDB CHARSET=utf8mb4"
	} else if strings.EqualFold(engine, "Memory") {
		setTableOption = "ENGINE=MEMORY CHARSET=utf8mb4"
	} else if strings.EqualFold(engine, "MyISAM") {
		setTableOption = "ENGINE=MyISAM CHARSET=utf8mb4"
	}
	if setTableOption != "" {
		db.Set("gorm:table_options", setTableOption)
	}

	for _, m := range models {
		if !db.Migrator().HasTable(m) {
			err := db.Migrator().CreateTable(m)
			if err != nil {
				Log.Errorf("m = %v, err = %v", reflect.TypeOf(m), err)
				continue
			}
		}
	}
	err := db.AutoMigrate(models...)
	if err != nil {
		Log.Error(err)
	}
}
