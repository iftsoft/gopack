package srv

import (
	"net/http"
	"io/ioutil"
	"encoding/base64"
)

type Source int8
const (
	Source_plain Source = iota
	Source_base64
	Source_file
)

const (
	Content_icon	= "image/x-icon"
	Content_png 	= "image/png"
	Content_css		= "text/css; charset=utf-8"
	Content_js		= "application/javascript; charset=utf-8"
)

type staticFile struct {
	byteDump	[]byte
	content 	string
}

var staticStack = map[string]*staticFile {}

func AddStaticDump(url string, cont string, data []byte) (err error) {
	file := &staticFile{ data, cont  }
	if file.byteDump != nil {
		staticStack[url] = file
	}
	return err
}

func AddStaticFile(url string, cont string, src Source, data string) (err error) {
	file := &staticFile{ nil, cont  }
	switch src {
	case Source_plain:
		file.byteDump = []byte(data)
	case Source_base64:
		file.byteDump, err = base64.StdEncoding.DecodeString(data)
	case Source_file:
		file.byteDump, err = ioutil.ReadFile(data)
	}
	if err == nil && file.byteDump != nil {
		staticStack[url] = file
	}
	return err
}

func WriteStaticFile(w http.ResponseWriter, url string) (err error) {
	file, ok := staticStack[url];
	if !ok {
		http.Error(w, "Not found", http.StatusNotFound)
		return err
	}
	w.Header().Set("Content-Type", file.content)
//	w.Header().Set("Content-Length", string(len(file.byteDump)))
//	w.Header().Set("Cache-Control", "public, max-age=7776000")
	w.Header().Set("Cache-Control", "no-cache")
	if _, err := w.Write(file.byteDump); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return err
}

