package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start server")
	http.HandleFunc("/get", Handle)
	err := http.ListenAndServe(":12345", new(MyHandler))
	if err != nil {
		fmt.Println(err)
	}
}

func Handle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req method:", req.Method, req.URL.Path, req.URL.Query())
	q := req.URL.Query()
	for k, v := range q {
		fmt.Println("query param", k, v)
	}
	for k, v := range req.Header {
		fmt.Fprintf(w, "%q=%v", k, v)
	}
}

type MyHandler struct {
}

func (m MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Println(req.URL.Path)
		fmt.Fprintln(w, req.URL.Path)
	case "/name":
		fmt.Println("hello :", req.URL.Query()["name"], req.URL.Query()["phone"], req.Form)
		fmt.Fprintln(w, "hello :", req.URL.Query()["name"], req.URL.Query()["phone"])
	default:
		fmt.Println("404")
		fmt.Fprintln(w, "404")
	}

}
