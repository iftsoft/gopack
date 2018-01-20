package htm

import (
	"strconv"
	"text/template"
	"go-ticket/htm/templ"
)

///////////////////////////////////////////////////////////////////////

// Collection of widgets with different templates
type Framework struct {
	frame string
	store map[string]*Widget
}

// FrameworkBootstrap creates new Widget map with bootstrap style
func FrameworkBootstrap() *Framework {
	return &Framework { templ.Frame_Bootstrap, make(map[string]*Widget) }
}

// FrameworkBootstrap creates new Widget map with bootstrap style
func FrameworkBaseStyle() *Framework {
	return &Framework { templ.Frame_BaseStyle, make(map[string]*Widget) }
}


// GetWidget returns existing Widget or create new one and add it to map
func (this *Framework) GetWidget(fldType string) *Widget {
	name, ok := type_templ[fldType]
	if !ok {
		panic("Unknown HTML type")
	}
	if wid, ok := this.store[name]; ok {
		return wid
	}
	tmpl := template.New(name)
	html := templ.GetTemplate(this.frame, name)
	tmpl = template.Must(tmpl.Parse(html))
	wid := &Widget{name, tmpl}
	this.store[name] = wid
	return wid
}

//
func (this Framework)IsFramework(frame string) bool {
	if this.frame == frame {
		return true
	}
	return false
}


///////////////////////////////////////////////////////////////////////
// Button creates a default generic button
func (this *Framework)Button(id, text string) *Button {
	ret := &Button{
		widget:		nil,
		id:         id,
		butType:    Type_BUTTON,
		class:      []string{},
		params:     map[string]string{},
		tag:        []string{},
		text:		text,
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_BUTTON)
	}
	return ret
}

// SubmitButton creates a default generic button
func (this *Framework)SubmitButton(id, text string) *Button {
	ret := this.Button(id, text)
	ret.SetType(Type_SUBMIT)
	if this.IsFramework(templ.Frame_Bootstrap) {
		ret.SetStyle(Style_Default)
	}
	return ret
}

// Field creates an empty field of the given type and identified by name.
func (this *Framework) Field(name, fldType string) *Field {
	ret := &Field{
		widget:		nil,
		fldType:    fldType,
		id:         name,
		name:       name,
		class:      []string{},
		params:     map[string]string{},
		tags:       []string{},
		value:      "",
		helpText:   "",
		extraData:  map[string]interface{}{},
	}
	if this != nil {
		ret.widget = this.GetWidget(fldType)
	}
	return ret
}

func (this *Framework) Search(name, fldType string) *Search {
	ret := &Search{
		widget:		nil,
		id:			name,
		name:		name,
		fldType:	fldType,
		inpClass:	[]string{},
		butClass:	[]string{},
		inpParams:	map[string]string{},
		butParams:	map[string]string{},
		tags:		[]string{},
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_SEARCH)
	}
	return ret
}

func (this *Framework) Label(text string, elems ...Renderer) *Label {
	ret := &Label{
		widget:	nil,
		id:		"",
		forId:	"",
		label:	text,
		class:      []string{},
		wrapper:	[]string{},
		fields:		elems,
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_LABEL)
	}
	return ret
}

// NewFieldSet creates and returns a new FieldSet with the given name and list of fields.
// Every method for FieldSet objects returns the object itself, so that call can be chained.
func (this *Framework) FieldSet(elems ...Renderer) *FieldSet {
	ret := &FieldSet{
		widget:	nil,
		id: 	"",
		class:	[]string{},
		tags:	[]string{},
		fields:	elems,
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_FIELDSET)
	}
	return ret
}

func (this *Framework) Panel(elems ...Renderer) *Panel {
	ret := &Panel {
		widget:		nil,
		id:			"",
		header:		"",
		footer:		nil,
		class:		[]string{},
		params:		map[string]string{},
		tags:		[]string{},
		fields:		elems,
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_PANEL)
	}
	return ret
}

func (this *Framework) Block() *Block {
	ret := &Block {
		widget:		nil,
		id:			"",
		fields:		[]Renderer{},
		left:		[]Renderer{},
		right:		[]Renderer{},
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_BLOCK)
	}
	return ret
}

func (this *Framework) TabGroup(name string) *TabGroup {
	ret := &TabGroup {
		widget:		nil,
		name:		name,
		panes:		[]TabPane{},
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_TABGROUP)
	}
	return ret
}

func (this *Framework) Form(id string, elems ...Renderer) *Form {
	ret := &Form {
		widget:	nil,
		id:		id,
		fields: elems,
		class:	[]string{},
		params:	map[string]string{},
		tags:	[]string{},
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_FORM)
	}
	return ret
}

func (this *Framework) Grid(heads ...string) *Grid {
	ret := &Grid {
		widget:		nil,
		id:			"",
		tbody:		"",
		class:		[]string{},
		params:		map[string]string{},
		tags:		[]string{},
		headers:	heads,
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_GRID)
	}
	return ret
}

func (this *Framework) MenuItem(href, text string) *MenuItem {
	ret := &MenuItem {
		widget:		nil,
		href:		href,
		text:		text,
		glyph:		"",
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_MENUITEM)
	}
	return ret
}

func (this *Framework) SubMenu(text string, elems []Renderer) *SubMenu {
	ret := &SubMenu {
		widget:		nil,
		id:			"",
		href:		"#",
		text:		text,
		glyph:		"",
		items:		elems,
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_SUBMENU)
	}
	return ret
}

func (this *Framework) Navbar(href, icon string, left []Renderer, right []Renderer) *Navbar {
	ret := &Navbar {
		widget:		nil,
		id:			"",
		href:		href,
		icon:		icon,
		text:		"",
		left:		left,
		right:		right,
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_TOPMENU)
	}
	return ret
}

func (this *Framework) Breadcrumb(elems []Renderer) *Breadcrumb {
	ret := &Breadcrumb {
		widget:		nil,
		items:		elems,
	}
	if this != nil {
		ret.widget = this.GetWidget(Type_BREADCRUMB)
	}
	return ret
}



// Hidden creates a default hidden input field based on the provided name.
func (this *Framework)Hidden(name string) *Field {
	ret := this.Field(name, Type_HIDDEN)
	return ret
}

// Static returns a static field with the provided name and content
func (this *Framework)Static(name, content string) *Field {
	ret := this.Field(name, Type_STATIC)
	ret.SetValue(content)
	return ret
}

// NumberField craetes a default number field with the provided name.
func (this *Framework)Number(name string) *Field {
	ret := this.Field(name, Type_NUMBER)
	return ret
}

// TextField creates a default text input field based on the provided name.
func (this *Framework)Text(name string) *Field {
	ret := this.Field(name, Type_TEXT)
	return ret
}

// PasswordField creates a default password text input field based on the provided name.
func (this *Framework)Password(name string) *Field {
	ret := this.Field(name, Type_PASSWORD)
	return ret
}

// DateField creates a default date input field
func (this *Framework)Date(name string) *Field {
	ret := this.Field(name, Type_DATE)
	return ret
}

// TimeField creates a default time input field
func (this *Framework)Time(name string) *Field {
	ret := this.Field(name, Type_TIME)
	return ret
}

// DatetimeField creates a default datetime input field
func (this *Framework)Datetime(name string) *Field {
	ret := this.Field(name, Type_DATETIME)
	return ret
}

// TextAreaField creates a default textarea input field based on the provided name and dimensions.
func (this *Framework)TextArea(name string, rows int) *Field {
	ret := this.Field(name, Type_TEXTAREA)
	ret.SetParam("rows", strconv.Itoa(rows))
	//	ret.SetParam("cols", strconv.Itoa(cols))
	return ret
}


// Checkbox creates a default checkbox field with the provided name.
func (this *Framework)Checkbox(name string) *Field {
	ret := this.Field(name, Type_CHECKBOX)
	return ret
}

func (this *Framework)Bitmask(name string, choices []InputChoice) *Field {
	ret := this.Field(name, Type_BITMASK)
	ret.SetInputChoices(choices)
	return ret
}

// RadioField creates a default radio button input field with the provided name and list of choices.
func (this *Framework)Radio(name string, choices []InputChoice) *Field {
	ret := this.Field(name, Type_RADIO)
	ret.SetInputChoices(choices)
	return ret
}

// SelectField creates a default select input field with the provided name and map of choices. Choices for SelectField are grouped
// by name (if <optgroup> is needed); "" group is the default one and does not trigger a <optgroup></optgroup> rendering.
func (this *Framework)Select(name string, choices []InputChoice) *Field {
	ret := this.Field(name, Type_SELECT)
	ret.SetInputChoices(choices)
	return ret
}

func (this *Framework)Enum(name string) *Field {
	ret := this.Field(name, Type_ENUM)
	return ret
}

func (this *Framework)SearchDate(name string) *Search {
	ret := this.Search(name, Type_DATE)
	return ret
}

func (this *Framework)SearchText(name string) *Search {
	ret := this.Search(name, Type_TEXT)
	return ret
}
