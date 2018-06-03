package main

import (
	"strings"
)

type SpecFile []*SpecEntry

type SpecEntry struct {
	Title       string
	Description string
	Package     string
	Entrypoint  string
	APIs        map[string]*SpecAPI
}

func (s *SpecEntry) toOutFile() OutFile {
	file := OutFile{
		APIs: map[string]*OutFileAPI{},
	}
	s.applyToOutFile(&file)
	return file
}

func (s *SpecEntry) applyToOutFile(o *OutFile) {
	// reset title/interface/path to spec data
	o.Title = s.Title
	o.Description = s.Description
	o.Package = s.Package
	o.Interface = strings.ToUpper(s.Entrypoint[0:1]) + s.Entrypoint[1:]
	o.Path = "/" + s.Entrypoint
	for key, val := range s.APIs {
		path := val.Path
		if path == "" {
			path = "/" + key
		}
		// add new API calls
		call, ok := o.APIs[key]
		if !ok {
			o.APIs[key] = &OutFileAPI{
				Method:     val.Method,
				Title:      val.Title,
				Path:       path,
				Parameters: map[string]interface{}{},
			}
		} else {
			// update title/method/path of existing APIs
			call.Title = val.Title
			call.Method = val.Method
			call.Path = path
			if val.Parameters != nil {
				call.Parameters = val.Parameters
			}
		}
	}
}

type SpecAPI struct {
	Method string
	Title  string
	Path   string
	Parameters map[string]interface{}
}

type OutFile struct {
	Title       string
	Description string `json:",omitempty"`
	Package     string
	Interface   string
	Path        string
	APIs        map[string]*OutFileAPI
}

type OutFileAPI struct {
	Method     string
	Title      string
	Description string `json:",omitempty"`
	Path       string
	Parameters map[string]interface{}
}
