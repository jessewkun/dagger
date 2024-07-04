package mysql

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	dlog "dagger/lib/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

const TAGNAME = "DAGGER_MYSQL"

// connList 数据库连接列表
var connList map[string]*gorm.DB
var mysqlCfg map[string]Config

func init() {
	connList = make(map[string]*gorm.DB)
}

// InitMysql 初始化数据库
func InitMysql(cfg map[string]Config) {
	mysqlCfg = cfg
	for dbName, conf := range cfg {
		conn, err := dbConnect(dbName, conf)
		if err != nil {
			dlog.ErrorWithMsg(context.Background(), TAGNAME, "connect to mysql %s error %s", dbName, err)
			continue
		}
		connList[dbName] = conn
		dlog.Info(context.Background(), TAGNAME, "connect to mysql %s succ", dbName)
	}
}

// dbConnect 连接数据库
func dbConnect(dbName string, conf Config) (*gorm.DB, error) {
	if _, ok := connList[dbName]; ok {
		if connList[dbName] != nil {
			return connList[dbName], nil
		}
	}

	if len(conf.Dsn) < 1 {
		return nil, errors.New(fmt.Sprintf("%s db dsn is empty", dbName))
	}

	logLevel := logger.Silent
	if conf.IsLog {
		logLevel = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logLevel,    // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	var dbOne *gorm.DB
	var err error
	master := conf.Dsn[0]

	slave := conf.Dsn[1:]
	dbOne, err = gorm.Open(mysql.Open(master), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 表前缀
		Logger:         newLogger,
	})

	if err != nil {
		return nil, err
	}

	// 配置读写分离
	dbResolverCfg := dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(master)},
		Replicas: []gorm.Dialector{},
		Policy:   dbresolver.RandomPolicy{},
	}

	// 设置从库
	if len(slave) > 0 {
		var replicas []gorm.Dialector
		for i := 0; i < len(slave); i++ {
			replicas = append(replicas, mysql.Open(slave[i]))
		}
		dbResolverCfg.Replicas = replicas
	}

	if conf.MaxIdleConn == 0 {
		conf.MaxIdleConn = 25
	}

	if conf.MaxConn == 0 {
		conf.MaxConn = 50
	}

	if conf.ConnMaxLife == 0 {
		conf.ConnMaxLife = 3600
	}

	dbOne.Use(
		dbresolver.Register(dbResolverCfg).
			// SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(time.Duration(conf.ConnMaxLife) * time.Second).
			SetMaxIdleConns(conf.MaxIdleConn).
			SetMaxOpenConns(conf.MaxConn),
	)

	return dbOne, nil
}

// GetConn 获取数据库连接
func GetConn(dbIns string) *gorm.DB {
	if len(connList) < 1 {
		return nil
	}
	if _, ok := connList[dbIns]; !ok {
		return nil
	}

	return connList[dbIns]
}

// mysql health check
func HealthCheck() map[string]string {
	resp := make(map[string]string)
	for dbName, conf := range mysqlCfg {
		_, err := dbConnect(dbName, conf)
		if err != nil {
			dlog.ErrorWithMsg(context.Background(), TAGNAME, "connect to mysql %s error %s", dbName, err)
			resp[dbName] = err.Error()
		} else {
			resp[dbName] = "succ"
		}
	}
	return resp
}
