// x-wwww-form-urlencodeed形式のPOSTメソッドの送信
// $ curl -d test=value http://localhost:18888

package main

import (
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"test": {"Value"},
	}
	res, err := http.PostForm("http://localhost:18888", values)
	if err != nil {
		panic(err)
	}
	log.Println("Status: ", res.Status)
}
