package htm

import (
	"strings"
)

// Form structure.
type Form struct {
	widget		*Widget			// Pointer to widget for rendering html
	id			string
	fields		[]Renderer
	class		[]string
	params		map[string]string
	tags		[]string			// Field tags
}


// Render executes the internal template and renders the form,
// returning the result as a string embeddable
// in any other template.
func (this *Form) Render() string {
	data := map[string]interface{}{
		"id":      this.id,
		"fields":  this.fields,
		"classes": this.class,
		"params":  this.params,
		"tags":    this.tags,
	}
	if this.widget != nil {
		return this.widget.Render(data)
	}
	return ""
}

func (this *Form) ToString() string {
	html := this.Render()
	return strings.Replace(html, "\n", "" , -1)
}


// SetId set the given id to the form.
func (this *Form) SetId(id string) *Form {
	this.id = id
	return this
}

// Elements adds the provided elements to the form.
func (this *Form) Elements(elems ...Renderer) *Form {
	for _, field := range elems {
		this.fields = append(this.fields, field)
	}
	return this
}

// AddClass associates the provided class to the Form.
func (this *Form) AddClass(class string) *Form {
	this.class = append(this.class, class)
	return this
}

func (this *Form) SetHorizontal() *Form {
	this.AddClass("form-horizontal")
	return this
}

// SetParam adds the given key-value pair to form parameters list.
func (this *Form) SetParam(key, value string) *Form {
	this.params[key] = value
	return this
}



///////////////////////////////////////////////////////////////////////////
// Tabs group structure.

type TabPane struct {
	Name		string
	Head		string
	Fields		[]Renderer
	Index		int
}

type TabGroup struct {
	widget		*Widget			// Pointer to widget for rendering html
	name		string
	panes		[]TabPane
}


// Name returns the name of the form.
func (this *TabGroup) Name() string {
	return this.name
}

func (this *TabGroup) Render() string {
	data := map[string]interface{}{
		"name":    this.name,
		"panes":   this.panes,
	}
	if this.widget != nil {
		return this.widget.Render(data)
	}
	return ""
}

func (this *TabGroup) ToString() string {
	html := this.Render()
	out := strings.Replace(string(html), "\n", "", -1)
	return out
}


func (this *TabGroup) AddTab(name, head string, elems ...Renderer) *TabGroup {
	tab := TabPane{ name, head, []Renderer{}, len(this.panes) }
	for _, field := range elems {
		tab.Fields = append(tab.Fields, field)
	}
	this.panes = append(this.panes, tab)
	return this
}





