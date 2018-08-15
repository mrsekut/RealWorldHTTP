// 自由なメソッドの送信
// $ curl -X DELETE http://localhost:18888

package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("DLETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}

	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
