package htm


///////////////////////////////////////////////////////////////////////////
// Search
type Search struct {
	widget		*Widget				// Pointer to widget for rendering html
	id			string				// Search id
	name		string				// Search name
	value		string				// Search value
	fldType		string				// Search type
	inpClass	[]string			// Input classes
	butClass	[]string			// Button classes
	inpParams	map[string]string	// Input params
	butParams	map[string]string	// Button params
	tags		[]string			// Search tags
}

// Prepare data for template rendering
func (this *Search) dataForRender() map[string]interface{} {
	data := map[string]interface{}{
		"id":			this.id,
		"name":			this.name,
		"type":			this.fldType,
		"value":        this.value,
		"inpClas":		this.inpClass,
		"butClas":		this.butClass,
		"inpPars":		this.inpParams,
		"butPars":		this.butParams,
		"tags":			this.tags,
	}
	return data
}

// Render packs all data and executes widget render method.
func (this *Search) Render() string {
	if this.widget != nil {
		data := this.dataForRender()
		return this.widget.Render(data)
	}
	return ""
}

// Name returns the name of the field.
func (this *Search) Name() string {
	return this.name
}

// SetId associates the given id to the field, overwriting any previous id.
func (this *Search) SetId(id string) *Search {
	this.id = id
	return this
}

// SetType change default field type to new type.
func (this *Search) SetType(typ string) *Search {
	this.fldType = typ
	return this
}

// SetValue sets the value parameter for the field.
func (this *Search) SetValue(value string) *Search {
	this.value = value
	return this
}

// AddTag adds a no-value parameter (e.g.: checked, disabled) to the field.
func (this *Search) AddTag(tag string) *Search {
	this.tags = append(this.tags, tag)
	return this
}

// AddClass adds a class to the field.
func (this *Search) AddInpClass(class string) *Search {
	this.inpClass = append(this.inpClass, class)
	return this
}
func (this *Search) AddButClass(class string) *Search {
	this.butClass = append(this.butClass, class)
	return this
}

// SetParam adds a parameter (defined as key-value pair) in the field.
func (this *Search) SetInpParam(key, value string) *Search {
	this.inpParams[key] = value
	return this
}
func (this *Search) SetButParam(key, value string) *Search {
	this.butParams[key] = value
	return this
}

// SetPlaceholder adds a parameter "placeholder" in the field.
func (this *Search) SetPlaceholder(value string) *Search {
	this.SetInpParam(Attr_placeholder, value)
	return this
}
// OnClick adds a parameter "onclick" in the field.
func (this *Search) OnClick(value string) *Search {
	this.SetButParam(Attr_onClick, value)
	return this
}

