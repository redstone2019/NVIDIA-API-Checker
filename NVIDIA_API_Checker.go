package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Welcome to the NVIDIA API Checker v0.2 by redstone2019.")
	fmt.Println("This will loop forever giving HTTP Status Codes.")
	fmt.Println("Please wait. (Usually takes aroung 10-30 seconds per loop)")
	for {
		resp, err := http.Get("https://api-prod.nvidia.com/direct-sales-shop/DR/products/en_us/USD/5379432500")
		// Print the HTTP Status Code and Status Name
		now := time.Now()
		fmt.Println(err)
		fmt.Println(now)
		fmt.Println("HTTP Response Status:", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
}
