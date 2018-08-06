package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB is the database connector for all the models
	DB  *gorm.DB
	err error
)

func init() {
	gorm.DefaultCallback.Query().Before("gorm:query").Register("gorm:query_start", )
	gorm.DefaultCallback.Query().After("gorm:after_query").Register("gorm:query_end", )	
}
func DB() *gorm.DB{

}

func DBAPM(ctx *context.Context) *gorm.DB{	
	return DB.Set("context",ctx)
}

func query_start(scope *gorm.Scope) {	
	tx, ok := scope.DB().Get("context")
	if ok {	
	txn := tx.(*context.Context)
	txnn,ok := txn.Get("")
	var txnnn newrelic.Transaction
	if ok{
		txnnn = txnn.(newrelic.Transaction)
	}
	ds := &newrelic.DatastoreSegment{
		StartTime:  newrelic.StartSegmentNow(txnnn),
		Product:    newrelic.DatastoreMySQL,
		Collection: "my_table",

	}
	scope.Set("newrelic",ds)
	}
}

func query_end(scope *gorm.Scope) {
	n,ok := scope.Get("newrelic")
	if ok{
	nn := n.(*newrelic.DatastoreSegment)
	nn.Operation = "Select"
	nn.ParameterizedQuery = scope.SQL
	nn.End()
}
}

type database struct {
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

func (db *database)toString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.Name)
}

func ConnectDatabase(config string) {
	var db database
	if _, err := toml.DecodeFile(config, &db); err != nil {
		panic(err)
	} else {
		dbInfo := 
		DB, err = connect(dbInfo)
		if err != nil {
			panic(err)
		}
	}
	DB.LogMode(true)
}

func connect(dbInfo string) (*gorm.DB, error) {
	DB, err = gorm.Open("mysql", dbInfo)
	if err != nil {
		return nil, err
	}
	return DB, err
}

func getDb(config string) *gorm.DB {
	if DB == nil {
		ConnectDatabase(config)
	}
	return DB
}
