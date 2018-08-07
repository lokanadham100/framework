package database

import (
	"context"
	"github.com/jinzhu/gorm"
)

var eventDispatcher, _ = event.GetWrapEvent("database")

func gormCreateStarted(scope *gorm.Scope) {	
	defaultGormStarted(scope)	
}

func gormCreateEnded(scope *gorm.Scope) {
	defaultGormEnded(scope,"Create")
}

func gormUpdateStarted(scope *gorm.Scope){
	defaultGormStarted(scope)
}

func gormUpdateEnded(scope *gorm.Scope) {
	defaultGormEnded(scope,"Update")
}

func gormQueryStarted(scope *gorm.Scope) {	
	defaultGormStarted(scope)
}

func gormQueryEnded(scope *gorm.Scope) {
	defaultGormEnded(scope,"Select")
}

func gormDeleteStarted(scope *gorm.Scope) {	
	defaultGormStarted(scope)
}

func gormDeleteEnded(scope *gorm.Scope) {
	defaultGormEnded(scope,"Delete")
}

func gormRowQueryStarted(scope *gorm.Scope) {	
	defaultGormStarted(scope)
}

func gormRowQueryEnded(scope *gorm.Scope) {
	defaultGormEnded(scope,"RowQuery")
}

func defaultGormStarted(scope *gorm.Scope){
	if ctx, ok := scope.DB().Get("context"); ok == true {	
		ctx := ctx.(*context.Context)		
		dbEvent, _ := eventDispatcher.Start(ctx)
		scope.Set("event",dbEvent)
	}
}

func defaultGormEnded(scope *gorm.Scope, qtype string){
	if dbEvent, ok := scope.Get("event"); ok == true {
		dbEvent := dbEvent.(*event.WrapInterface)
		dbEvent.stop(
			{
				'qtype' : qtype,
				'query' : scope.SQL,
			},
		)
	}
}