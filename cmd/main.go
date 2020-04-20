package main

import (
	"week3"
	"week3/exporter"
	"week3/facebook"
	"week3/linkedIn"
	"week3/twitter"
)

func main() {
	// define the structs

	/*
	* important: I noticed that it would not allow the struct to work
	* if both(all) methods in the interface are not used
	 */

	fb := new(facebook.Facebook)
	twtr := new(twitter.Twitter)
	lnkd := new(linkedIn.LinkedIn)

	week3.ScrollFeeds(fb, twtr, lnkd)
	//export TXT files
	err := exporter.ExportTXT(fb, "fbdata")
	err = exporter.ExportTXT(twtr, "twtrdata")
	err = exporter.ExportTXT(lnkd, "lnkddata")
	// export JSON files
	err = exporter.ExportJSON(fb, "fbdata")
	err = exporter.ExportJSON(twtr, "twtrdata")
	err = exporter.ExportJSON(lnkd, "lnkddata")
	// export XML files
	err = exporter.ExportXML(fb, "fbdata")
	err = exporter.ExportXML(twtr, "twtrdata")
	err = exporter.ExportXML(lnkd, "lnkddata")
	if err != nil {
		panic(err)
	}

}
