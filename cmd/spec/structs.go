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
	APIs        []*SpecAPI
}

func (s *SpecEntry) toOutFile() OutFile {
	file := OutFile{
		APIs: []*OutFileAPI{},
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

	namedAPIs := o.NamedAPIs()

	for _, val := range s.APIs {
		path := val.Path
		if path == "" {
			path = "/" + val.Name
		}
		// add new API calls
		call, ok := namedAPIs[val.Name]
		if !ok {
			o.APIs = append(o.APIs, &OutFileAPI{
				Name:       val.Name,
				Method:     val.Method,
				Title:      val.Title,
				Path:       path,
				Parameters: val.Parameters,
			})
		} else {
			// update title/method/path of existing APIs
			call.Name = val.Name
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
	Name       string
	Method     string
	Title      string
	Path       string
	Parameters map[string]interface{}
}

type OutFile struct {
	Title       string
	Description string `json:",omitempty"`
	Package     string
	Interface   string
	Path        string
	APIs        []*OutFileAPI
}

func (o *OutFile) NamedAPIs() map[string]*OutFileAPI {
	apis := map[string]*OutFileAPI{}
	for _, api := range o.APIs {
		apis[api.Name] = api
	}
	return apis
}

type OutFileAPI struct {
	Name        string
	Method      string
	Title       string
	Description string `json:",omitempty"`
	Path        string
	Parameters  map[string]interface{}
}
