package htm

import "strings"

///////////////////////////////////////////////////////////////////////////
// Panel structure.
type Panel struct {
	widget		*Widget				// Pointer to widget for rendering html
	id			string
	header		string
	footer		Renderer
	fields		[]Renderer
	class		[]string
	params		map[string]string
	tags		[]string			// Panel tags
}

// Render executes the internal template and renders the panel,
// returning the result as a string embeddable
// in any other template.
func (this *Panel) Render() string {
	data := map[string]interface{}{
		"id":		this.id,
		"header":	this.header,
		"footer":	this.footer,
		"fields":	this.fields,
		"classes":	this.class,
		"params":	this.params,
		"tags":		this.tags,
	}
	if this.widget != nil {
		return this.widget.Render(data)
	}
	return ""
}

func (this *Panel) ToString() string {
	html := this.Render()
	return strings.Replace(html, "\n", "" , -1)
}

// SetId associates the given id to the panel, overwriting any previous id.
func (this *Panel) SetId(id string) *Panel {
	this.id = id
	return this
}

func (this *Panel) SetHeader(header string) *Panel {
	this.header = header
	return this
}

func (this *Panel) SetFooter(footer Renderer) *Panel {
	this.footer = footer
	return this
}

// Elements adds the provided elements to the panel.
func (this *Panel) Elements(elems ...Renderer) *Panel {
	for _, field := range elems {
		this.fields = append(this.fields, field)
	}
	return this
}

// AddClass associates the provided class to the panel.
func (f *Panel) AddClass(class string) *Panel {
	f.class = append(f.class, class)
	return f
}

// SetParam adds the given key-value pair to form parameters list.
func (this *Panel) SetParam(key, value string) *Panel {
	this.params[key] = value
	return this
}

func (this *Panel) SetStyle(style string) *Panel {
	switch style {
	case Style_Default:
		this.AddClass("panel-default")
	case Style_Primary:
		this.AddClass("panel-primary")
	case Style_Success:
		this.AddClass("panel-success")
	case Style_Info:
		this.AddClass("panel-info")
	case Style_Warning:
		this.AddClass("panel-warning")
	case Style_Danger:
		this.AddClass("panel-danger")
	}
	return this
}



///////////////////////////////////////////////////////////////////////////
// Block structure.
type Block struct {
	widget		*Widget				// Pointer to widget for rendering html
	id			string
	fields		[]Renderer
	left		[]Renderer
	right		[]Renderer
}

// Render executes the internal template and renders the panel,
// returning the result as a string embeddable
// in any other template.
func (this *Block) Render() string {
	data := map[string]interface{}{
		"id":		this.id,
		"fields":	this.fields,
		"left": 	this.left,
		"right":	this.right,
	}
	if this.widget != nil {
		return this.widget.Render(data)
	}
	return ""
}

func (this *Block) ToString() string {
	html := this.Render()
	return strings.Replace(html, "\n", "" , -1)
}

// SetId associates the given id to the block, overwriting any previous id.
func (this *Block) SetId(id string) *Block {
	this.id = id
	return this
}

// Elements adds the provided elements to the block.
func (this *Block) Elements(elems ...Renderer) *Block {
	for _, field := range elems {
		this.fields = append(this.fields, field)
	}
	return this
}
func (this *Block) AddLeft(elems ...Renderer) *Block {
	for _, field := range elems {
		this.left = append(this.left, field)
	}
	return this
}
func (this *Block) AddRight(elems ...Renderer) *Block {
	for _, field := range elems {
		this.right = append(this.right, field)
	}
	return this
}


