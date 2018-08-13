// クエリの送信
// $ curl -G --data-urlencode "query=hello world" http://localhost:18888

package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"query": {"hello world"},
	}
	res, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	log.Println(string(body))
}
