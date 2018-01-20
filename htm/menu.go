package htm


///////////////////////////////////////////////////////////////////////////
// Menu item structure.
type MenuItem struct {
	widget		*Widget				// Pointer to widget for rendering html
	href		string				// Menu item URL
	text		string				// Menu item text
	glyph		string				// Menu item glyphicon
}


func (this *MenuItem) Render() string {
	if this.widget != nil {
		data := map[string]interface{}{
			"href":		this.href,
			"text":		this.text,
			"glyph":	this.glyph,
		}
		return this.widget.Render(data)
	}
	return ""
}

func (this *MenuItem) SetHref(href string) *MenuItem {
	this.href = href
	return this
}

func (this *MenuItem) SetText(text string) *MenuItem {
	this.text = text
	return this
}

func (this *MenuItem) SetGlyphIcon(glyph string) *MenuItem {
	this.glyph = glyph
	return this
}



///////////////////////////////////////////////////////////////////////////
// Sub menu structure.
type SubMenu struct {
	widget		*Widget				// Pointer to widget for rendering html
	id			string				// Menu Id
	href		string				// Menu URL
	text		string				// Menu text
	glyph		string				// Menu glyphicon
	items		[]Renderer
}

func (this *SubMenu) Render() string {
	if this.widget != nil {
		data := map[string]interface{}{
			"id":		this.id,
			"href":		this.href,
			"text":		this.text,
			"glyph":	this.glyph,
			"items":	this.items,
		}
		return this.widget.Render(data)
	}
	return ""
}

func (this *SubMenu) SetId(id string) *SubMenu {
	this.id = id
	return this
}

func (this *SubMenu) SetText(text string) *SubMenu {
	this.text = text
	return this
}

func (this *SubMenu) SetHref(href string) *SubMenu {
	this.href = href
	return this
}

func (this *SubMenu) SetGlyphIcon(glyph string) *SubMenu {
	this.glyph = glyph
	return this
}

func (this *SubMenu) Elements(elems ...Renderer) *SubMenu {
	for _, menu := range elems {
		this.items = append(this.items, menu)
	}
	return this
}


///////////////////////////////////////////////////////////////////////////
// Navbar structure.
type Navbar struct {
	widget		*Widget				// Pointer to widget for rendering html
	id			string				// Navbar Id
	href		string				// Navbar URL
	icon		string				// Navbar icon
	text		string				// Navbar text
	left		[]Renderer
	right		[]Renderer
}

func (this *Navbar) Render() string {
	if this.widget != nil {
		data := map[string]interface{}{
			"id":		this.id,
			"href":		this.href,
			"icon":		this.icon,
			"text":		this.text,
			"left":		this.left,
			"right":	this.right,
		}
		return this.widget.Render(data)
	}
	return ""
}

func (this *Navbar) SetId(id string) *Navbar {
	this.id = id
	return this
}

func (this *Navbar) SetHref(href string) *Navbar {
	this.href = href
	return this
}

func (this *Navbar) SetIcon(icon string) *Navbar {
	this.icon = icon
	return this
}

func (this *Navbar) SetText(text string) *Navbar {
	this.text = text
	return this
}

func (this *Navbar) LeftElements(elems ...Renderer) *Navbar {
	for _, menu := range elems {
		this.right = append(this.right, menu)
	}
	return this
}

func (this *Navbar) RightElements(elems ...Renderer) *Navbar {
	for _, menu := range elems {
		this.left = append(this.left, menu)
	}
	return this
}




///////////////////////////////////////////////////////////////////////////
// Breadcrumb structure.
type Breadcrumb struct {
	widget		*Widget				// Pointer to widget for rendering html
	items		[]Renderer
}

func (this *Breadcrumb) Render() string {
	if this.widget != nil {
		data := map[string]interface{}{
			"items":	this.items,
		}
		return this.widget.Render(data)
	}
	return ""
}

func (this *Breadcrumb) Elements(elems ...Renderer) *Breadcrumb {
	for _, menu := range elems {
		this.items = append(this.items, menu)
	}
	return this
}




