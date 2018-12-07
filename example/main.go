package main

import (
	"fmt"
	"log"
	"mdreminder"
)

func main() {
	s := mdreminder.ScanDir("/Users/aaron/go/src/mdreminder")
	fmt.Println("Found ", len(s), " link(s) in this directory.")
	for _, v := range s {
		if err := mdreminder.AccessUrl(v); err != nil {
			log.Println(err)
		}
	}
}
