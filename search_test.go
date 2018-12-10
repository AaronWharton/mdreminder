package mdreminder

import (
	"fmt"
	"testing"
)

func TestScanDir(t *testing.T) {
	if _, err := ScanDir("???"); err == nil {
		t.Errorf("Test for ScanDir failed! Directory path is: ???\n")
	} else {
		fmt.Printf("Test for ScanDir succeed! Directory path is: ???\n")
	}

	// access root directory `/`, it may take a long time.
	// Maybe you can choose other directory to test.
	if _, err := ScanDir("/"); err != nil {
		t.Errorf("Test for ScanDir failed! Directory path is: /\n")
	} else {
		fmt.Printf("Test for ScanDir succeed! Directory path is: /\n")
	}
}
