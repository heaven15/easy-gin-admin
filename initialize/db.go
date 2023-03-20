package initialize

import (
	"fmt"
	"github.com/develop-kevin/easy-gin-vue-admin/app/v1/model"
	"github.com/develop-kevin/easy-gin-vue-admin/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"os"
)

func InitDB() {
	dbType := global.EGVA_CONFIG.System.DbType
	if dbType == "" {
		panic("system not config db type")
	}
	m := global.EGVA_CONFIG.GormDB
	if m.DataBase == "" {
		panic("gorm db not database")
	}
	var dbConfig gorm.Dialector
	switch dbType {
	case "MsSQL":
		mssqlConfig := sqlserver.Config{
			DSN:               "sqlserver://" + m.UserName + ":" + m.PassWord + "@" + m.Host + ":" + m.Port + "?database=" + m.DataBase + "&encrypt=disable", // DSN data source name
			DefaultStringSize: 191,                                                                                                                           // string 类型字段的默认长度
		}
		dbConfig = sqlserver.New(mssqlConfig)
	case "MySQL":
		mysqlConfig := mysql.Config{
			DSN:                       m.UserName + ":" + m.PassWord + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.DataBase + "?" + m.Config, // DSN data source name
			DefaultStringSize:         191,                                                                                                  // string 类型字段的默认长度
			SkipInitializeWithVersion: false,                                                                                                // 根据版本自动配置
		}
		dbConfig = mysql.New(mysqlConfig)
	case "Oracle":
		oracleConfig := mysql.Config{
			DSN:               "oracle://" + m.UserName + ":" + m.PassWord + "@" + m.Host + ":" + m.Port + "/" + m.DataBase + "?" + m.Config, // DSN data source name
			DefaultStringSize: 191,                                                                                                           // string 类型字段的默认长度
		}
		dbConfig = mysql.New(oracleConfig)
	case "PgSQL":
		pgsqlConfig := postgres.Config{
			DSN:                  "host=" + m.Host + " user=" + m.UserName + " password=" + m.PassWord + " dbname=" + m.DataBase + " port=" + m.Port + " " + m.Config, // DSN data source name
			PreferSimpleProtocol: false,
		}
		dbConfig = postgres.New(pgsqlConfig)
	}
	var err error
	global.EGVA_DB, err = gorm.Open(dbConfig, GormDBNew.Config(m.Prefix, m.Singular))
	if err != nil {
		panic(fmt.Sprintf("gorm:%s connect failed:%s", dbType, err.Error()))
	} else {
		global.EGVA_DB.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := global.EGVA_DB.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleCons)
		sqlDB.SetMaxOpenConns(m.MaxOpenCons)
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	db := global.EGVA_DB
	err := db.AutoMigrate(
		// 系统模块表
		&model.SysUser{},
		&model.SysDepartment{},
		&model.SysRole{},
		&model.SysMenu{},
		&model.SysPermission{},
		&model.SysPost{},
		&model.SysRank{},
		&model.SysConfig{},
		&model.SysDictionary{},
		&model.SysElement{},
		&model.SysFile{},
		&model.SysGroup{},
		&model.SysOperation{},
	)
	if err != nil {
		global.EGVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.EGVA_LOG.Info("register table success")
}
