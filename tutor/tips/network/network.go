package network

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com/",
	"http://www.topgoer.com/",
	"https://blog.golang.org/",
}

func ArrPrint2() {
	// Execute an HTTP HEAD request for all url's
	// and returns the HTTP status string or an error string.
	for _, url := range urls {
		resp, err := http.Head(url)

		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}
