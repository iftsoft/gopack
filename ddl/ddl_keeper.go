package ddl

import (
	"reflect"
	"go-ticket/lla"
	"errors"
)


type ColumnMap map[string]string

var ddlTableColumns map[string]ColumnMap

func init(){
	ddlTableColumns = make(map[string]ColumnMap)
}

func RegisterObject(unit interface{}, table, alias string) {
	v := reflect.ValueOf(unit)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		panic(strErrNullDataPtr)
	}
	obj := v.Elem()
	if obj.Kind() == reflect.Slice {
		panic(strErrDataIsSlice)
	}
	if 	_, ok := ddlTableColumns[table]; !ok {
		m := createColumnMap(unit, alias)
		ddlTableColumns[table] = m
	}
}

func GetTableColumnMap(table string) (m ColumnMap, ok bool) {
	m, ok = ddlTableColumns[table]
	return m, ok
}

func GetTableColumnName(table, field string) (string, bool) {
	if m, ok := ddlTableColumns[table]; ok {
		col, is := m[field]
		return col, is
	}
	return "", false
}

func createColumnMap(unit interface{}, alias string) ColumnMap {
	// Init and fill map of Value holders
	m := make(ColumnMap)
	iterateColumnMap(reflect.ValueOf(unit), m, alias, "", "")
	return m
}

// Recursive iterate through object structure and fill column map with value holders
func iterateColumnMap(value reflect.Value, m ColumnMap, tab, pref, base string) {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		if !value.IsNil() {
			// Recursive call for pointer to struct
			iterateColumnMap(value.Elem(), m, tab, pref, base)
		}
	case reflect.Struct:
		t := value.Type()
		// Iterate through struct fields
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.PkgPath != "" && !field.Anonymous {
				// Skip unexported field
				continue
			}
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				// Recursive call for anonymous include struct
				iterateColumnMap(value.Field(i), m, tab, pref, base)
				continue
			}
			info := FieldInfo{}
			info.InitFieldInfo(field)
			if info.DdlType == Field_Skip {
				// Skip ignored field
				continue
			}
			key := info.FldName
			if base != "" {
				key = base + "." + info.FldName
			}
			if (info.DdlType & Field_Inc) == Field_Inc {
				// Recursive call for regular include struct
				iterateColumnMap(value.Field(i), m, tab, info.ColName, key)
				continue
			}
			if (info.DdlType & Field_Ref) == Field_Ref {
				// Recursive call for reference to other struct
				iterateColumnMap(value.Field(i), m, info.ColName, "", key)
				continue
			}
			// Add struct field to column map
			col := getFullColumnName(info.ColName, pref, tab)
			if _, ok := m[key]; !ok {
				m[key] = col
			}
		}
	}
}



type DBaseDAO struct {
	DBaseLink
	Name  string
	timer lla.DurationTimer
}

func (dao *DBaseDAO)StartTimer(){
	dao.timer.StartTimer()
}

func (dao *DBaseDAO)StopTimer() int {
	return dao.timer.Microseconds()
}


func (dao *DBaseDAO) MakeSelectQuery(unit interface{}, bld *SelectBuilder) (err error){
	if bld != nil {
		err = dao.ExecuteSelect(bld.BuildQuery(), bld.ParSlice(), unit)
	} else {
		err = errors.New(strErrNullBldPtr)
	}
	return err
}

func (dao *DBaseDAO) MakeSearchQuery(list interface{}, bld *SelectBuilder) (err error){
	if bld != nil {
		err = dao.ExecuteSearch(bld.BuildQuery(), bld.ParSlice(), list)
	} else {
		err = errors.New(strErrNullBldPtr)
	}
	return err
}

func (dao *DBaseDAO) MakeEstimateQuery(result interface{}, bld *SelectBuilder) (err error){
	if bld != nil {
		err = dao.ExecuteEstimate(bld.BuildQuery(), bld.ParSlice(), result)
	} else {
		err = errors.New(strErrNullBldPtr)
	}
	return err
}

func (dao *DBaseDAO) MakeUpdateQuery(bld *UpdateBuilder) (err error){
	if bld != nil {
		err = dao.ExecuteUpdate(bld.BuildQuery(), bld.ParSlice())
	} else {
		err = errors.New(strErrNullBldPtr)
	}
	return err
}

func (dao *DBaseDAO) MakeInsertQuery(bld *InsertBuilder) (err error){
	if bld != nil {
		if bld.IsAutoInsert() {
			err = dao.ExecuteAutoInsert(bld.BuildQuery(), bld.ParSlice(), bld.AutoKeys(),  bld.Dialect().IsReturnKey())
		} else {
			err = dao.ExecuteInsert(bld.BuildQuery(), bld.ParSlice())
		}
	} else {
		err = errors.New(strErrNullBldPtr)
	}
	return err
}

func (dao *DBaseDAO) MakeDeleteQuery(bld *DeleteBuilder) (err error){
	if bld != nil {
		err = dao.ExecuteDelete(bld.BuildQuery(), bld.ParSlice())
	} else {
		err = errors.New(strErrNullBldPtr)
	}
	return err
}

