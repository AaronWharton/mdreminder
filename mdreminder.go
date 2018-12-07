package mdreminder

import (
	"fmt"
)

// ExecuteByPath scans directory you specified to find all urls in this directory,
// then return the urls test result in console.
// View search.go and access.go to find more details.
func ExecuteByPath(path string) {
	urlList := ScanDir(path)
	fmt.Println("Found ", len(urlList), " link(s) in this directory.")
	for _, v := range urlList {
		// ignore the errors cause AccessUrl() has output the error information
		_ = AccessUrl(v)
	}
}