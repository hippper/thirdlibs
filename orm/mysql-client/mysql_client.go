package mysqlclient

import "github.com/jinzhu/gorm"

type MysqlClient struct {
	URL         string
	IdleSize    int
	MaxSize     int
	MaxLifeTime int64
	SqlDebug    int
	Memory      bool
}

func MysqlClientInit(configs []MysqlClient) {
	mysqlClientManager.configs = configs

	if len(configs) == 0 {
		Log.Warnf("db config not found...")
		return
	}

	for idx, conf := range configs {

		var db *gorm.DB
		if conf.Memory {
			db = mysqlClientManager.openMemDBConn(conf)
		} else {
			db = mysqlClientManager.openDBConn(conf)
		}

		mysqlClientManager.dbs = append(mysqlClientManager.dbs, db)

		if idx == 0 {
			mysqlClientManager.masterDB = db
			mysqlClientManager.slaveDB = db
		}

		if idx == 1 { // 如果有配置从库
			mysqlClientManager.slaveDB = db
		}

		if mysqlClientManager.memoryDB == nil && conf.Memory { // 只支持一个内存库
			mysqlClientManager.memoryDB = db
		}
	}

	if len(mysqlClientManager.dbs) == 1 {
		mysqlClientManager.masterDB = mysqlClientManager.openDBConn(configs[0])
		mysqlClientManager.slaveDB = mysqlClientManager.masterDB
	}

	if len(configs) == 1 {
		mysqlClientManager.masterDB = mysqlClientManager.openDBConn(configs[0])
		mysqlClientManager.slaveDB = mysqlClientManager.masterDB
	} else {
		mysqlClientManager.masterDB = mysqlClientManager.openDBConn(configs[0])
		mysqlClientManager.slaveDB = mysqlClientManager.openDBConn(configs[1])
	}
}

func MysqlClientRegisterModels(models ...interface{}) {
	mysqlClientManager.models = models
}

func MysqlClientRegisterMemoryModels(models ...interface{}) {
	mysqlClientManager.memoryModels = models
}
