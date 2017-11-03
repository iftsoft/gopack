package ddl

import (
	"unicode"
	"reflect"
	"fmt"
	"time"
	"math/big"
	"strconv"
	"encoding/json"
	"database/sql"
	"github.com/iftsoft/gopack/lla"
)

// Format to snake string, XxYy to xx_yy
func FormatColumnName(name string) string {
	runes := []rune(name)
	size  := len(runes)
	buf := make([]rune, 0, size+8)
	for i := 0; i < size; i++ {
		buf = append(buf, unicode.ToLower(runes[i]))
		if i != size-1 && unicode.IsUpper(runes[i+1]) &&
			(unicode.IsLower(runes[i]) || unicode.IsDigit(runes[i]) ||
			(i != size-2 && unicode.IsLower(runes[i+2]))) {
			buf = append(buf, '_')
		}
	}
	return string(buf)
}


// Get column name for select query
func GetColumnName(col string, pref string, tab string, as bool) string {
	name := getFullColumnName(col, pref, tab)
	if as {
		name += " AS "
		name += getNickColumnName(col, pref, tab)
	}
	return name
}

// Fully declared column name
func getFullColumnName(col string, pref string, tab string) string {
	name := ""
	if tab != "" {
		name += tab + "."
	}
	if pref != "" {
		name += pref + "_"
	}
	if col != "" {
		name += col
	}
	return name
}

// Generated column alias
func getNickColumnName(col string, pref string, tab string) string {
	name := ""
	if tab != "" {
		name += tab + "_"
	}
	if pref != "" {
		name += pref + "_"
	}
	if col != "" {
		name += col
	}
	return name
}


// StrTo is the target string
type StrTo string

// Set string
func (f *StrTo) Set(v string) {
	if v != "" {
		*f = StrTo(v)
	} else {
		f.Clear()
	}
}

// Clear string
func (f *StrTo) Clear() {
	*f = StrTo(0x1E)
}

// Exist check string exist
func (f StrTo) Exist() bool {
	return string(f) != string(0x1E)
}

// Bool string to bool
func (f StrTo) Bool() (bool, error) {
	return strconv.ParseBool(f.String())
}

// Float32 string to float32
func (f StrTo) Float32() (float32, error) {
	v, err := strconv.ParseFloat(f.String(), 32)
	return float32(v), err
}

// Float64 string to float64
func (f StrTo) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 64)
}

// Int string to int
func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int(v), err
}

// Int8 string to int8
func (f StrTo) Int8() (int8, error) {
	v, err := strconv.ParseInt(f.String(), 10, 8)
	return int8(v), err
}

// Int16 string to int16
func (f StrTo) Int16() (int16, error) {
	v, err := strconv.ParseInt(f.String(), 10, 16)
	return int16(v), err
}

// Int32 string to int32
func (f StrTo) Int32() (int32, error) {
	v, err := strconv.ParseInt(f.String(), 10, 32)
	return int32(v), err
}

// Int64 string to int64
func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(f.String(), 10) // octal
		if !ok {
			return v, err
		}
		return ni.Int64(), nil
	}
	return v, err
}

// Uint string to uint
func (f StrTo) Uint() (uint, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint(v), err
}

// Uint8 string to uint8
func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

// Uint16 string to uint16
func (f StrTo) Uint16() (uint16, error) {
	v, err := strconv.ParseUint(f.String(), 10, 16)
	return uint16(v), err
}

// Uint32 string to uint31
func (f StrTo) Uint32() (uint32, error) {
	v, err := strconv.ParseUint(f.String(), 10, 32)
	return uint32(v), err
}

// Uint64 string to uint64
func (f StrTo) Uint64() (uint64, error) {
	v, err := strconv.ParseUint(f.String(), 10, 64)
	if err != nil {
		i := new(big.Int)
		ni, ok := i.SetString(f.String(), 10)
		if !ok {
			return v, err
		}
		return ni.Uint64(), nil
	}
	return v, err
}

// String string to string
func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
	return ""
}
func (f StrTo) Json() []byte {
	if f.Exist() {
		return []byte(string(f))
	}
	return nil
}

// ToStr interface to string
func ToStr(value interface{}) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f',10, 32)
	case float64:
		s = strconv.FormatFloat(v, 'f', 10, 64)
	case int:
		s = strconv.FormatInt(int64(v), 10)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

// ToInt64 interface to int64
func ToInt64(value interface{}) (d int64) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		panic(fmt.Errorf("ToInt64 need numeric not `%T`", value))
	}
	return
}

func CheckValueToJson(value interface{}) (interface{}) {
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return value
		} else {
			v = v.Elem()
		}
	}
	switch v.Kind() {
	// Slice or map
	case reflect.Slice, reflect.Map :
		if txt, err := json.Marshal(value); err == nil {
			return string(txt)
		} else {
			ddlLog.Warn("Can't convert value to json = %v", value)
			return nil
		}
	// Structure
	case reflect.Struct:
		switch v.Interface().(type) {
		// Time value
		case time.Time, lla.DateOnly, lla.DateTime, lla.TimeOnly, lla.FullTime:
			return value
		// Sql NullValue
		case sql.NullString, sql.NullBool, sql.NullInt64, sql.NullFloat64:
			return value
		// Undefined structure
		default:
			if txt, err := json.Marshal(value); err == nil {
				return string(txt)
			} else {
				ddlLog.Warn("Can't convert value to json = %v", value)
				return nil
			}
		}
	}
	return value
}


