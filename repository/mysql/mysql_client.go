package mysql

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	"github.com/qisst/ms-nadra-verification/conf"
	"github.com/qisst/ms-nadra-verification/repository"
	"github.com/siddontang/go/log"
)

var store repository.Store
var storeOnce sync.Once

type Store struct {
	db *gorm.DB
}

// SharedStore return global or single instance of firebase connection (bounded in sync once)
func SharedStore() repository.Store {
	storeOnce.Do(func() {
		conn, err := initDb()
		if err != nil {
			panic(err)
		}
		store = NewStore(conn)
	})

	return store
}

// NewStore create store object
func NewStore(conn *gorm.DB) *Store {
	return &Store{
		db: conn,
	}
}

// initDb initiate firebase connections
func initDb() (*gorm.DB, error) {

	log.Info("Grabbing variables from config file.")

	dbConfig := conf.GetConfig().Mysql

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.DbUserName,
		dbConfig.DbPassword,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
	)

	log.Info("Connecting MYSQL Database.")

	connection, err := gorm.Open("mysql", dbUri)
	if err != nil {
		log.Errorf("ERROR: Couldn't establish database connection: %v", err.Error())
		return connection, err
	}
	//defer connection.Close()

	log.Info("MYSQL Database Connected Successfully.")

	return connection, nil
}
