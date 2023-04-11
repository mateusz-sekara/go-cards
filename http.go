package main

import (
	"fmt"
	"net/http"
	"os"
)

func httpSandbox() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Connection error", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	//_, err = resp.Body.Read(bs)
	//if err != nil {
	//	fmt.Println("Reading error", err)
	//	os.Exit(1)
	//}
	fmt.Println(string(bs))
}
