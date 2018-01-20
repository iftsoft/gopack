package htm


// Field is a generic type containing all data associated to an input field.
type Field struct {
	widget		*Widget				// Pointer to widget for rendering html
	fldType		string				// Field type
	id			string				// Field id
	name		string				// Field name
	class		[]string			// Field classes
	params		map[string]string	// Field params
	tags		[]string			// Field tags
	value		string				// Field value
	helpText	string				// Field help text
	extraData	map[string]interface{}	// Extra data for inner elements
}

// Id - Value pair used to define an option for select and radio input fields.
type InputChoice struct {
	Id	string		// Option id
	Val	string		// Option value
}


// Prepare data for template rendering
func (this *Field) dataForRender() map[string]interface{} {
	data := map[string]interface{}{
		"id":           this.id,
		"name":         this.name,
		"type":         this.fldType,
		"classes":      this.class,
		"params":       this.params,
		"tags":         this.tags,
		"value":        this.value,
		"helpText":     this.helpText,
	}
	for k, v := range this.extraData {
		data[k] = v
	}
	return data
}

// Name returns the name of the field.
func (this *Field) Name() string {
	return this.name
}

// Render packs all data and executes widget render method.
func (this *Field) Render() string {
	if this.widget != nil {
		data := this.dataForRender()
		return this.widget.Render(data)
	}
	return ""
}



// AddClass adds a class to the field.
func (this *Field) AddClass(class string) *Field {
	this.class = append(this.class, class)
	return this
}

// SetId associates the given id to the field, overwriting any previous id.
func (this *Field) SetId(id string) *Field {
	this.id = id
	return this
}

// SetType change default field type to new type.
func (this *Field) SetType(typ string) *Field {
	this.fldType = typ
	return this
}

// SetParam adds a parameter (defined as key-value pair) in the field.
func (this *Field) SetParam(key, value string) *Field {
	this.params[key] = value
	return this
}

// SetPlaceholder adds a parameter "placeholder" in the field.
func (this *Field) SetPlaceholder(value string) *Field {
	this.SetParam(Attr_placeholder, value)
	return this
}
// OnClick adds a parameter "onclick" in the field.
func (this *Field) OnClick(value string) *Field {
	this.SetParam(Attr_onClick, value)
	return this
}
//func (this *Field) TypeInt() *Field {
//	this.SetParam(Attr_valType, Var_Int)
//	return this
//}
//func (this *Field) TypeFloat() *Field {
//	this.SetParam(Attr_valType, Var_Float)
//	return this
//}
//func (this *Field) TypeDate() *Field {
//	this.SetParam(Attr_valType, Var_Date)
//	return this
//}
//func (this *Field) TypeBool() *Field {
//	this.SetParam(Attr_valType, Var_Bool)
//	return this
//}
//func (this *Field) TypeString() *Field {
//	this.SetParam(Attr_valType, Var_String)
//	return this
//}

// Disabled add the "disabled" tag to the field, making it unresponsive in some environments (e.g. Bootstrap).
func (this *Field) Disabled() *Field {
	this.AddTag(Tag_disabled)
	return this
}

// AddTag adds a no-value parameter (e.g.: checked, disabled) to the field.
func (this *Field) AddTag(tag string) *Field {
	this.tags = append(this.tags, tag)
	return this
}

// SetValue sets the value parameter for the field.
func (this *Field) SetValue(value string) *Field {
	this.value = value
	return this
}

// SetHelptext saves the field helptext.
func (this *Field) SetHelpText(text string) *Field {
	this.helpText = text
	return this
}

// SetInputChoices sets an array of InputChoice objects as the possible choices for a field.
func (this *Field) SetInputChoices(choices []InputChoice) *Field {
	this.extraData["choices"] = choices
	return this
}



