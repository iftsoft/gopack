package templ

const (
	Frame_BaseStyle	= "BaseStyle"
	Frame_Bootstrap	= "Bootstrap"
	Frame_Material	= "Material"
)
const (
	Templ_LABEL			= "label"
	Templ_INPUT			= "input"
	Templ_STATIC		= "static"
	Templ_BUTTON		= "button"
	Templ_TEXTAREA		= "textarea"
	Templ_ENUM			= "enum"
	Templ_SELECT		= "select"
	Templ_RADIO			= "radio"
	Templ_BITMASK		= "bitmask"
	Templ_CHECKBOX		= "checkbox"
	Templ_SEARCH        = "search"
	Templ_FIELDSET		= "fieldset"
	Templ_PANEL			= "panel"
	Templ_BLOCK			= "block"
	Templ_FORM			= "form"
	Templ_GRID			= "grid"
	Templ_TABGROUP		= "tabgroup"
	Templ_MENUITEM		= "menuitem"
	Templ_SUBMENU		= "submenu"
	Templ_TOPMENU		= "topmenu"
	Templ_BREADCRUMB	= "breadcrumb"

)

var base_style = map[string]string {
	Templ_LABEL :		base_label,
	Templ_INPUT :		base_input,
	Templ_STATIC :		base_static,
	Templ_BUTTON :		base_button,
	Templ_TEXTAREA :	base_textarea,
	Templ_ENUM :		base_enum,
	Templ_SELECT :		base_select,
	Templ_RADIO :		base_radio,
	Templ_BITMASK:		base_bitmask,
	Templ_CHECKBOX :	base_checkbox,
	Templ_SEARCH :		base_search,
	Templ_FIELDSET:		base_fieldset,
	Templ_PANEL :		base_panel,
	Templ_BLOCK :		base_block,
	Templ_FORM :		base_form,
	Templ_GRID :		base_grid,
	Templ_TABGROUP :	base_tabgroup,
	Templ_MENUITEM :	base_menuitem,
	Templ_SUBMENU :		base_submenu,
	Templ_TOPMENU :		base_topmenu,
	Templ_BREADCRUMB:	base_breadcrumb,
}

var boot_style = map[string]string {
	Templ_LABEL :		boot_label,
	Templ_INPUT :		boot_input,
	Templ_STATIC :		boot_static,
	Templ_BUTTON :		boot_button,
	Templ_TEXTAREA :	boot_textarea,
	Templ_ENUM :		boot_enum,
	Templ_SELECT :		boot_select,
	Templ_RADIO :		boot_radio,
	Templ_BITMASK:		boot_bitmask,
	Templ_CHECKBOX :	boot_checkbox,
	Templ_SEARCH :		boot_search,
	Templ_FIELDSET:		boot_fieldset,
	Templ_PANEL :		boot_panel,
	Templ_BLOCK :		boot_block,
	Templ_FORM :		boot_form,
	Templ_GRID :		boot_grid,
	Templ_TABGROUP :	boot_tabgroup,
	Templ_MENUITEM :	boot_menuitem,
	Templ_SUBMENU :		boot_submenu,
	Templ_TOPMENU :		boot_topmenu,
	Templ_BREADCRUMB:	boot_breadcrumb,
}



func GetTemplate(frame, name string) string {
	switch frame {
	case Frame_BaseStyle :
		if str, ok := base_style[name]; ok {
			return str
		}
	case Frame_Bootstrap :
		if str, ok := boot_style[name]; ok {
			return str
		}
	}
	return ""
}
