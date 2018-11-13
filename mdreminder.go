package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

// wg ensures all goroutines finish their task before main function ends.
var wg sync.WaitGroup
// TODO:
var urlConn = make(chan string, 100)

// AccessUrl tests each url. If url can access, return code 200 with no error,
// Otherwise return the error code (like 404) and error.
func AccessUrl(url string) error {

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Url parse error: ", err)
		return err
	}

	response, err := client.Do(request)
	statusCode := response.StatusCode
	if statusCode != 200 {
		fmt.Println("Error occurs when execute \"NewRequest()\" on url, code is ", statusCode, url)
		return err
	}
	fmt.Printf("Url %s is available.\n", url)
	return nil
}

// ScanDir scans the project folder to find the file (mainly *.md) which contains url
// and store the file into []string, then using goroutine to open the file separately
// to search the urls, all these urls will return finally.
func ScanDir(dir string) {

	fList, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("Error occurs when read the directory", err)
	}

	// string of "/"
	pathSep := string(os.PathSeparator)

	for _, f := range fList {
		// joint path name
		subDirOrFile := dir + pathSep + f.Name()
		if f.IsDir() {
			ScanDir(subDirOrFile)
		}

		// ignore the letter case to judge the concrete type
		if strings.ToLower(path.Ext(f.Name())) == ".md" {
			//TODO: open other goroutine to search url in file
			//TODO: sync data from goroutine before main function ends
			// Use another goroutine to open file and search the links,
			//
			wg.Add(1)
			go func(filename string) {
				defer wg.Done()

				file, err := os.Open(filename)
				if err != nil {
					log.Fatal("Error occurs when open the .md file", err)
				}
				defer file.Close()

				res, err := ioutil.ReadFile(filename)
				if err != nil {
					log.Fatal("Error occurs when read file by ReadFile:", err)
				}
				// TODO: prints the file content for test
				re := regexp.MustCompile(`(\w+):\/\/([^\/:]+)\/([^\s\)]*)*`)
				set := re.FindAllString(string(res), -1)
				for _, url := range set {
					fmt.Println("Url is: ====================", url)
					urlConn <- url
					// TODO: test code should be removed later...
					fmt.Println("urlConn <- url has been finished............")
				}

				// TODO: parse the content and extract the links
				url := "https://github.com/AaronWharton"
				urlConn <- url
				fmt.Println("urlConn <- url has been finished............")

			}(subDirOrFile)

			wg.Wait()
			// TODO: test code should be removed later...
			fmt.Println("<-urlConn has been finished............")
			AccessUrl(<-urlConn)
		}
	}
}

func main() {
	// test
	ScanDir("/Users/aaron/Go/src/go-pro/src")
}
