// HTTPプロトコルバージョンを確認
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// resp, err := http.Get("http://google.com")  // HTTP/1.1
	resp, err := http.Get("https://google.com") // HTTP/2.0
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("Protocol Version: %s\n", resp.Proto)
}
