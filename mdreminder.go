package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

// global variable to hold all the urls
var urls = make([]string, 0, 10)

// AccessUrl tests each url. If url can access, return code 200 with no error,
// Otherwise return the error code (like 404) and error.
func AccessUrl(url string) (code int, err error) {
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error occurs when execute \"NewRequest()\" on url:", err)
	}

	response, err := client.Do(request)
	statusCode := response.StatusCode
	if statusCode != 200 {
		return statusCode, err
	}

	return statusCode, nil
}

// ScanDir scans the project folder to find the file (mainly *.md) which contains url
// and store the file into []string, then using goroutine to open the file separately
// to search the urls, all these urls will return finally.
func ScanDir(dir string) []string {

	fList, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("Error occurs when read the directory", err)
	}

	for _, f := range fList {
		if f.IsDir() {
			ScanDir(f.Name())
		}

		// ignore the letter case to judge the file type
		if strings.ToLower(path.Ext(f.Name())) == ".md" {
			//TODO: open other goroutine to search url in file
			go func(filename string) {
				file, err := os.Open(filename)
				if err != nil {
					log.Fatal("Error occurs when open the .md file", err)
				}
				defer file.Close()

				// read all data from files using ioutil
				res, err := ioutil.ReadFile(filename)
				if err != nil {
					log.Fatal("Error occurs when read file by ReadFile:", err)
				}
				fmt.Println(res)
			}(f.Name())
		}
	}
	return urls
}

func main() {

}
