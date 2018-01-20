package templ

var boot_button = `{{define "button"}}` +
	`<button type="{{.type}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="btn {{if .classes}}{{range .classes}} {{.}}{{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>{{.text}}</button> {{end}}`

var boot_static = `{{define "static"}}` +
	`<p name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="form-control-static {{if .classes}}{{range .classes}}{{.}} {{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>{{.value}}</p>` +
	`{{end}}`

var boot_search = `{{define "search"}}` +
	`<div class="input-group">` +
	`<input type="{{.type}}" name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="form-control{{if .inpClas}} {{range .inpClas}}{{.}} {{end}}{{end}}"` +
	`{{if .inpPars}}{{range $k, $v := .inpPars}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .inpTags}}{{range .inpTags}} {{.}}{{end}}{{end}}` +
	`{{if .value}} value="{{.value}}"{{end}}>` +
	`<div class="input-group-btn">` +
	`<button class="btn{{if .butClas}} {{range .butClas}}{{.}} {{end}}{{end}}"` +
	`{{if .butPars}}{{range $k, $v := .butPars}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .butTags}}{{range .butTags}} {{.}}{{end}}{{end}}>` +
//	` type="button">` +
	`<i class="glyphicon glyphicon-search"></i></button>` +
	`</div></div>` +
	`{{end}}`

var boot_input = `{{define "input"}}` +
	`<input type="{{.type}}" name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="form-control{{if .classes}} {{range .classes}}{{.}} {{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`{{if .value}} value="{{.value}}"{{end}}>` +
	`{{if .helpText}}<span class="help-block">{{.helpText }}</span>{{end}}` +
	`{{ end }}`

var boot_textarea = `{{define "textarea"}}` +
	`<textarea name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="form-control{{ if .classes }} {{range .classes}}{{.}} {{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>{{.value}}</textarea>` +
	`{{if .helpText}}<span class="help-block">{{.helpText }}</span>{{end}}` +
	`{{end}}`

var boot_enum = `{{define "enum"}}` +
	`<select name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="form-control{{if .classes}}{{range .classes}} {{.}}{{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>{{$p := . }}` +
	`{{range .choices}}<option value="{{.Id}}">{{.Val}}</option>{{end}}` +
	`</select>` +
	`{{if .helpText}}<span class="help-block">{{.helpText }}</span>{{end}}` +
	`{{end}}`

var boot_select = `{{define "select"}}` +
	`<select name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="form-control{{if .classes}}{{range .classes}} {{.}}{{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>{{$p := . }}` +
	`{{range .choices}}<option value="{{.Id}}">{{.Val}}</option>{{end}}` +
	//`{{range $k, $v := .choices}}{{if $k }}<optgroup label="{{$k}}">{{end}}{{range $v}}` +
	//`<option value="{{.Id}}"{{if $p.tags.multiple }}{{$id := .Id}}` +
	//`{{range $k2, $p2 := $p.multValues}}{{if eq $k2 $id}} selected{{end}}{{end}}` +
	//`{{else}}{{ if eq $p.value .Id}} selected{{end}}{{end}}>` +
	//`{{.Val}}</option>{{end}}` +
	//`{{if $k}}</optgroup>{{end}}{{end}}` +
	`</select>` +
	`{{if .helpText}}<span class="help-block">{{.helpText }}</span>{{end}}` +
	`{{end}}`

var boot_checkbox = `{{define "checkbox"}}` +
	`<div class="checkbox form-control"` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}>` +
	`<label class="form-check-label">` +
	`<input type="checkbox" name="{{.name}}"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="form-check-input {{if .classes}}{{range .classes}}{{.}} {{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>{{.value}}</label>` +
	`{{if .helpText}}<span class="help-block">{{.helpText }}</span>{{end}}` +
	`</div> {{end}}`

var boot_bitmask = `{{define "bitmask"}}{{$p := . }}` +
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

var boot_radio = `{{define "radio"}}{{$p := . }}` +
	`<div class="radio form-control" style="float:left; height:auto;"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}>` +
	`{{ range .choices }}` +
		`<label class="col-xs-12 col-sm-6 col-md-4 col-lg-3">` +
		`<input type="radio" name="{{$p.name}}"` +
		`{{if $p.classes}} class="{{range $p.classes}}{{.}} {{end}}"{{end}}` +
		`{{if $p.params}}{{range $k2, $v2 := $p.params}} {{$k2}}="{{$v2}}"{{end}}{{end}}` +
		`{{if $p.tags}}{{range $p.tags}} {{.}}{{end}}{{end}}` +
		` value="{{.Id}}">{{.Val}}</label>` +
	`{{end}}` +
	`{{if .helpText}}<span class="help-block">{{.helpText }}</span>{{end}}` +
	`</div> {{end}}`

var boot_label = `{{define "label"}}` +
	`<div class="form-group">` +
	`<label class="control-label{{if .classes}}{{range .classes}} {{.}}{{end}}{{end}}"` +
	`{{if .forId}} for="{{.forId}}"{{end}}` +
	`>{{.label}}</label>` +
	`{{if .wrapper}}<div class="{{range .wrapper}}{{.}} {{end}}">{{end}}` +
	`{{range .fields}}{{ .Render }}{{end}}` +
	`{{if .wrapper}}</div>{{end}}` +
	`</div> {{ end }}`

var boot_fieldset = `{{define "fieldset"}}` +
	`<fieldset{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .classes}} class="{{range $k,$v := .classes}} {{$k}}{{end}}"{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>{{range .fields}}    {{ .Render }}
{{end}}  </fieldset> {{end}}`

var boot_tabgroup = `{{define "tabgroup"}}` +
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

var boot_panel = `{{define "panel"}}` +
	`<div{{if .id}} id="{{.id}}"{{end}}` +
	` class="panel {{if .classes}}{{range .classes}}{{.}} {{end}}{{end}}"` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>{{if .header}}<div class="panel-heading">{{.header}}</div>{{end}}
<div class="panel-body">
{{range .fields}}    {{.Render}}
{{end}}  </div>` +
	`{{if .footer}}<div class="panel-footer">{{.footer.Render}}</div>{{end}}` +
	`</div> {{end}}`

var boot_block = `{{define "block"}}` +
	`{{if .fields}}<div class="row"><div class="col-xs-12">` +
	`{{range .fields}}{{.Render}}{{end}}` +
	`</div></div>{{end}}` +
	`<div class="row"><div class="col-xs-6">` +
	`{{if .left}}{{range .left}}{{.Render}}{{end}}{{end}}` +
	`</div><div class="col-xs-6 text-right">` +
	`{{if .right}}{{range .right}}{{.Render}}{{end}}{{end}}` +
	`</div></div>` +
	` {{end}}`

var boot_form = `{{define "form"}}` +
	`<form role="form"` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	`{{if .classes}} class="{{range .classes}}{{.}} {{end}}"{{end}}` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}` +
	`>
{{range .fields}}  {{.Render}}
{{end}}</form> {{end}}`

var boot_grid = `{{define "grid"}}` +
	`<table` +
	`{{if .id}} id="{{.id}}"{{end}}` +
	` class="table {{if .classes}}{{range .classes}}{{.}} {{end}}{{end}}"` +
	`{{if .params}}{{range $k, $v := .params}} {{$k}}="{{$v}}"{{end}}{{end}}` +
	`{{if .tags}}{{range .tags}} {{.}}{{end}}{{end}}>` +
	`{{if .headers}}<thead><tr>{{range .headers}}<th>{{.}}</th>{{end}}</tr></thead>{{end}}` +
	`<tbody{{if .tbody}} id="{{.tbody}}"{{end}}></tbody>` +
	`</table>` +
	` {{end}}`


// Menu templates

var boot_menuitem = `{{define "menuitem"}}` +
	`<li{{if .href}}><a href="{{.href}}">{{else}} active>{{end}}` +
	`{{if .glyph}}<span class="glyphicon {{.glyph}}"></span>&nbsp;{{end}}` +
	`{{.text}}{{if .href}}</a>{{end}}` +
	`</li> {{end}}`

var boot_submenu = `{{define "submenu"}}` +
	`<li class="dropdown">` +
	`<a href="#" class="dropdown-toggle" data-toggle="dropdown">` +
	`{{if .glyph}}<span class="glyphicon {{.glyph}}"></span>&nbsp;{{end}}` +
	`{{.text}}<span class="caret"></span></a>` +
	`<ul class="dropdown-menu" role="menu">
{{range .items}}{{.Render}}
{{end}}</ul>` +
	`</li> {{end}}`

var boot_topmenu = `{{define "topmenu"}}` +
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


var boot_breadcrumb = `{{define "breadcrumb"}}` +
	`<ul class="breadcrumb" style="margin-bottom:0px;">` +
	`{{if .items}}{{range .items}}{{.Render}}{{end}}{{end}}` +
	`</ul> {{end}}`

//<div class="container"></div>