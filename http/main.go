package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)
//
func main() {
	// the site http://www.google.com is passed as the first argument
	arg := os.Args[1]
	resp, err := http.Get(arg)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}


	if resp.StatusCode == 200 {
	//	bs := make([]byte, 99999)
	//	resp.Body.Read(bs)
	//	fmt.Println(string(bs))

		// This does what the above does
		io.Copy(os.Stdout, resp.Body)
	} else {
		fmt.Println("Site is DOWN")
	}
}
