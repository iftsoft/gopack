package ddl

import (
	"fmt"
	"strings"
)

///////////////////////////////////////////////////////////////////////
//
// Common interface for dialect translator
//
type Dialector interface {
	QuoteIdent(id string) string
	ParamHolder(n int) string
	IsReturnKey() bool
}

func GetDialector(driver string) Dialector {
	switch driver {
	case "postgres":
		return dialPostgre{}
	case "mysql":
		return dialMysql{}
	case "sqlite":
		return dialSqlite3{}
	}
	return dialDefault{}
}

///////////////////////////////////////////////////////////////////////
//
// Default dialect translator
//

type dialDefault struct{}

func (d dialDefault) QuoteIdent(s string) string {
	return quoteIdent(s, `"`)
}

func (d dialDefault) ParamHolder(n int) string {
	return "?"
}

func (d dialDefault) IsReturnKey() bool {
	return false
}


///////////////////////////////////////////////////////////////////////
//
// PostgreSQL dialect translator
//

type dialPostgre struct{}

func (d dialPostgre) QuoteIdent(s string) string {
	return quoteIdent(s, `"`)
}

func (d dialPostgre) ParamHolder(n int) string {
	return fmt.Sprintf("$%d", n+1)
}

func (d dialPostgre) IsReturnKey() bool {
	return true
}


///////////////////////////////////////////////////////////////////////
//
// MySQL dialect translator
//

type dialMysql struct{}

func (d dialMysql) QuoteIdent(s string) string {
	return quoteIdent(s, "`")
}

func (d dialMysql) ParamHolder(_ int) string {
	return "?"
}

func (d dialMysql) IsReturnKey() bool {
	return false
}


///////////////////////////////////////////////////////////////////////
//
// SQLite dialect translator
//

type dialSqlite3 struct{}

func (d dialSqlite3) QuoteIdent(s string) string {
	return quoteIdent(s, `"`)
}

func (d dialSqlite3) ParamHolder(_ int) string {
	return "?"
}

func (d dialSqlite3) IsReturnKey() bool {
	return false
}


///////////////////////////////////////////////////////////////////////
//  Helper functions
//
func quoteIdent(s, quote string) string {
	part := strings.SplitN(s, ".", 2)
	if len(part) == 2 {
		return quote + part[0] + quote + "." + quote + part[1] + quote
	}
	return quote + s + quote
}
