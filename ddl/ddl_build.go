package ddl

import (
	"reflect"
	"fmt"
	"github.com/iftsoft/gopack/lla"
)

const ddl_estimate_result = "estimate_result"

///////////////////////////////////////////////////////////////////////
//
// Common SQL Builder
//
type CommonBuilder struct {
	parSlice  []interface{}
	dialect   Dialector
	TableName string
	MainAlias string
	SqlText   string
	paramStr  string
	paramCnt  int
}

func (this *CommonBuilder)ParSlice() []interface{} {
	return this.parSlice
}

func (this *CommonBuilder)Dialect() Dialector {
	return this.dialect
}

func (this *CommonBuilder)And() (*CommonBuilder) {
	this.paramStr += " AND "
	return this
}
func (this *CommonBuilder)Or() (*CommonBuilder) {
	this.paramStr += " OR "
	return this
}
func (this *CommonBuilder)GrpBeg() (*CommonBuilder) {
	this.paramStr += "( "
	return this
}
func (this *CommonBuilder)GrpEnd() (*CommonBuilder) {
	this.paramStr += " )"
	return this
}

func (this *CommonBuilder)IsNull(alias string, name string) (*CommonBuilder) {
	if alias != "" {
		alias += "."
	}
	this.paramStr += fmt.Sprintf("%s%s IS NULL", alias, name)
	return this
}
func (this *CommonBuilder)IsNotNull(alias string, name string) (*CommonBuilder) {
	if alias != "" {
		alias += "."
	}
	this.paramStr += fmt.Sprintf("%s%s IS NOT NULL", alias, name)
	return this
}

func (this *CommonBuilder)Equal(alias string, name string, value interface{}) (*CommonBuilder) {
	this.appendParam(alias, name, " = ", value)
	return this
}
func (this *CommonBuilder)NotEqual(alias string, name string, value interface{}) (*CommonBuilder) {
	this.appendParam(alias, name, " != ", value)
	return this
}
func (this *CommonBuilder)Grater(alias string, name string, value interface{}) (*CommonBuilder) {
	this.appendParam(alias, name, " > ", value)
	return this
}
func (this *CommonBuilder)GrEqual(alias string, name string, value interface{}) (*CommonBuilder) {
	this.appendParam(alias, name, " >= ", value)
	return this
}
func (this *CommonBuilder)Later(alias string, name string, value interface{}) (*CommonBuilder) {
	this.appendParam(alias, name, " < ", value)
	return this
}
func (this *CommonBuilder)LtEqual(alias string, name string, value interface{}) (*CommonBuilder) {
	this.appendParam(alias, name, " <= ", value)
	return this
}

func (this *CommonBuilder)appendParam(alias string, name string, cmd string, value interface{}) (*CommonBuilder) {
	if alias != "" {
		this.paramStr += alias + "."
	}
	this.paramStr += name
	this.paramStr += cmd
	this.paramStr += this.dialect.ParamHolder(this.paramCnt)
	this.paramCnt++
	this.AppendParamValue(value)
	return this
}

func (this *CommonBuilder)AppendParamValue(value interface{}) {
	this.parSlice = append(this.parSlice, CheckValueToJson(value))
}

func (this *CommonBuilder)SetParamList(keys lla.ParamList) (*CommonBuilder) {
	var count = 0
	for _, par := range keys {
		if col, ok := GetTableColumnName(this.TableName, par.Field); ok {
			if this.paramCnt > 0 {	this.And()	}
			this.Equal("", col, par.Value)
			count++
		}
	}
	return this
}

func (this *CommonBuilder)getWhereClause() string {
	if this.paramStr != "" {
		return " WHERE " + this.paramStr
	}
	return ""
}



///////////////////////////////////////////////////////////////////////
//
// Select SQL Builder
//
type SelectBuilder struct {
	CommonBuilder
	fieldCnt int
	fieldStr string
	JoinList string
	OrderBy  string
	Limit	uint32
	Offset	uint32
}

func CreateSelectBuilder(driver, table, alias string) *SelectBuilder {
	bld := new(SelectBuilder)
	if bld == nil {
		ddlLog.Error(strErrMakeBuilder)
		return nil
	}
	// Reserve slice for params
	bld.parSlice  = make([]interface{}, 0, 8)
	bld.dialect   = GetDialector(driver)
	bld.TableName = table
	bld.MainAlias = alias
	bld.Limit	= 0
	bld.Offset	= 0
	return bld
}

func (this *SelectBuilder)SetLimitOffset(lim, off uint32) (*SelectBuilder) {
	this.Limit	= lim
	this.Offset	= off
	return this
}

func (this *SelectBuilder) SetupEstimateCount(prkey string) {
	this.fieldStr = fmt.Sprintf( " count(%s.%s) as %s ",
		this.MainAlias, prkey, ddl_estimate_result)
	this.fieldCnt++
}

// Create column list for select
func (this *SelectBuilder) SetupSelectColumns(unit interface{}) {
	this.iterateSelectFields(reflect.ValueOf(unit), this.MainAlias, "", false)
}

// Recursive iterate through object structure and fill column map with value holders
func (this *SelectBuilder) iterateSelectFields(value reflect.Value, tab string, pref string, as bool) {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		if !value.IsNil() {
			this.iterateSelectFields(value.Elem(), tab, pref, as)
		}
	case reflect.Struct:
		t := value.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" && !field.Anonymous {
				continue
			}
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				this.iterateSelectFields(value.Field(i), tab, pref, as)
				continue
			}
			info := FieldInfo{}
			info.InitFieldInfo(field)
			if info.DdlType == Field_Skip {
				// Skip ignored field
				continue
			}
			if (info.DdlType & Field_Inc) == Field_Inc {
				this.iterateSelectFields(value.Field(i), tab, info.ColName, as)
				continue
			}
			if (info.DdlType & Field_Ref) == Field_Ref {
				this.iterateSelectFields(value.Field(i), info.ColName, "", true)
				continue
			}
			// Append Object field to column list
			if this.fieldCnt > 0 {
				this.fieldStr += ", "
			}
			this.fieldStr += GetColumnName(info.ColName, pref, tab, as)
			this.fieldCnt++
		}
	}
}


func (this *SelectBuilder) defineUnitFields(value reflect.Value, tab string, pref string, as bool) {
	var info FieldInfo

	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		if !value.IsNil() {
			this.defineUnitFields(value.Elem(), tab, pref, as)
		}
	case reflect.Struct:
		t := value.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" && !field.Anonymous {
				continue
			}
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				this.defineUnitFields(value.Field(i), tab, pref, as)
				continue
			}
			info.InitFieldInfo(field)
			if (info.DdlType & Field_Inc) == Field_Inc {
				this.defineUnitFields(value.Field(i), tab, info.ColName, as)
				continue
			}
			if (info.DdlType & Field_Ref) == Field_Ref {
				this.defineUnitFields(value.Field(i), info.ColName, "", true)
				continue
			}
			this.AppendField(info.ColName, pref, tab, as)
		}
	}
}

func (this *SelectBuilder) AppendField(col string, pref string, tab string, as bool) {
	if this.fieldCnt > 0 {
		this.fieldStr += ", "
	}
	this.fieldStr += GetColumnName(col, pref, tab, as)
	this.fieldCnt++
}

func (this *SelectBuilder) JoinTable(table string, alias string, col string, key string) {
	this.JoinList += fmt.Sprintf(" LEFT JOIN %s %s ON %s.%s = %s.%s ",
		table, alias, this.MainAlias, col, alias, key)
}
func (this *SelectBuilder) JoinTable2(table string, alias string, col1 string, key1 string, col2 string, key2 string) {
	this.JoinList += fmt.Sprintf(" LEFT JOIN %s %s ON %s.%s = %s.%s AND %s.%s = %s.%s ",
		table, alias, this.MainAlias, col1, alias, key1, this.MainAlias, col2, alias, key2)
}

func (this *SelectBuilder)AndParamMap(keys map[string]interface{}) (*SelectBuilder) {
	var count = 0
	for key, val := range keys {
		if col, ok := GetTableColumnName(this.TableName, key); ok {
			if this.paramCnt > 0 {	this.And()	}
			this.Equal("", col, val)
			count++
		}
	}
	return this
}

func (this *SelectBuilder)SetListFilter(filter *lla.ClauseList) (*SelectBuilder) {
	this.iterateCondGroup(filter, lla.Type_Group_AND)
	return this
}

func (this *SelectBuilder) iterateCondGroup(val interface{}, grpType lla.ClauseType) {
	if val == nil { return }
	ref := reflect.ValueOf(val)
	switch ref.Kind() {
	case reflect.Ptr:
		if !ref.IsNil() {
			this.iterateCondGroup(ref.Elem().Interface(), grpType)
		}
	case reflect.Slice:
		if list, ok := val.(lla.ClauseList); ok {
			for i, cond := range list {
				if i > 0 {
					switch grpType {
					case lla.Type_Group_AND:	this.And()
					case lla.Type_Group_OR: 	this.Or()
					}
				}
				switch cond.Type {
				case lla.Type_Statement:
					if col, ok := GetTableColumnName(this.TableName, cond.Field); ok {
						switch cond.Stat {
						case lla.Stat_Equal :
							this.Equal("", col, cond.Value)
						case lla.Stat_NotEqual :
							this.NotEqual("", col, cond.Value)
						case lla.Stat_Grater :
							this.Grater("", col, cond.Value)
						case lla.Stat_GrEqual :
							this.GrEqual("", col, cond.Value)
						case lla.Stat_Later :
							this.Later("", col, cond.Value)
						case lla.Stat_LtEqual :
							this.LtEqual("", col, cond.Value)
						case lla.Stat_IsNull :
							this.IsNull("", col)
						case lla.Stat_IsNotNull :
							this.IsNotNull("", col)
						}
					}
				case lla.Type_Group_AND, lla.Type_Group_OR:
					this.GrpBeg()
					this.iterateCondGroup(cond.Value, cond.Type)
					this.GrpEnd()
				}
			}
		}
	}
	return
}



func (this *SelectBuilder)SetSortOrder(sortList *lla.OrderList) (*SelectBuilder) {
	var count = 0
	if sortList == nil {
		return this
	}
	for _, sort := range *sortList {
		if col, ok := GetTableColumnName(this.TableName, sort.Field); ok {
			if count > 0 {
				this.OrderBy += ", "
			}
			this.OrderBy += col
			switch sort.Sort {
			case lla.Sort_Asc :
				this.OrderBy += " asc"
			case lla.Sort_Desc :
				this.OrderBy += " desc"
			}
			count++
		}
	}
	return this
}


func (this *SelectBuilder)getOrderClause() string {
	if this.OrderBy != "" {
		return " ORDER BY " + this.OrderBy
	}
	return ""
}

func (this *SelectBuilder)getLimitClause() string {
	str := ""
	if this.Limit > 0 {
		str = fmt.Sprintf(" LIMIT %d", this.Limit)
		if this.Offset > 0 {
			str += fmt.Sprintf(" OFFSET %d", this.Offset)
		}
	}
	return str
}

// Generate SELECT query
func (this *SelectBuilder) BuildQuery() string {
	this.SqlText = fmt.Sprintf("SELECT %s FROM %s %s%s%s%s%s;",
		this.fieldStr, this.TableName, this.MainAlias, this.JoinList,
		this.getWhereClause(), this.getOrderClause(), this.getLimitClause())
	return this.SqlText
}



///////////////////////////////////////////////////////////////////////
//
// Update SQL Builder
//
type UpdateBuilder struct {
	CommonBuilder
	setText  	string
}

func CreateUpdateBuilder(driver, table string) *UpdateBuilder {
	bld := new(UpdateBuilder)
	if bld == nil {
		ddlLog.Error(strErrMakeBuilder)
		return nil
	}
	// Reserve slice for params
	bld.parSlice  = make([]interface{}, 0, 32)
	bld.dialect   = GetDialector(driver)
	bld.TableName = table
	return bld
}


// Create set param list for update
func (this *UpdateBuilder) SetupUpdateParams(unit interface{}) {
	this.defineUnitParams(reflect.ValueOf(unit), "")
}

func (this *UpdateBuilder) defineUnitParams(value reflect.Value, pref string) {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		if !value.IsNil() {
			this.defineUnitParams(value.Elem(), pref)
		}
	case reflect.Struct:
		t := value.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" && !field.Anonymous {
				continue
			}
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				this.defineUnitParams(value.Field(i), pref)
				continue
			}
			info := FieldInfo{}
			info.InitFieldInfo(field)
			if info.DdlType == Field_Skip {
				continue
			}
			if (info.DdlType & Skip_Update) != 0 {
				continue
			}
			if (info.DdlType & Field_Inc) == Field_Inc {
				this.defineUnitParams(value.Field(i), info.ColName)
				continue
			}
			this.AppendField(value.Field(i), info.ColName, pref, "")
		}
	}
}

func (this *UpdateBuilder) AppendField(value reflect.Value, col string, pref string, tab string) {
	if this.paramCnt > 0 {
		this.setText += ", "
	}
	this.parSlice = append(this.parSlice, value.Addr().Interface())
	this.setText += getFullColumnName(col, pref, "")
	this.setText += " = "
	this.setText += this.dialect.ParamHolder(this.paramCnt)
	this.paramCnt++
}

// Create set param list for modify
func (this *UpdateBuilder) SetupModifyParams(vals lla.ParamList) {
	for _, par := range vals {
		if col, ok := GetTableColumnName(this.TableName, par.Field); ok {
			if this.paramCnt > 0 {
				this.setText += ", "
			}
			this.parSlice = append(this.parSlice, par.Value)
			this.setText += col
			this.setText += " = "
			this.setText += this.dialect.ParamHolder(this.paramCnt)
			this.paramCnt++
		}
	}
}

// Generate UPDATE query
func (this *UpdateBuilder) BuildQuery() string {
	this.SqlText = fmt.Sprintf("UPDATE %s SET %s WHERE %s;",
		this.TableName, this.setText, this.paramStr)
	return this.SqlText
}


///////////////////////////////////////////////////////////////////////
//
// Insert SQL Builder
//
type InsertBuilder struct {
	CommonBuilder
	autoKeys map[string]reflect.Value
	fieldStr string
}

func CreateInsertBuilder(driver, table string) *InsertBuilder {
	bld := new(InsertBuilder)
	if bld == nil {
		ddlLog.Error(strErrMakeBuilder)
		return nil
	}
	// Reserve slice for params
	bld.parSlice  = make([]interface{}, 0, 32)
	bld.autoKeys  = make(map[string]reflect.Value)
	bld.dialect   = GetDialector(driver)
	bld.TableName = table
	return bld
}

func (this *InsertBuilder) IsAutoInsert() bool {
	return len(this.autoKeys) > 0
}

func (this *InsertBuilder) AutoKeys() map[string]reflect.Value {
	return this.autoKeys
}


// Create set param list for insert
func (this *InsertBuilder) SetupInsertParams(unit interface{}) {
	this.defineUnitParams(reflect.ValueOf(unit), "")
}

func (this *InsertBuilder)defineUnitParams(value reflect.Value, pref string) {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		if !value.IsNil() {
			this.defineUnitParams(value.Elem(), pref)
		}
	case reflect.Struct:
		t := value.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" && !field.Anonymous {
				continue
			}
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				this.defineUnitParams(value.Field(i), pref)
				continue
			}
			info := FieldInfo{}
			info.InitFieldInfo(field)
			if info.DdlType == Field_Skip {
				continue
			}
			if (info.DdlType & Skip_Return) == Skip_Return {
				this.autoKeys[info.ColName] = value.Field(i)
			}
			if (info.DdlType & Skip_Insert) != 0 {
				continue
			}
			if (info.DdlType & Field_Inc) == Field_Inc {
				this.defineUnitParams(value.Field(i), info.ColName)
				continue
			}
			this.AppendField(value.Field(i), info.ColName, pref)
		}
	}
}

// Append inserted field to list
func (this *InsertBuilder) AppendField(value reflect.Value, col string, pref string) {
	if this.paramCnt > 0 {
		this.fieldStr += ", "
		this.paramStr += ", "
	}
	this.fieldStr += getFullColumnName(col, pref, "")
	this.paramStr += this.dialect.ParamHolder(this.paramCnt)
	this.paramCnt++
	this.parSlice = append(this.parSlice, value.Addr().Interface())
}

func (this *InsertBuilder) genReturningText() string {
	var str string
	if this.dialect.IsReturnKey() && len(this.autoKeys) > 0 {
		for key, _ := range this.autoKeys {
			if str != "" {
				str += ", "
			}
			str += key
		}
		return fmt.Sprintf(" RETURNING %s", str)
	}
	return str
}

// Generate INSERT query
func (this *InsertBuilder) BuildQuery() string {
	this.SqlText = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)%s;",
		this.TableName, this.fieldStr, this.paramStr, this.genReturningText())
	return this.SqlText
}



///////////////////////////////////////////////////////////////////////
//
// Delete SQL Builder
//
type DeleteBuilder struct {
	CommonBuilder
}

func CreateDeleteBuilder(driver, table string, alias string) *DeleteBuilder {
	bld := new(DeleteBuilder)
	if bld == nil {
		ddlLog.Error(strErrMakeBuilder)
		return nil
	}
	// Reserve slice for params
	bld.parSlice  = make([]interface{}, 0, 4)
	bld.dialect   = GetDialector(driver)
	bld.TableName = table
	bld.MainAlias = alias
	return bld
}

// Generate DELETE query
func (this *DeleteBuilder) BuildQuery() string {
	this.SqlText = fmt.Sprintf("DELETE FROM %s %s WHERE %s;",
		this.TableName, this.MainAlias, this.paramStr)
	return this.SqlText
}



