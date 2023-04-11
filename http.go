package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func httpSandbox() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Connection error", err)
		os.Exit(1)
	}

	_, err = io.Copy(logWriter{}, resp.Body)
	if err != nil {
		fmt.Println("Copy error", err)
		os.Exit(1)
	}
}
