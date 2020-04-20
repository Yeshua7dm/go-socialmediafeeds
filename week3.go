package week3

import "fmt"

// Always remember Eric, a struct is an object

// an interface is  collection of method signatures basically, that is a method
// and the xpected output
type SocialMedia interface { //social media functions
	Feed() []string //a method that returns an array of strings
	Fame() int      // a method that returns an integer value
}

//function that loops through feeds and display them
func ScrollFeeds(platforms ...SocialMedia) {
	// read as: for platform in platforms
	for _, platform := range platforms {
		for _, feed := range platform.Feed() {
			fmt.Println(feed)
		}
		fmt.Println("**************************************")
	}
}
