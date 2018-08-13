// 任意の文字列をPOST送信
// $ curl -H "Content-Type: text/plain" -H "Content-Length"  -d "テキスト"  http://localhost:18888

package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("テキスト")
	res, err := http.Post("http://localhost:18888", "test/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res.Status)
}
