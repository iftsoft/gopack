package ddl

import (
	"reflect"
	"database/sql"
	"errors"
	"time"
	"encoding/json"
	"go-ticket/lla"
)


const (
	strErrMakeBuilder	= "Can't create query builder"
	//	strErrNullBasePtr	= "Invalid pointer to DB connect"
	strErrNullDaoPtr	= "Invalid pointer to DAO object"
	strErrNullBldPtr	= "Invalid pointer to SQL builder"
	strErrNullDataPtr	= "Attempt to load into an invalid pointer"
	strErrDataIsSlice	= "Unit target is slice"
	strErrDataIsUnit	= "List target is not slice"
)

//var ddlErrMakeBuilder	= errors.New("Can't create query builder")
var ddlErrNullDaoPtr	= errors.New(strErrNullDaoPtr)
var ddlErrNullBldPtr	= errors.New(strErrNullBldPtr)
//var ddlErrNullDataPtr	= errors.New(strErrNullDataPtr)
//var ddlErrDataIsSlice	= errors.New(strErrDataIsSlice)
//var ddlErrDataIsUnit	= errors.New(strErrDataIsUnit)


// Fetch columns data from selected record and fill object structure
func fetchReturnedRow(rows *sql.Rows, m map[string]reflect.Value) (int64, error) {
	defer rows.Close()

	column, err := rows.Columns()
	if err != nil {
		return 0, err
	}
	// Make column data holders
	colsNum := len(column)
	refs := getDummyColumnValues(colsNum)

	var count int64 = 0
	if rows.Next() {
		// Scan column data from record
		err = rows.Scan(refs...)
		if err != nil {
			return 0, err
		}
		// Fill value map with column data
		err = iterateColumnValues(refs, column, m)
		if err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

// Fetch columns data from selected record and fill object structure
func fetchSelectedRow(rows *sql.Rows, value interface{}) (int64, error) {
	defer rows.Close()

	column, err := rows.Columns()
	if err != nil {
		return 0, err
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		panic(strErrNullDataPtr)
	}
	obj := v.Elem()
	if obj.Kind() == reflect.Slice {
		panic(strErrDataIsSlice)
	}

	// Make column data holders
	colsNum := len(column)
	refs := getDummyColumnValues(colsNum)
	m := DefineColumnValues(obj)

	var count int64 = 0
	if rows.Next() {
		// Scan column data from record
		err = rows.Scan(refs...)
		if err != nil {
			return 0, err
		}
		// Fill Object with column data
		err = iterateColumnValues(refs, column, m)
		if err != nil {
			return 0, err
		}
		count++
	}
	return count, nil
}

// Fetch columns data from record set and fill objects slice
func fetchSearchedRows(rows *sql.Rows, value interface{}) (int64, error) {
	defer rows.Close()

	column, err := rows.Columns()
	if err != nil {
		return 0, err
	}

	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		panic(strErrNullDataPtr)
	}
	list := v.Elem()
	if list.Kind() != reflect.Slice {
		panic(strErrDataIsUnit)
	}
	var count int64 = 0
	colsNum := len(column)

	for rows.Next() {
		// Make column data holders and scan data from current row
		refs := getDummyColumnValues(colsNum)
		err = rows.Scan(refs...)
		if err != nil {
			return 0, err
		}
		// Create new Object for slice
		elem := reflect.New(list.Type().Elem()).Elem()
		m := DefineColumnValues(elem)
		// Fill Object with column data
		err = iterateColumnValues(refs, column, m)
		if err != nil {
			return 0, err
		}
		count++
		list.Set(reflect.Append(list, elem))
	}
	return count, nil
}

// Return slice of empty interfaces for column data holders
func getDummyColumnValues(colsNum int) []interface{} {
	refs := make([]interface{}, colsNum)
	for i := range refs {
		var ref interface{}
		refs[i] = &ref
	}
	return refs
}


// Fill Object structure with column data
func iterateColumnValues(refs []interface{}, column []string, m map[string]reflect.Value) error {
	// Iterate through columns in record
	for i, col := range column {
		if dst, ok := m[col]; ok {
			// Get column data interface
			src := reflect.Indirect(reflect.ValueOf(refs[i]))
			// Fill field value in object
			err := setUnitColumnValue(src, dst, col)
			if err != nil {
				return err
			}
		}
	}
	return nil
}


func DefineColumnValues(value reflect.Value) map[string]reflect.Value {
	// Init and fill map of Value holders
	m := make(map[string]reflect.Value)
	iterateObjectValues(value, m, "", "")
	return m
}

// Recursive iterate through object structure and fill column map with value holders
func iterateObjectValues(value reflect.Value, m map[string]reflect.Value, tab string, pref string) {
	switch value.Kind() {
	case reflect.Ptr:
		if value.IsNil() {
			value.Set(reflect.New(value.Type().Elem()))
		}
		if !value.IsNil() {
			// Recursive call for pointer to struct
			iterateObjectValues(value.Elem(), m, tab, pref)
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
				iterateObjectValues(value.Field(i), m, tab, pref)
				continue
			}
			info := FieldInfo{}
			info.InitFieldInfo(field)
			if info.DdlType == Field_Skip {
				// Skip ignored field
				continue
			}
			if (info.DdlType & Field_Inc) == Field_Inc {
				// Recursive call for regular include struct
				iterateObjectValues(value.Field(i), m, tab, info.ColName)
				continue
			}
			if (info.DdlType & Field_Ref) == Field_Ref {
				// Recursive call for reference to other struct
				iterateObjectValues(value.Field(i), m, info.ColName, "")
				continue
			}
			// Add struct field to column map
			key := getNickColumnName(info.ColName, pref, tab)
			if _, ok := m[key]; !ok {
				m[key] = value.Field(i)
			}
		}
	}
}

// Copy data from record column to object field
func setUnitColumnValue(val reflect.Value, dst reflect.Value, col string) error {
	srcType := val.Type()
	dstType := dst.Type()
	dstKind := dst.Kind()
	src := val.Interface()
	ddlLog.Trace("Column %s, Target = %v (%v), Source = %v, Value = %v", col, dstType, dstKind, srcType, src)

	var fld reflect.Value
	// Find or create actual field
	if dstKind == reflect.Ptr {
		if src == nil {
			dst.Set(reflect.Zero(dst.Type()))
			return nil
		}
		if dst.IsNil() {
			dst.Set(reflect.New(dst.Type().Elem()))
		}
		fld = dst.Elem()
	} else {
		fld = dst
	}

	var str *StrTo
	// Check for string type of source data
	switch v := src.(type) {
	case []byte:
		s := StrTo(string(v))
		str = &s
	case string:
		s := StrTo(v)
		str = &s
	}

	fldKind := fld.Kind()
	// Check field type
	switch fldKind {
	// String
	case reflect.String:
		if src == nil {
			fld.Set(reflect.Zero(fld.Type()))
		} else {
			if str != nil {
				fld.SetString(str.String())
			} else {
				fld.SetString(ToStr(src))
			}
		}
	// Boolean
	case reflect.Bool:
		if src == nil {
			fld.Set(reflect.Zero(fld.Type()))
		} else {
			if str != nil {
				b, err := str.Bool()
				if err != nil {
					return err
				}
				fld.SetBool(b)
			} else {
				fld.SetBool(src.(bool))
			}
		}
	// Integer
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if src == nil {
			fld.Set(reflect.Zero(fld.Type()))
		} else {
			if str != nil {
				i, err := str.Int64()
				if err != nil {
					return err
				}
				fld.SetInt(i)
			} else {
				fld.SetInt(src.(int64))
			}
		}
	// Unsigned
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if src == nil {
			fld.Set(reflect.Zero(fld.Type()))
		} else {
			if str != nil {
				u, err := str.Uint64()
				if err != nil {
					return err
				}
				fld.SetUint(u)
			} else {
				fld.SetUint(src.(uint64))
			}
		}
	// Float
	case reflect.Float32, reflect.Float64:
		if src == nil {
			fld.Set(reflect.Zero(fld.Type()))
		} else {
			if str != nil {
				v, err := str.Float64()
				if err != nil {
					return err
				}
				fld.SetFloat(v)
			} else {
				fld.SetFloat(src.(float64))
			}
		}
	// Structure
	case reflect.Struct:
		switch fld.Interface().(type) {
		case time.Time, lla.DateOnly, lla.DateTime, lla.TimeOnly, lla.FullTime:
//			ddlLog.Info("Column %s, Target = %v (%v), Source = %v, Value = %v", col, dstType, dstKind, srcType, src)
			if src == nil {
				fld.Set(reflect.Zero(fld.Type()))
			} else {
				if tm, ok := src.(time.Time); ok {
					switch fld.Interface().(type) {
					case time.Time:
						fld.Set(reflect.ValueOf(tm))
					case lla.DateOnly:
						fld.Set(reflect.ValueOf(lla.DateOnly{tm}))
					case lla.TimeOnly:
						fld.Set(reflect.ValueOf(lla.TimeOnly{tm}))
					case lla.DateTime:
						fld.Set(reflect.ValueOf(lla.DateTime{tm}))
					case lla.FullTime:
						fld.Set(reflect.ValueOf(lla.FullTime{tm}))
					}
				} else {
					fld.Set(reflect.Zero(fld.Type()))
				}
			}
		// NullString
		case sql.NullString:
			if ns, ok := src.(sql.NullString); ok {
				if src == nil {
					ns.Valid = false
				} else {
					if str != nil {
						ns.String = str.String()
						ns.Valid  = true
					} else {
						ns.String = ToStr(src)
						ns.Valid  = true
					}
				}
				fld.Set(reflect.ValueOf(ns))
			} else {
				fld.Set(reflect.Zero(fld.Type()))
			}
		// NullBool
		case sql.NullBool:
			if nb, ok := src.(sql.NullBool); ok {
				if src == nil {
					nb.Valid = false
				} else {
					if str != nil {
						b, err := str.Bool()
						if err != nil {
							nb.Valid = false
						} else {
							nb.Bool = b
							nb.Valid = true
						}
					} else {
						nb.Bool = src.(bool)
						nb.Valid = true
					}
				}
				fld.Set(reflect.ValueOf(nb))
			} else {
				fld.Set(reflect.Zero(fld.Type()))
			}
		// NullInt
		case sql.NullInt64:
			if ni, ok := src.(sql.NullInt64); ok {
				if src == nil {
					ni.Valid = false
				} else {
					if str != nil {
						i, err := str.Int64()
						if err != nil {
							ni.Valid = false
						} else {
							ni.Int64 = i
							ni.Valid = true
						}
					} else {
						ni.Int64 = src.(int64)
						ni.Valid = true
					}
				}
				fld.Set(reflect.ValueOf(ni))
			} else {
				fld.Set(reflect.Zero(fld.Type()))
			}
		// NullFloat
		case sql.NullFloat64:
			if nf, ok := src.(sql.NullFloat64); ok {
				if src == nil {
					nf.Valid = false
				} else {
					if str != nil {
						f, err := str.Float64()
						if err != nil {
							nf.Valid = false
						} else {
							nf.Float64 = f
							nf.Valid = true
						}
					} else {
						nf.Float64 = src.(float64)
						nf.Valid = true
					}
				}
				fld.Set(reflect.ValueOf(nf))
			} else {
				fld.Set(reflect.Zero(fld.Type()))
			}
		// Undefined structure
		default:
			if str == nil {
				fld.Set(reflect.Zero(fld.Type()))
			} else {
				val := reflect.New(fld.Type())
				if err := json.Unmarshal(str.Json(), val.Interface()); err == nil {
					fld.Set(val)
				} else {
					fld.Set(reflect.Zero(fld.Type()))
				}
			}
		}
	// Slice or map
	case reflect.Slice, reflect.Map :
		if str == nil {
			fld.Set(reflect.Zero(fld.Type()))
		} else {
			val := reflect.New(fld.Type())
			if err := json.Unmarshal(str.Json(), val.Interface()); err == nil {
				fld.Set(val)
			} else {
				fld.Set(reflect.Zero(fld.Type()))
			}
		}
	}
	return nil
}


