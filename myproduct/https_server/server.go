// package main
//
// import (
// 	"fmt"
// 	"log"
// 	"net/http"
//
// 	"golang.org/x/net/http2"
// )
//
// func main() {
// 	var s http.Server
// 	http2.VerboseLogs = true
// 	s.Addr = ":18443"
//
// 	http2.ConfigureServer(&s, nil)
//
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "text/plain")
// 		fmt.Fprintf(w, "Hello World")
// 	})
//
// 	log.Fatal(s.ListenAndServeTLS("../../keys/server.crt", "../../keys/server.key"))
// }

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

var image []byte

func init() {
	var err error
	image, err = ioutil.ReadFile("./image.png")
	if err != nil {
		panic(err)
	}
}

// func handlerHtml(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	fmt.Fprintf(w, `
// 		<html>
// 			<body>
// 				<img src="/image"/>
// 			</body>
// 		</html>`)
// }

func handlerHtml(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t := template.Must(template.ParseFiles("index.html.tpl"))
	if err := t.ExecuteTemplate(w, "index.html.tpl", []int{1, 2, 3}); err != nil {
		log.Fatal(err)
	}
}

func handlerImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Write(image)
}

func main() {
	var s http.Server
	http2.VerboseLogs = true
	s.Addr = ":18443"

	http2.ConfigureServer(&s, nil)
	http.HandleFunc("/", handlerHtml)
	http.HandleFunc("/image", handlerImage)

	fmt.Println("start http listening :18443")
	log.Fatal(s.ListenAndServeTLS("../../keys/server.crt", "../../keys/server.key"))
}

// func main() {
// 	http.HandleFunc("/", handlerHtml)
// 	fmt.Println("start http listening :18443")
// 	err := http.ListenAndServeTLS(":18443", "../../keys/server.crt", "../../keys/server.key", nil)
// 	fmt.Println(err)
// }
