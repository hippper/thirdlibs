package mysqlclient

import (
	"sync"

	. "github.com/luckyweiwei/base/logger"
	"gorm.io/gorm"
)

type MysqlClientManager struct {
	mysqlClientMap *sync.Map
}

var mysqlClientManager *MysqlClientManager = nil

func MysqlClientManagerInit(configs []MysqlClientConfig) error {
	mysqlClientManager = &MysqlClientManager{
		mysqlClientMap: &sync.Map{},
	}

	for _, config := range configs {
		client, err := NewMysqlClient(config)
		if err != nil {
			Log.Error(err)
			return err
		}

		mysqlClientManager.mysqlClientMap.Store(config.Name, client)
	}
	return nil
}

func GetMysqlClientManager() *MysqlClientManager {
	return mysqlClientManager
}

func (r *MysqlClientManager) GetMysqlClient(name string) *gorm.DB {
	clientobj, ok := r.mysqlClientMap.Load(name)
	if !ok {
		Log.Errorf("can't find client in map, name=%v", name)
		return nil
	}

	client := clientobj.(*gorm.DB)

	return client
}
