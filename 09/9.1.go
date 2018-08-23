package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Printf("Protocol Version: $s\n", resp.Proto)
}
