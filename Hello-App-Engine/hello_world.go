package helloworld

import (
	"net/http"
	"fmt"
)

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test/", testHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "handler() from go-playing/helloWorld/ # from github")
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "testHandler() from go-playing/helloWorld/test/")
}
