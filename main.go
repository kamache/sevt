package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>This is a house of learned doctors.</h1>")
	io.WriteString(w, "<a href=\"https://www.youtube.com/watch?v=hh1oaumUoyc\">https://www.youtube.com/watch?v=hh1oaumUoyc</a>")
	fmt.Fprintf(w, "<h2>Hi there, I love %s!</h2>", r.URL.Path[1:])
}

func main() {
	fmt.Printf("Starting http server...\n")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err.Error())
}
