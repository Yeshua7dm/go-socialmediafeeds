package exporter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"week3"
)

func ExportTXT(u week3.SocialMedia, filename string) error {
	f, err := os.OpenFile(filename+".txt", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	for _, feed := range u.Feed() {
		n, err := f.Write([]byte(feed + "\n"))
		if err != nil {
			return err
		}
		fmt.Printf("wrote %d bytes\n", n)
	}
	return nil
}

func ExportJSON(u week3.SocialMedia, filename string) error {
	// give the file the indexes
	jsonMap := make(map[int]interface{})
	for i := 0; i < len(u.Feed()); i++ {
		jsonMap[i+1] = u.Feed()[i]
	}
	// create the json file
	file, err := json.MarshalIndent(jsonMap, "", " ")
	err = ioutil.WriteFile(filename+".json", file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ExportXML(u week3.SocialMedia, filename string) error {
	file, err := xml.MarshalIndent(u.Feed(), "", " ")
	err = ioutil.WriteFile(filename+".xml", file, 0644)
	if err != nil {
		return err
	}

	return nil
}
