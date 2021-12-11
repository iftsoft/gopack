package ddl

import (
	"reflect"
	"strings"
)

// Reserved words
const (
	Tag_Ddl  = "ddl"
	Tag_Skip = "-"
	Tag_Col  = "col"
	Tag_Key  = "key"
	Tag_Auto = "auto"
	Tag_Inc  = "inc"
	Tag_Ref  = "ref"
)

// Field attributes
const Field_Skip = 0
const (
	Field_Col = 1 << iota
	Field_Key
	Field_Auto
	Field_Inc
	Field_Ref
)
const Skip_Select = Field_Skip
const Skip_Return = Field_Key | Field_Auto
const Skip_Insert = Field_Auto | Field_Ref
const Skip_Update = Field_Key | Field_Auto | Field_Ref

///////////////////////////////////////////////////////////////////////
//
// Field information structure
type FieldInfo struct {
	DdlType int
	FldKind reflect.Kind
	FldName string
	ColName string
}

// Init field information
func (info *FieldInfo) InitFieldInfo(field reflect.StructField) {
	if info == nil {
		return
	}
	info.DdlType = Field_Skip
	info.FldKind = field.Type.Kind()
	info.FldName = field.Name
	info.ColName = FormatColumnName(field.Name)
	tag := field.Tag.Get(Tag_Ddl)
	if len(tag) > 0 {
		info.ParseFieldTag(tag)
	} else {
		info.DdlType = Field_Col
	}
	ddlLog.Trace("Field name:%s, column:%s, kind:%s, type:%d", info.FldName, info.ColName, info.FldKind.String(), info.DdlType)
}

// Parse field attributes from "ddl" tag
func (info *FieldInfo) ParseFieldTag(tag string) {
	if info == nil {
		return
	}
	if tag == Tag_Skip {
		info.DdlType = Field_Skip
		return
	}
	for i, v := range strings.Split(tag, ",") {
		if i == 0 && v != "" {
			info.ColName = v
		}
		if i > 0 {
			switch v {
			case Tag_Col:
				info.DdlType |= Field_Col
			case Tag_Key:
				info.DdlType |= Field_Key
			case Tag_Auto:
				info.DdlType |= Field_Auto
			case Tag_Ref:
				info.DdlType |= Field_Ref
			case Tag_Inc:
				info.DdlType |= Field_Inc
			}
		}
	}
	if info.DdlType == Field_Skip {
		info.DdlType = Field_Col
	}
}
