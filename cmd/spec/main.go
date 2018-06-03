package main

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	debug := false

	raw, err := ioutil.ReadFile("./spec.json")
	if err != nil {
		log.Fatal(err)
	}

	var spec SpecFile
	json.Unmarshal(raw, &spec)

	os.Mkdir("./spec", 0755)
	for _, val := range spec {
		filename := val.Entrypoint + ".json"
		var file OutFile
		contents, err := ioutil.ReadFile("./" + filename)
		if err != nil {
			file = val.toOutFile()
		} else {
			err = json.Unmarshal(contents, &file)
			if err != nil {
				log.Fatal("Error parsing ", filename, ": ", err)
			}
			val.applyToOutFile(&file)
		}
		raw, _ := json.MarshalIndent(file, "", "  ")
		ioutil.WriteFile("./spec/"+filename, raw, 0644)
		fmt.Println(filename)
	}

	if debug {
		spew.Dump(spec)
	}
}
