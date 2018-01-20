package htm

import "strings"

///////////////////////////////////////////////////////////////////////////
//  structure.
type Grid struct {
	widget		*Widget				// Pointer to widget for rendering html
	id			string
	tbody		string
	class		[]string
	params		map[string]string
	tags		[]string			// Grid tags
	headers		[]string
}

// Render executes the internal template and renders the grid,
// returning the result as a string embeddable
// in any other template.
func (this *Grid) Render() string {
	data := map[string]interface{}{
		"id":		this.id,
		"tbody":	this.tbody,
		"classes":	this.class,
		"params":	this.params,
		"tags":		this.tags,
		"headers":	this.headers,
	}
	if this.widget != nil {
		return this.widget.Render(data)
	}
	return ""
}

func (this *Grid) ToString() string {
	html := this.Render()
	return strings.Replace(html, "\n", "" , -1)
}


// SetId associates the given id to the grid, overwriting any previous id.
func (this *Grid) SetId(id string) *Grid {
	this.id = id
	return this
}
func (this *Grid) SetDataId(id string) *Grid {
	this.tbody = id
	return this
}

// AddClass associates the provided class to the Grid.
func (this *Grid) AddClass(class string) *Grid {
	this.class = append(this.class, class)
	return this
}

// SetParam adds the given key-value pair to Grid parameters list.
func (this *Grid) SetParam(key, value string) *Grid {
	this.params[key] = value
	return this
}


