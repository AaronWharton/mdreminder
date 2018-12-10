package mdreminder

import (
	"fmt"
	"testing"
)

func TestAccessUrl(t *testing.T) {
	if err := AccessUrl("...."); err == nil {
		t.Errorf("Test for AccessUrl failed! Url is: ...\n")
	} else {
		fmt.Printf("Test for AccessUrl succeed! Url is: ...\n")
	}

	if err := AccessUrl("https://github.com"); err != nil {
		t.Errorf("Test for AccessUrl failed! Url is https://github.com\n ")
	} else {
		fmt.Printf("Test for AccessUrl succeed! Url is https://github.com\n")
	}
}
