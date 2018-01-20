package templ

var base_button = `{{define "button"}}` +
	`<button type="{{.type}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .classes}} class="{{range .classes}}{{.}} {{end}}"{{end}}` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .css}} style="{{range $k, $v := .css}}{{$k}}:{{$v}}; {{end}}"{{end}}` +
	`{{range .tags}} {{.}}{{end}}` +
	`>{{.text}}</button> {{end}}`

var base_static = `{{define "static"}}` +
	`{{if .label}}<label{{if .labelClasses}} class="{{range .labelClasses}}{{.}} {{end}}"{{end}}` +
	`{{if .id}} for="{{.id}}"{{end}}>{{.label}}</label>{{end}}` +
	`<p name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .classes}} class="{{range .classes}}{{.}} {{end}}"{{end}}` +
	`{{if .css}} style="{{range $k, $v := .css}}{{$k}}:{{$v}}; {{end}}"{{end}}` +
	`{{range $k,$v := .tags}} {{$k}}{{end}}>{{.text}}</p> {{end}}`

var base_search = `{{define "search"}}` +
	`<div class="input-group">` +
	`<input type="date" class="form-control" placeholder="Search" name="search">` +
	`<div class="input-group-btn">` +
	`<button class="btn btn-default" type="button">` +
	`<i class="glyphicon glyphicon-search"></i></button>` +
	`</div>` +
	`</div>` +
	` {{end}}`

var base_input = `{{define "input"}}` +
	`{{if .label}}<label{{if .labelClasses}} class="{{range .labelClasses}}{{.}} {{end}}"{{end}}` +
	`{{if .id}} for="{{.id}}"{{end}}>{{.label}}</label>{{end}}` +
	`<input type="{{.type}}" name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .classes}} class="{{range .classes}}{{.}} {{end}}"{{end}}` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .css}} style="{{range $k, $v := .css}}{{$k}}:{{$v}}; {{end}}"{{end}}` +
	`{{range $k,$v := .tags}} {{$k}}{{end}}` +
	`{{if .value}} value="{{.value}}"{{end}}>`

var base_textarea = `{{define "textarea"}}` +
	`{{if .label }}<label{{if .labelClasses}} class="{{range .labelClasses}}{{.}} {{end}}"{{end}}` +
	`{{if .id}} for="{{.id}}"{{end}}>{{.label}}</label>{{end}}` +
	`<textarea name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .classes }} class="{{range .classes}}{{.}} {{end}}"{{end}}` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .css}} style="{{range $k, $v := .css}}{{$k}}:{{$v}}; {{end}}"{{end}}` +
	`{{range $k,$v := .tags}} {{$k}}{{end}}>` +
	`{{.text}}</textarea> {{end}}`

var base_enum = `{{define "enum"}}` +
	`<div class="form-group{{if .errors}} has-error{{end}}">` +
	`{{if .label}}<label class="control-label col-sm-3` +
	`{{if .labelClasses}}{{range .labelClasses}} {{.}}{{end}}{{end}}"` +
	`{{if .id}} for="{{.id}}"{{end}}>{{.label}}</label>{{end}}` +
	`<div class="col-sm-9">` +
	`<select name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="form-control {{if .classes}}{{range .classes}}{{.}} {{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .css}} style="{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}"{{end}}` +
	`{{range $k,$v := .tags}} {{$k}}{{end}}>{{ $p := . }}` +
	`{{range .choices}}` +
	`<option value="{{.Id}}">` +
	`{{.Val}}</option>{{end}}` +
	`</select>` +
	`</div></div> {{end}}`

var base_select = `{{define "select"}}` +
	`{{if .label}}<label{{if .labelClasses}} class="{{range .labelClasses}}{{.}} {{end}}"{{end}}` +
	`{{if .id}} for="{{.id}}"{{end}}>{{.label}}</label>{{end}}` +
	`<select name="{{.name}}"` +
	`{{if .classes}} class="{{range .classes}}{{.}} {{end}}"{{end}}` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .css}} style="{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}"{{end}}` +
	`{{range $k,$v := .tags}} {{$k}}{{end}}>` +
	`{{$p := . }}{{range $k, $v := .choices}}{{if $k }}<optgroup label="{{$k}}">{{end}}` +
	`{{range $v}}<option value="{{.Id}}"{{if $p.tags.multiple }}{{$id := .Id}}{{range $k2, $p2 := $p.multValues}}{{if eq $k2 $id}} selected{{end}}{{end}}` +
	`{{else}}{{ if eq $p.value .Id}} selected{{end}}{{end}}>{{.Val}}</option>{{end}}` +
	`{{if $k}}</optgroup>{{end}}{{end}}` +
	`</select> {{end}}`

var base_checkbox = `{{define "checkbox"}}` +
	`<input type="checkbox" name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .classes}} class="{{range .classes}}{{.}} {{end}}"{{end}}` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .css}} style="{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}"{{end}}` +
	`{{range $k,$v := .tags}} {{$k}}{{end}}` +
	`>{{.label}}<br> {{end}}`

var base_bitmask = `{{define "bitmask"}}{{$p := . }}` +
	`<div class="checkbox form-control" style="float:left; height:auto;"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}>` +
	`{{range .choices }}` +
	`<label class="col-xs-12 col-sm-6 col-md-4 col-lg-3">` +
	`<input type="checkbox" name="{{$p.name}}"` +
	`{{if $p.classes}} class="{{range $p.classes}}{{.}} {{end}}"{{end}}` +
	`{{if $p.params}}{{range $k, $v := $p.params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if $p.tags}}{{range $p.tags}} {{.}}{{end}}{{end}}` +
	` value="{{.Id}}">{{.Val}}</label>` +
	`{{end}}` +
	`</div> {{end}}`

var base_radio = `{{define "radio"}}{{$p := .}}` +
	`{{range $k, $v := .choices}}{{if $v}}` +
	`<label{{if $p.labelClasses}} class="{{range $p.labelClasses}}{{.}} {{end}}"{{end}} ` +
	`for="{{.Id}}">{{.Val}}</label>` +
	`<input type="radio" name="{{$p.name}}"` +
	`{{if $p.classes}} class="{{range $p.classes}}{{.}} {{end}}"{{end}} id="{{.Id}}"` +
	`{{if $p.params}}{{range $k2, $v2 := $p.params}} {{$k2}}="{{$v2}}"{{end}}{{end}}` +
	`{{if $p.css}} style="{{range $k2, $v2 := .css}}{{$k2}}: {{$v2}}; {{end}}"{{end}}` +
	`{{if eq $p.value .Id}} checked{{end}}>{{end}}{{end}}` +
	` {{end}}`

var base_label = `{{define "label"}}` +
	`<div class="form-group">` +
	`<label class="control-label{{if .classes}}{{range .classes}} {{.}}{{end}}{{end}}"` +
	`{{if .forId}} for="{{.forId}}"{{end}}` +
	`>{{.label}}</label>` +
	`{{if .wrapper}}<div class="{{range .wrapper}}{{.}} {{end}}">{{end}}` +
	`{{range .fields}}{{ .Render }}{{end}}` +
	`{{if .wrapper}}</div>{{end}}` +
	`</div> {{ end }}`

var base_fieldset = `{{define "fieldset"}}` +
	`<fieldset{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .classes}} class="{{range $k,$v := .classes}} {{$k}}{{end}}"{{end}}` +
	`{{if .tags}}{{range $k,$v := .tags}} {{$k}}{{end}}{{end}}>
{{range .fields}}    {{ .Render }}
{{end}}  </fieldset> {{end}}`

var base_tabgroup = `{{define "tabgroup"}}` +
	`<ul class="nav nav-tabs">` +
	`{{range .panes}}` +
	`<li{{if eq .Index 0}} class="active"{{end}}>` +
	`<a data-toggle="tab" href="#{{.Name}}">{{.Head}}</a></li>` +
	`{{end}}</ul>` +
	`<div class="tab-content">` +
	`{{range .panes}}` +
	`<div id="{{.Name}}" class="tab-pane fade{{if eq .Index 0}} in active{{end}}">` +
	`{{range .Fields}}{{.Render}}{{end}}` +
	`</div>{{end}}` +
	`</div> {{end}}`

var base_panel = `{{define "panel"}}` +
	`<div` +
	`{{if .name}} name="{{.name}}"{{end}}` +
	`{{if .classes}} class="{{range $k,$v := .classes}} {{$k}}{{end}}"{{end}}` +
	`{{if .tags}}{{range $k,$v := .tags}} {{$k}}{{end}}{{end}}>` +
	`{{if .header}}<div>{{.header}}</div>{{end}}<div>
{{range .fields}}    {{.Render}}
{{end}}  </div>` +
	`{{if .footer}}<div>{{.footer.Render}}</div>{{end}}` +
	`</div> {{end}}`

var base_block = `{{define "block"}}` +
	`{{if .fields}}<div class="row"><div class="col-sm-12">` +
	`{{range .fields}}{{.Render}}{{end}}` +
	`</div></div>{{end}}` +
	`<div class="row"><div class="col-sm-6">` +
	`{{if .left}}{{range .left}}{{.Render}}{{end}}{{end}}` +
	`&nbsp;</div><div class="col-sm-6 text-right">&nbsp;` +
	`{{if .right}}{{range .right}}{{.Render}}{{end}}{{end}}` +
	`</div></div>` +
	` {{end}}`

var base_form = `{{define "form"}}` +
	`<form` +
	`{{if .name}} name="{{.name}}"{{end}}` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .classes}} class="{{range .classes}}{{.}} {{end}}"{{end}}` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .css}} style="{{range $k, $v := .css}}{{$k}}: {{$v}}; {{end}}"{{end}}` +
	` method="{{.method}}" action="{{.action}}">
{{ range .fields}}  {{ .Render}}
{{end}}</form> {{end}}`

var base_grid = `{{define "grid"}}` +
	`<table` +
	`{{if .name}} name="{{.name}}"{{end}}` +
	` class="table {{if .classes}}{{range .classes}}{{.}} {{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range $k,$v := .tags}} {{$k}}{{end}}{{end}}>` +
	`{{if .headers}}<thead><tr>{{range .headers}}<th>{{.}}</th>{{end}}</tr></thead>{{end}}` +
	`<tbody {{if .tbody}} id="{{.tbody}}"{{end}}>
{{if .rowset}}{{range .rowset}}  {{.Render}}
{{end}}{{end}}</tbody>` +
	`</table> {{end}}`


// Menu templates

var base_menuitem = `{{define "menuitem"}}` +
	`<li{{if .href}}><a href="{{.href}}">{{else}} active>{{end}}` +
	`{{if .glyph}}<span class="glyphicon {{.glyph}}"></span>&nbsp;{{end}}` +
	`{{.text}}{{if .href}}</a>{{end}}` +
	`</li> {{end}}`

var base_submenu = `{{define "submenu"}}` +
	`<li class="dropdown">` +
	`<a href="#" class="dropdown-toggle" data-toggle="dropdown">` +
	`{{if .glyph}}<span class="glyphicon {{.glyph}}"></span>&nbsp;{{end}}` +
	`{{.text}}<span class="caret"></span></a>` +
	`<ul class="dropdown-menu" role="menu">
{{range .items}}{{.Render}}
{{end}}</ul>` +
	`</li> {{end}}`

var base_topmenu = `{{define "topmenu"}}` +
	`<div class="navbar-header">` +
	`<button class="navbar-toggle" type="button" data-toggle="collapse" data-target=".bs-navbar-collapse">` +
	`<span class="icon-bar"></span>` +
	`<span class="icon-bar"></span>` +
	`<span class="icon-bar"></span>` +
	`</button><a href="{{.href}}">` +
	`{{if .icon}}<img id="logo" src="{{.icon}}"/>{{end}}` +
	`{{if .text}}<h3>{{.text}}</h3>{{end}}` +
	`</a></div>` +
	`<nav class="collapse navbar-collapse bs-navbar-collapse" role="navigation">` +
	`{{if .left}}<ul class="nav navbar-nav">{{range .left}}{{.Render}}
{{end}}</ul>{{end}}` +
	`{{if .right}}<ul class="nav navbar-nav navbar-right">{{range .right}}{{.Render}}
{{end}}</ul>{{end}}` +
	`</nav> {{end}}`

var base_breadcrumb = `{{define "breadcrumb"}}` +
	`<ul class="breadcrumb">` +
	`{{if .items}}{{range .items}}{{.Render}}{{end}}{{end}}` +
	`</ul> {{end}}`

