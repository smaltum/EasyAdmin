package bootstrap

import (
	"easy-admin/bootstrap/dao"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"time"
)

var DB = new(database)

type database struct {
	GDB *gorm.DB
}

//Init 初始化数据库
func (r *database) _init() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		Cfg.GetVal("database.default.user").String(),
		Cfg.GetVal("database.default.password").String(),
		Cfg.GetVal("database.default.host").String(),
		Cfg.GetVal("database.default.name").String(),
		Cfg.GetVal("database.default.config").String(),
	)
	dCfg := &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   nil,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		DisableNestedTransaction:                 false,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          0,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                           dsn,
		SkipInitializeWithVersion:     false,
		DisableDatetimePrecision:      false,
		DontSupportRenameIndex:        false,
		DontSupportRenameColumn:       false,
		DontSupportForShareClause:     false,
		DontSupportNullAsDefaultValue: false,
	}), dCfg)

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	//.SingularTable(true)
	//sqlDB.LogMode(Cfg.GetVal("database.default.logMode").Bool())
	sqlDB.SetMaxIdleConns(Cfg.GetVal("database.default.maxIdleConn").Int())
	sqlDB.SetMaxOpenConns(Cfg.GetVal("database.default.maxOpenConn").Int())
	sqlDB.SetConnMaxLifetime(time.Hour)
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return Cfg.GetVal("database.default.prefix").String() + defaultTableName
	//}
	if err != nil {
		panic(err)
	}

	r.GDB = db
	// 表迁移
	autoMigrate(r.GDB)

}

// autoMigrateDb 自动迁移
func autoMigrate(db *gorm.DB) {
	// 系统用户表
	if !db.Migrator().HasTable(dao.AdminUsers{}.TableName()) {
		_ = db.AutoMigrate(dao.AdminUsers{})
	}
}
