package ddl

import (
	"database/sql"
	"go-ticket/lla"
	"errors"
	"reflect"
)

const (
	ddlTimeReturnError  = "SQL %s works %d mcs; Return error: %v"
	ddlTimeRowsFetched  = "SQL %s works %d mcs; Rows fetched: %d"
	ddlTimeRowsAffected = "SQL %s works %d mcs; Rows affected:%d"
)
const (
	ddl_Select = "Select"
	ddl_Update = "Update"
	ddl_Insert = "Insert"
	ddl_Delete = "Delete"
	ddl_Search = "Search"
	ddl_Rating = "Rating"
)


func DdlPanicRecover(err *error) {
	if r := recover(); r != nil {
		ddlLog.Panic("Panic is recovered: %+v", r)
		if err != nil {
			*err = errors.New("Panic have been recovered")
		}
	}
}

///////////////////////////////////////////////////////////////////////
//
// Common linker to database storage
//
type DBaseLink struct {
	conn *DBaseConn		// Pointer to parent database connection
	Err  error			// Error from database layer
	Done int64			// Rows affected or fetch count
	SqlTime int			// Wait for database reply (microsecond)
}

// Set connection to DAO object
func (link *DBaseLink) SetConnection(conn *DBaseConn) {
	link.conn = conn
	link.Err = nil
	link.Done = 0
	link.SqlTime = 0
}

func (link *DBaseLink) Driver() string {
	if link.conn != nil && link.conn.pool != nil {
		return link.conn.pool.driver
	}
	return ""
}


///////////////////////////////////////////////////////////////////////
// Select one row from database
func (link *DBaseLink) ExecuteSelect(sqlText string, params []interface{}, unit interface{}) error {

	ddlLog.Dump(sqlText)
	link.Err = link.executeSelectSql(sqlText, params, unit)

	if link.Err != nil {
		ddlLog.Error(ddlTimeReturnError, ddl_Select, link.SqlTime, link.Err)
	} else {
		ddlLog.Debug(ddlTimeRowsFetched, ddl_Select, link.SqlTime, link.Done)
	}
	return link.Err
}


///////////////////////////////////////////////////////////////////////
// Search all rows from database
func (link *DBaseLink) ExecuteSearch(sqlText string, params []interface{}, list interface{}) error {

	ddlLog.Dump(sqlText)
	link.Err = link.executeSearchSql(sqlText, params, list)

	if link.Err != nil {
		ddlLog.Error(ddlTimeReturnError, ddl_Search, link.SqlTime, link.Err)
	} else {
		ddlLog.Debug(ddlTimeRowsFetched, ddl_Search, link.SqlTime, link.Done)
	}
	return link.Err
}


///////////////////////////////////////////////////////////////////////
// Select one row from database
func (link *DBaseLink) ExecuteEstimate(sqlText string, params []interface{}, result interface{}) error {

	ddlLog.Dump(sqlText)
	out := make(map[string]reflect.Value)
	out[ddl_estimate_result] = reflect.ValueOf(result)
	link.Err = link.executeReturnSql(sqlText, params, out)

	if link.Err != nil {
		ddlLog.Error(ddlTimeReturnError, ddl_Rating, link.SqlTime, link.Err)
	} else {
		ddlLog.Debug(ddlTimeRowsFetched, ddl_Rating, link.SqlTime, link.Done)
	}
	return link.Err
}


///////////////////////////////////////////////////////////////////////
// Update query execution
func (link *DBaseLink) ExecuteUpdate(sqlText string, params []interface{}) error {
	ddlLog.Dump(sqlText)
	link.Err = link.executeCommandSql(sqlText, params)

	if link.Err != nil {
		ddlLog.Error(ddlTimeReturnError, ddl_Update, link.SqlTime, link.Err)
	} else {
		ddlLog.Debug(ddlTimeRowsAffected, ddl_Update, link.SqlTime, link.Done)
	}
	return link.Err
}


///////////////////////////////////////////////////////////////////////
// Insert query execution
func (link *DBaseLink) ExecuteInsert(sqlText string, params []interface{}) error {

	ddlLog.Dump(sqlText)
	link.Err = link.executeCommandSql(sqlText, params)

	if link.Err != nil {
		ddlLog.Error(ddlTimeReturnError, ddl_Insert, link.SqlTime, link.Err)
	} else {
		ddlLog.Debug(ddlTimeRowsAffected, ddl_Insert, link.SqlTime, link.Done)
	}
	return link.Err
}

func (link *DBaseLink) ExecuteAutoInsert(sqlText string, params []interface{},
	out map[string]reflect.Value, ret bool) error {

	ddlLog.Dump(sqlText)
	if ret{
		link.Err = link.executeReturnSql(sqlText, params, out)
	} else {
		link.Err = link.executeCommandSql(sqlText, params)
	}
	if link.Err != nil {
		ddlLog.Error(ddlTimeReturnError, ddl_Insert, link.SqlTime, link.Err)
	} else {
		ddlLog.Debug(ddlTimeRowsAffected, ddl_Insert, link.SqlTime, link.Done)
	}
	return link.Err
}


///////////////////////////////////////////////////////////////////////
// Delete query execution
func (link *DBaseLink) ExecuteDelete(sqlText string, params []interface{}) error {
	ddlLog.Dump(sqlText)
	link.Err = link.executeCommandSql(sqlText, params)

	if link.Err != nil {
		ddlLog.Error(ddlTimeReturnError, ddl_Delete, link.SqlTime, link.Err)
	} else {
		ddlLog.Debug(ddlTimeRowsAffected, ddl_Delete, link.SqlTime, link.Done)
	}
	return link.Err
}


///////////////////////////////////////////////////////////////////////
//
// Execute select SQL
func (link *DBaseLink) executeSelectSql(sqlText string, params []interface{}, unit interface{}) (err error) {
	defer DdlPanicRecover(&err)
	var timer lla.DurationTimer
	timer.StartTimer()

	var row *sql.Rows
	row, err = link.conn.doQuery(sqlText, params...)
	link.SqlTime = timer.Microseconds()
	if err == nil {
		link.Done, err = fetchSelectedRow(row, unit)
	}
	ddlLog.Debug("SQL SelectOne return %s", lla.GetErrorText(err))
	return err
}

// Execute search SQL
func (link *DBaseLink) executeSearchSql(sqlText string, params []interface{}, list interface{}) (err error) {
	defer DdlPanicRecover(&err)
	var timer lla.DurationTimer
	timer.StartTimer()

	var row *sql.Rows
	row, err = link.conn.doQuery(sqlText, params...)
	link.SqlTime = timer.Microseconds()
	if err == nil {
		link.Done, err = fetchSearchedRows(row, list)
	}
	ddlLog.Debug("SQL SelectAll return %s", lla.GetErrorText(err))
	return err
}

// Execute command SQL
func (link *DBaseLink) executeCommandSql(sqlText string, params []interface{}) (err error) {
	defer DdlPanicRecover(&err)
	var timer lla.DurationTimer
	timer.StartTimer()

	var res sql.Result
	res, err = link.conn.doExec(sqlText, params...)
	link.SqlTime = timer.Microseconds()
	if err == nil {
		link.Done, err = res.RowsAffected()
	}
	ddlLog.Debug("SQL ExecQuery return %s", lla.GetErrorText(err))
	return err
}

// Execute Return SQL
func (link *DBaseLink) executeReturnSql(sqlText string, params []interface{}, m map[string]reflect.Value) (err error) {
	defer DdlPanicRecover(&err)
	var timer lla.DurationTimer
	timer.StartTimer()

	var row *sql.Rows
	row, err = link.conn.doQuery(sqlText, params...)
	link.SqlTime = timer.Microseconds()
	if err == nil {
		link.Done, err = fetchReturnedRow(row, m)
	}
	ddlLog.Debug("SQL ReturnRow return %s", lla.GetErrorText(err))
	return err
}




