package mdreminder

import (
	"errors"
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

var urlConn = make(chan string, 100)

// AccessUrl tests each url. If url can access, return code 200 with no error,
// Otherwise return the error code (like 404) and error.
func AccessUrl(url string) error {
	fmt.Println("This url is :", url)

	if url == "" {
		log.Println("This file does not contain any url!")
		return errors.New("this file does not contain any url")
	}
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
				re := regexp.MustCompile(`(\w+):\/\/([^\/:]+)\/([^\s\)]*)*`)
				set := re.FindAllString(string(res), -1)
				if len(set) == 0 {
					urlConn <- ""
				}

				for _, url := range set {
					urlConn <- url
				}

			}(subDirOrFile)

			AccessUrl(<-urlConn)
			wg.Wait()
		}
	}
}
