package database

import (
	"fmt"
	"context"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/voonik/framework/pkg/config"
	"github.com/voonik/framework/pkg/logger"
)

var (
	// DB is the database connector for all the models
	DB  *gorm.DB
	err error
)

func Init(){
	ConnectDatabase()
	ConfigDatabaseConnection()
}

func DBAPM(ctx *context.Context) *gorm.DB{	
	return DB.Set("context",ctx)
}

func toString(db interface{}) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.Name)
}

func ConnectDatabase() {
	dbInfo := toString(config.DatabaseConfig())
	DB, err = connect(dbInfo)
	if err != nil {
		panic(err)
	}		
}

func connect(dbInfo string) (*gorm.DB, error) {
	DB, err = gorm.Open("mysql", dbInfo)
	if err != nil {
		return nil, err
	}
	return DB, err
}

func ConfigDatabaseConnection(){
	DB.LogMode(true)
	DB.SetMaxIdleConns(100)  //TODO : make configurable
	DB.SetMaxOpenConns(100)  //TODO : make configurable
	DB.SetLogger(logger.getLoggerWithName("mysql"))
}