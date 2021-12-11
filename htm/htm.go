package htm

import (
	"errors"
	"github.com/iftsoft/gopack/lla"
)

const (
	Method_POST = "POST"
	Method_GET  = "GET"
)

const (
	Style_Default = "default"
	Style_Primary = "primary"
	Style_Success = "success"
	Style_Info    = "info"
	Style_Warning = "warning"
	Style_Danger  = "danger"
	Style_Active  = "active"
	Style_Link    = "link"
)

const (
	Mode_EDIT = iota
	Mode_VIEW
	Mode_MAKE
	Mode_FIND
)

const (
	Type_LABEL      = "label"
	Type_STATIC     = "static"
	Type_BUTTON     = "button"
	Type_SUBMIT     = "submit"
	Type_RESET      = "reset"
	Type_TEXTAREA   = "textarea"
	Type_TEXT       = "text"
	Type_PASSWORD   = "password"
	Type_NUMBER     = "number"
	Type_DATE       = "date"
	Type_TIME       = "time"
	Type_DATETIME   = "datetime"
	Type_ENUM       = "enum"
	Type_SELECT     = "select"
	Type_RADIO      = "radio"
	Type_BITMASK    = "bitmask"
	Type_CHECKBOX   = "checkbox"
	Type_SEARCH     = "search"
	Type_FIELDSET   = "fieldset"
	Type_PANEL      = "panel"
	Type_BLOCK      = "block"
	Type_FORM       = "form"
	Type_GRID       = "grid"
	Type_TABGROUP   = "tabgroup"
	Type_MENUITEM   = "menuitem"
	Type_SUBMENU    = "submenu"
	Type_TOPMENU    = "topmenu"
	Type_BREADCRUMB = "breadcrumb"

	Type_HIDDEN         = "hidden"
	Type_DATETIME_LOCAL = "datetime-local"
	Type_EMAIL          = "email" // Not yet implemented
	Type_FILE           = "file"  // Not yet implemented
	Type_IMAGE          = "image" // Not yet implemented
	Type_MONTH          = "month" // Not yet implemented
	Type_RANGE          = "range"
	Type_TEL            = "tel"   // Not yet implemented
	Type_URL            = "url"   // Not yet implemented
	Type_WEEK           = "week"  // Not yet implemented
	Type_COLOR          = "color" // Not yet implemented
)

// Log Agent for handler layer logging
var htmLog lla.LogAgent

func InitLoggerHTM(level int) {
	htmLog.Init(level, "HTM")
}

func HtmPanicRecover(err *error) {
	if r := recover(); r != nil {
		htmLog.Panic("Panic is recovered: %+v", r)
		if err != nil {
			*err = errors.New("Panic have been recovered")
		}
	}
}
