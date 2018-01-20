package htm

// Label is a generic container for any input field.
type Label struct {
	widget		*Widget				// Pointer to widget for rendering html
	id			string				// Label id
	forId		string				// Field for id
	class		[]string			// Label classes
	wrapper		[]string			// Field wrapper classes
	label		string				// Label text
	fields		[]Renderer
}

// Render translates a Label into HTML code and returns it as a string.
func (this *Label) Render() string {
	if this.widget != nil {
		data := map[string]interface{}{
			"id":		this.id,
			"forId":	this.forId,
			"classes":	this.class,
			"wrapper":	this.wrapper,
			"label":	this.label,
			"fields":	this.fields,
		}
		return this.widget.Render(data)
	}
	return ""
}


// AddClass saves the provided class for the fieldset.
func (this *Label) AddClass(class string) *Label {
	this.class = append(this.class, class)
	return this
}

func (this *Label) AddWrapper(class string) *Label {
	this.wrapper = append(this.wrapper, class)
	return this
}

// SetId associates the given id to the field group, overwriting any previous id.
func (this *Label) SetId(id string) *Label {
	this.id = id
	return this
}

// SetLabel saves the label to be rendered along with the field.
func (this *Label) SetLabel(label string) *Label {
	this.label = label
	return this
}

func (this *Label) SetHorizontal() *Label {
	this.AddClass("col-sm-2 col-md-3")
	this.AddWrapper("col-sm-10 col-md-9")
	return this
}



// FieldSet is a collection of fields grouped within a form.
type FieldSet struct {
	widget		*Widget			// Pointer to widget for rendering html
	id			string
	class		[]string		// FieldSet classes
	tags		[]string			// Field tags
	fields		[]Renderer
}

// Render translates a FieldSet into HTML code and returns it as a string.
func (this *FieldSet) Render() string {
	data := map[string]interface{}{
		"id":      this.id,
		"fields":  this.fields,
		"classes": this.class,
		"tags":    this.tags,
	}
	if this.widget != nil {
		return this.widget.Render(data)
	}
	return ""
}


func (this *FieldSet) SetId(id string) *FieldSet {
	this.id = id
	return this
}

// AddClass saves the provided class for the fieldset.
func (this *FieldSet) AddClass(class string) *FieldSet {
	this.class = append(this.class, class)
	return this
}

// AddTag adds a no-value parameter (e.g.: "disabled", "checked") to the fieldset.
func (this *FieldSet) AddTag(tag string) *FieldSet {
	this.tags = append(this.tags, tag)
	return this
}

// RemoveTag removes a tag from the fieldset, if it was present.
// Disable adds tag "disabled" to the fieldset, making it unresponsive in some environment (e.g.: Bootstrap).
func (this *FieldSet) Disable() *FieldSet {
	this.AddTag("disabled")
	return this
}


