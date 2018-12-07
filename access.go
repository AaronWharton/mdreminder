package mdreminder

import (
	"fmt"
	"log"
	"net/http"
)

// AccessUrl tests each url. If url can access, return code 200 with no error,
// Otherwise return the error code (like 404) and error.
func AccessUrl(url string) error {
	fmt.Println("This url is :", url)

	if url == "" {
		// TODO: to make hint more friendly
		return nil
	}
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Url parse error: ", err)
		return err
	}

	response, err := client.Do(request)
	// handle err like: `net/http: TLS handshake timeout` etc.
	if err != nil {
		log.Println("Error occurred when execute `client.Do()` : ", err)
		return err
	}
	if response.StatusCode != 200 {
		fmt.Println("Error occurs when execute \"NewRequest()\" on url, code is ", response.StatusCode, url)
		return err
	}
	fmt.Printf("Url %s is available.\n", url)
	return nil
}
