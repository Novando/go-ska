package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/novando/go-ska/pkg/logger"
	"time"
)

type Config struct {
	SSL     bool
	Port    uint
	MaxTime uint
	MaxConn int
	Host    string
	Addr    string
	User    string
	Pass    string
	Name    string
}

type MySQL struct {
	db  *sql.DB
	log *logger.Logger
}

func InitMySQL(config Config, l ...*logger.Logger) *MySQL {
	log := logger.Call()
	if len(l) > 0 {
		log = l[0]
	}
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port == 0 {
		config.Port = 3306
	}
	addr := fmt.Sprintf("%v:%v", config.Host, config.Port)
	if config.Addr != "" {
		addr = config.Addr
	}
	if config.User == "" {
		config.User = "root"
	}
	if config.Pass == "" {
		config.Pass = ""
	}
	if config.MaxTime == 0 {
		config.MaxTime = 3
	}
	if config.MaxConn == 0 {
		config.MaxConn = 10
	}
	if config.Name == "" {
		log.Fatalf("mysql Name required")
	}
	dsn := fmt.Sprintf(
		"%v:%v@(%v)/%v?tls=%v",
		config.User,
		config.Pass,
		addr,
		config.Name,
		config.SSL,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Call().Fatalf("%v/%v?sslMode=%v mysql failed: %v", addr, config.Name, config.SSL, err.Error())
		return nil
	}
	if err = db.Ping(); err != nil {
		logger.Call().Fatalf("mysql %v ping failed: %v", config.Name, err.Error())
		return nil
	}
	db.SetConnMaxLifetime(time.Minute * time.Duration(config.MaxTime))
	db.SetMaxOpenConns(config.MaxConn)
	db.SetMaxIdleConns(config.MaxConn)

	logger.Call().Infof("%v/%v?sslMode=%v connected", addr, config.Name, config.SSL)
	return &MySQL{db: db, log: log}
}
