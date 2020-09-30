package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var svar string
	flag.StringVar(&svar, "url", "https://google.com", "a url to check needs to be inputed here")
	flag.Parse()
	fmt.Println("Welcome to the NVIDIA API Checker v0.4 by redstone2019.")
	fmt.Println("This will loop forever giving HTTP Status Codes.")
	fmt.Println("Please wait. (Usually takes aroung 10-30 seconds per loop)")
	fmt.Println("URL:", svar)
	for {
		resp, err := http.Get(svar)
		now := time.Now()
		// Print the HTTP Status Code and Status Name
		fmt.Println(err)
		fmt.Println(now)
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		time.Sleep(10 * time.Second)
	}
}
