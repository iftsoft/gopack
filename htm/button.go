package htm


type Button struct {
	widget		*Widget				// Pointer to widget for rendering html
	id			string				// Button id
	butType		string				// Button type
	class		[]string			// Button classes
	params		map[string]string	// Button params
	tag			[]string			// Button tags
	text		string				// Button text
}

// Render packs all data and executes widget render method.
func (this *Button) Render() string {
	if this.widget != nil {
		data := map[string]interface{}{
			"id":		this.id,
			"type":		this.butType,
			"classes":	this.class,
			"params":	this.params,
			"tags":		this.tag,
			"text":		this.text,
		}
		return this.widget.Render(data)
	}
	return ""
}

// SetId associates the given id to the button, overwriting any previous id.
func (this *Button) SetId(id string) *Button {
	this.id = id
	return this
}

// SetType change default button type to new type.
func (this *Button) SetType(typ string) *Button {
	this.butType = typ
	return this
}

// SetText change default button text to new text.
func (this *Button) SetText(text string) *Button {
	this.text = text
	return this
}

// AddClass adds a class to the button.
func (this *Button) AddClass(class string) *Button {
	this.class = append(this.class, class)
	return this
}

// SetParam adds a parameter (defined as key-value pair) in the field.
func (this *Button) SetParam(key, value string) *Button {
	this.params[key] = value
	return this
}

// AddTag adds a no-value parameter (e.g.: checked, disabled) to the field.
func (this *Button) AddTag(tag string) *Button {
	this.tag = append(this.tag, tag)
	return this
}

// Disabled add the "disabled" tag to the button,
// making it unresponsive in some environments (e.g. Bootstrap).
func (this *Button) Disabled() *Button {
	this.AddTag(Tag_disabled)
	return this
}

// OnClick adds onClick handler to the button
func (this *Button) OnClick(value string) *Button {
	this.SetParam(Attr_onClick, value)
	return this
}

func (this *Button) SetStyle(style string) *Button {
	switch style {
	case Style_Default:
		this.AddClass("btn-default")
	case Style_Primary:
		this.AddClass("btn-primary")
	case Style_Success:
		this.AddClass("btn-success")
	case Style_Info:
		this.AddClass("btn-info")
	case Style_Warning:
		this.AddClass("btn-warning")
	case Style_Danger:
		this.AddClass("btn-danger")
	case Style_Link:
		this.AddClass("btn-link")
	}
	return this
}


