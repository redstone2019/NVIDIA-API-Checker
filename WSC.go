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
	fmt.Println("Welcome to the NVIDIA API Checker v0.51 by redstone2019.")
	fmt.Println("This will loop forever giving HTTP Status Codes.")
	fmt.Println("Please wait. (Usually takes around 1 second per loop for fast websites, slow websites may take much longer)")
	fmt.Println("Warning you must include either https:// or http:// before your domain or you will encounter a panic runtime error")
	fmt.Println("URL:", svar)
	time.Sleep(5 * time.Second)
	for {
		fmt.Println("Pinging...")
		resp, err := http.Get(svar)
		now := time.Now()
		// Print the HTTP Status Code and Status Name
		fmt.Println(now)
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
		fmt.Println("Error?", err)
		fmt.Println(" ")
		time.Sleep(1 * time.Second)
	}
}
