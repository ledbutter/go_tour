package main

import (
	"net/http"
	"fmt"
)

type String string

type Struct struct {
	Greeting string
	Punct    string
	Who	 string
}

func (s String) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello String: %s\n", string(s))
}

func (s *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s%s%s\n", s.Greeting, s.Punct, s.Who)
}

func main() {
	http.Handle("/string", String("Hello World!"))
	http.Handle("/struct", &Struct{"Hello", ":", "World!"})
	http.ListenAndServe("localhost:4000", nil)
}

