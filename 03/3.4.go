// GETメソッドの送信
// $ curl http://localhost:18888

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	log.Println("status: ", res.Status)
	log.Println("status: ", res.StatusCode)
	log.Println("status: ", res.Header)
}
