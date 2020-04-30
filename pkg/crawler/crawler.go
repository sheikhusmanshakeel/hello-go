package crawler

import "fmt"

/*
This function will crawl. it won't run, or walk. it will crawl!
*/
func Crawl() {
	fmt.Println("This is the crawler")
}

// This function on the other hand will strip!
func StripHTML(a string){
	fmt.Println(fmt.Sprintf("This is the parameter passed %s",a))
}