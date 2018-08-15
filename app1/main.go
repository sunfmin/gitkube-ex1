package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/sunfmin/go-cicd-example1/app1/pkg1"
)

func main() {
	pkg1.Hello()

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		b, _ := httputil.DumpRequest(r, true)
		fmt.Println(string(b))
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":4000", nil))

}
