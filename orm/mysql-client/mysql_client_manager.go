package mysqlclient

import (
	"reflect"
	"time"

	"github.com/jinzhu/gorm"
)

type MysqlClientManager struct {
	configs      []DataSource
	dbs          []*gorm.DB
	masterDB     *gorm.DB
	slaveDB      *gorm.DB
	memoryDB     *gorm.DB
	models       []interface{}
	memoryModels []interface{}
}

var mysqlClientManager = &MysqlClientManager{}

func DataSourceInstance() *MysqlClientManager {
	return mysqlClientManager
}

func (d *MysqlClientManager) openDBConn(ds DataSource) *gorm.DB {
	db, err := gorm.Open("mysql", ds.URL)
	if err != nil {
		Log.Errorf("connect to mysql failed, err = %v", err)
		return nil
	}

	db.DB().SetMaxIdleConns(ds.IdleSize)
	db.DB().SetMaxOpenConns(ds.MaxSize)
	db.DB().SetConnMaxLifetime(time.Duration(ds.MaxLifeTime) * time.Second)

	// 设置字符编码
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	db.SingularTable(true)
	if ds.SqlDebug == 1 {
		db.LogMode(true)
		db.SetLogger(TSQLLogger{})
	}

	for _, m := range d.models {
		if !db.HasTable(m) {
			// Log.Debugf("m = %v", reflect.TypeOf(m))
			err := db.CreateTable(m).Error
			if err != nil {
				Log.Errorf("m = %v, err = %v", reflect.TypeOf(m), err)
			}
		}
	}
	db.AutoMigrate(d.models...)

	return db
}

func (d *MysqlClientManager) openMemDBConn(ds DataSource) *gorm.DB {
	db, err := gorm.Open("mysql", ds.URL)
	if err != nil {
		Log.Errorf("connect to mysql failed, err = %v", err)
		return nil
	}

	db.DB().SetMaxIdleConns(ds.IdleSize)
	db.DB().SetMaxOpenConns(ds.MaxSize)
	db.DB().SetConnMaxLifetime(time.Duration(ds.MaxLifeTime) * time.Second)

	// 设置字符编码
	db = db.Set("gorm:table_options", "ENGINE=MEMORY CHARSET=utf8mb4")
	db.SingularTable(true)
	if ds.SqlDebug == 1 {
		db.LogMode(true)
		db.SetLogger(TSQLLogger{})
	}

	for _, m := range d.memoryModels {
		if !db.HasTable(m) {
			err := db.CreateTable(m).Error
			if err != nil {
				Log.Error(err)
			}
		}
	}
	db.AutoMigrate(d.memoryModels...)

	return db
}

func (d *MysqlClientManager) Master() *gorm.DB {
	return d.masterDB
}

func (d *MysqlClientManager) Slave() *gorm.DB {
	return d.slaveDB
}

func (d *MysqlClientManager) Memory() *gorm.DB {
	return d.memoryDB
}
