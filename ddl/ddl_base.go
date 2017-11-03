package ddl

import (
	"database/sql"
	"fmt"
	"time"
	"errors"
	"github.com/iftsoft/gopack/lla"
)

// Database Configuration
type DBaseConfig struct {
	DbDriver string
	HostName string
	HostPort string
	BaseName string
	UserName string
	UserPass string
	MaxIdle int
	MaxOpen int
	MaxTime int
}


// Print config data to console
func (cfg *DBaseConfig) PrintData() {
	fmt.Println("DbDriver ", cfg.DbDriver)
	fmt.Println("HostName ", cfg.HostName)
	fmt.Println("HostPort ", cfg.HostPort)
	fmt.Println("BaseName ", cfg.BaseName)
	fmt.Println("UserName ", cfg.UserName)
	fmt.Println("UserPass ", cfg.UserPass)
	fmt.Println("MaxTime  ", cfg.MaxTime)
	fmt.Println("MaxIdle  ", cfg.MaxIdle)
	fmt.Println("MaxOpen  ", cfg.MaxOpen)
}
// Get formatted string with config data
func (cfg *DBaseConfig) String() string {
	str := fmt.Sprintf("Database config: " +
		"DbDriver = %s, HostName = %s, HostPort = %s, BaseName = %s, " +
		"UserName = %s, UserPass = %s, MaxTime = %d, MaxIdle = %d, MaxOpen = %d.",
		cfg.DbDriver, cfg.HostName, cfg.HostPort, cfg.BaseName,
		cfg.UserName, cfg.UserPass, cfg.MaxTime, cfg.MaxIdle, cfg.MaxOpen)
	return str
}

// Log Agent for sql layer logging
var ddlLog lla.LogAgent

func InitLoggerSQL(level int){
	ddlLog.Init(level, "SQL")
}

///////////////////////////////////////////////////////////////////////
//
// Database storage descriptor
//
type DBaseStore struct {
	base   *sql.DB
	dsn    string
	driver string
}

func GetDsnPostgres(cfg *DBaseConfig) string {
	return fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.HostName, cfg.HostPort, cfg.BaseName, cfg.UserName, cfg.UserPass)
}

// Open Database store from config data
func (store *DBaseStore) Open(cfg *DBaseConfig) (err error){
	store.Close()
	if cfg != nil {
		store.driver = cfg.DbDriver
		store.dsn = GetDsnPostgres(cfg)
		ddlLog.Info("Open database: %s", store.dsn)
		store.base, err = sql.Open(cfg.DbDriver, store.dsn)
		if err == nil {
			store.base.SetMaxIdleConns(cfg.MaxIdle)
			store.base.SetMaxOpenConns(cfg.MaxOpen)
			store.base.SetConnMaxLifetime(time.Duration(cfg.MaxTime) * time.Second)
		} else {
			ddlLog.Error("Return error: %s", err.Error())
		}
	}
	return err
}

// Close Database store
func (store *DBaseStore) Close() (err error) {
	if store.base != nil {
		err = store.base.Close()
	}
	return err
}

// Get underlay sql.DB
func (store *DBaseStore) GetSqlDB() *sql.DB {
	return store.base
}


///////////////////////////////////////////////////////////////////////
//
// Connection to database storage
//
type TransactMaker interface {
	Begin() error
	Commit()
	Rollback()
}

type DBaseConn struct {
	pool *DBaseStore
	base *sql.DB
	tran *sql.Tx
}

// Init connection to Database store
func (conn *DBaseConn) SetStore(store *DBaseStore) {
	conn.pool = store
	conn.base = store.base
	conn.tran = nil
}

// Return transaction interface
func (conn *DBaseConn) GetTransactMaker() TransactMaker {
	return conn
}

// Begin database transaction
func (conn *DBaseConn) Begin() (err error) {
	ddlLog.Info("Transaction begin")
	if conn.base != nil {
		conn.tran, err = conn.base.Begin()
	} else {
		err = errors.New("Invalid pointer to DB connect")
	}
	if err == nil {
		ddlLog.Debug("Transaction begin done")
	} else {
		ddlLog.Error("Transaction begin: %s", err.Error())
	}
	return err
}

// Commit database transaction
func (conn *DBaseConn) Commit() {
	ddlLog.Info("Transaction commit")
	if conn.tran != nil {
		err := conn.tran.Commit()
		if err == nil {
			ddlLog.Debug("Transaction commit done")
		} else {
			ddlLog.Error("Transaction commit error: %s", err.Error())
		}
		conn.tran = nil
	} else {
		ddlLog.Warn("Transaction is not started (commit)")
	}
}

// Rollback database transaction
func (conn *DBaseConn) Rollback() {
	ddlLog.Info("Transaction rollback")
	if conn.tran != nil {
		err := conn.tran.Rollback()
		if err == nil {
			ddlLog.Debug("Transaction rollback done")
		} else {
			ddlLog.Error("Transaction rollback error: %s", err.Error())
		}
		conn.tran = nil
	} else {
		ddlLog.Warn("Transaction is not started (rollback)")
	}
}

// Get underlay sql.DB
func (conn *DBaseConn) GetSqlDB() *sql.DB {
	if conn.base != nil {
		return conn.base
	}
	return nil
}

// doQuery executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
func (conn *DBaseConn) doQuery(query string, args ...interface{}) (*sql.Rows, error) {
	if conn.tran != nil {
		return conn.tran.Query(query, args...)
	}
	if conn.base != nil {
		return conn.base.Query(query, args...)
	}
	return nil, ddlErrNullDaoPtr
}

// doExec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
func (conn *DBaseConn) doExec(query string, args ...interface{}) (sql.Result, error) {
	if conn.tran != nil {
		return conn.tran.Exec(query, args...)
	}
	if conn.base != nil {
		return conn.base.Exec(query, args...)
	}
	var res sql.Result
	return res, ddlErrNullDaoPtr
}



