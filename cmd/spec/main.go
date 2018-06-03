package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
)

func main() {
	raw, err := ioutil.ReadFile("./_spec.json")
	if err != nil {
		log.Fatal(err)
	}

	var spec SpecFile
	json.Unmarshal(raw, &spec)

	for _, val := range spec {
		filename := val.Entrypoint + ".json"
		var file OutFile
		contents, err := ioutil.ReadFile("./" + filename)
		if err != nil {
			file = val.toOutFile()
		} else {
			json.Unmarshal(contents, &file)
			val.applyToOutFile(&file)
		}
		raw, _ := json.MarshalIndent(file, "", "  ")
		ioutil.WriteFile("./"+filename, raw, 0644)
	}

	spew.Dump(spec)
}
