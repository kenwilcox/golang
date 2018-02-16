package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
)

type Lang struct {
	Name string
	Year int
	URL  string
}

func main() {
	lang := Lang{"Go", 2009, "http://golang.org/"}
	//fmt.Printf("%v\n", lang)
	//fmt.Printf("%+v\n", lang)
	//fmt.Printf("%#v\n", lang)
	json, err := json.Marshal(lang)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", json)

	xml, err := xml.MarshalIndent(lang, "", "  ")
	//xml, err := xml.Marshal(lang)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", xml)
}
