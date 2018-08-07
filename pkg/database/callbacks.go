package database

import (
	"github.com/jinzhu/gorm"
)

func init() {
	registerCreateCallbacks()
	registerUpdateCallbacks()
	registerQueryCallbacks()
	registerDeleteCallbacks()
	registerRowQueryCallbacks()	
}

func registerCreateCallbacks(){
	gorm.DefaultCallback.Query().Before("gorm:create").Register("trace:create_start", gormCreateStarted)
	gorm.DefaultCallback.Query().After("gorm:create").Register("trace:create_end", gormCreateEnded)	
}

func registerUpdateCallbacks(){
	gorm.DefaultCallback.Query().Before("gorm:update").Register("trace:update_start", gormUpdateStarted)
	gorm.DefaultCallback.Query().After("gorm:update").Register("trace:update_end", gormUpdateEnded)	
}

func registerQueryCallbacks(){
	gorm.DefaultCallback.Query().Before("gorm:query").Register("trace:query_start", gormQueryStarted)
	gorm.DefaultCallback.Query().After("gorm:query").Register("trace:query_end", gormQueryEnded)	
}

func registerDeleteCallbacks(){
	gorm.DefaultCallback.Query().Before("gorm:delete").Register("trace:delete_start", gormDeleteStarted)
	gorm.DefaultCallback.Query().After("gorm:delete").Register("trace:delete_end", gormDeleteEnded)	
}

func registerRowQueryCallbacks(){
	gorm.DefaultCallback.Query().Before("gorm:row_query").Register("trace:row_query_start", gormRowQueryStarted)
	gorm.DefaultCallback.Query().After("gorm:row_query").Register("trace:row_query_end", gromRowQueryEnded)	
}