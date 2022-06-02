package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func GreetHandler(rw http.ResponseWriter, r *http.Request) {
	Greet(rw, "Internet")
}

func main() {
	http.ListenAndServe(":8000", http.HandlerFunc(GreetHandler))
	Greet(os.Stdout, "Console")
}
