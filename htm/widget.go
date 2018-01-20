package htm

import (
	"bytes"
	"text/template"
	"github.com/iftsoft/gopack/htm/templ"
)

type Renderer interface {
	Render() string
}

type Stringer interface {
	ToString() string
}

const (
	Attr_valType		= "valtype"
	Attr_onClick		= "onclick"
	Attr_placeholder	= "placeholder"
)
const (
	Tag_disabled	= "disabled"
)
const (
	Var_String	= "string"
	Var_Int		= "int"
	Var_Float	= "float"
	Var_Bool	= "bool"
	Var_Date	= "date"
	Var_Enum	= "enum"
	Var_Bitmask	= "bitmask"
)

var type_templ = map[string]string{
	Type_BUTTON : 		templ.Templ_BUTTON,
	Type_SUBMIT : 		templ.Templ_BUTTON,
	Type_RESET: 		templ.Templ_BUTTON,
	Type_LABEL : 		templ.Templ_LABEL,
	Type_STATIC : 		templ.Templ_STATIC,
	Type_TEXTAREA : 	templ.Templ_TEXTAREA,
	Type_TEXT : 		templ.Templ_INPUT,
	Type_PASSWORD : 	templ.Templ_INPUT,
	Type_NUMBER :		templ.Templ_INPUT,
	Type_DATE : 		templ.Templ_INPUT,
	Type_TIME :			templ.Templ_INPUT,
	Type_DATETIME : 	templ.Templ_INPUT,
	Type_HIDDEN : 		templ.Templ_INPUT,
	Type_ENUM : 		templ.Templ_ENUM,
	Type_SELECT : 		templ.Templ_SELECT,
	Type_RADIO : 		templ.Templ_RADIO,
	Type_BITMASK :		templ.Templ_BITMASK,
	Type_CHECKBOX :		templ.Templ_CHECKBOX,
	Type_SEARCH :		templ.Templ_SEARCH,
	Type_FIELDSET :		templ.Templ_FIELDSET,
	Type_PANEL :		templ.Templ_PANEL,
	Type_BLOCK :		templ.Templ_BLOCK,
	Type_FORM : 		templ.Templ_FORM,
	Type_GRID : 		templ.Templ_GRID,
	Type_TABGROUP : 	templ.Templ_TABGROUP,
	Type_MENUITEM :		templ.Templ_MENUITEM,
	Type_SUBMENU :		templ.Templ_SUBMENU,
	Type_TOPMENU :		templ.Templ_TOPMENU,
	Type_BREADCRUMB :	templ.Templ_BREADCRUMB,
	//Type_DATETIME_LOCAL = "datetime-local"
	//Type_EMAIL          = "email" // Not yet implemented
	//Type_FILE           = "file"  // Not yet implemented
	//Type_IMAGE          = "image" // Not yet implemented
	//Type_MONTH          = "month" // Not yet implemented
	//Type_RANGE          = "range"
	//Type_SEARCH         = "search" // Not yet implemented
	//Type_TEL            = "tel" // Not yet implemented
	//Type_URL            = "url"  // Not yet implemented
	//Type_WEEK           = "week" // Not yet implemented
	//Type_COLOR          = "color" // Not yet implemented
}

// Simple widget object that gets executed at render time.
type Widget struct {
	name  string
	temp  *template.Template
}

// Render executes the internal template and returns the result as a string.
func (w *Widget) Render(data interface{}) string {
	var s string
	if w.temp != nil {
		buf := bytes.NewBufferString(s)
		w.temp.ExecuteTemplate(buf, w.name, data)
		s = buf.String()
	}
	htmLog.Trace( "Render Widget \"%s\" to '%s'", w.name, s)
	return s
}


