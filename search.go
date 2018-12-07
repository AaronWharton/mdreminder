package mdreminder

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
)

// wg ensures all goroutines finish their task before main function ends.
var wg sync.WaitGroup

// urlList stores urls as []string.
var urlList []string

// ScanDir scans the project folder to find the file (mainly *.md) which contains url
// and store the file into []string, then using goroutine to open the file separately
// to search the urls, all these urls will return finally.
func ScanDir(dir string) []string {

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

				for _, url := range set {
					urlList = append(urlList, url)
				}

			}(subDirOrFile)

			wg.Wait()
		}
	}
	return urlList
}
