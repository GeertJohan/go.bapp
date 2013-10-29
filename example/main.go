package main

import (
	"fmt"
	"github.com/GeertJohan/go.bapp"
	"log"
	"net/http"
)

func main() {
	// create new bapp instance
	b, err := bapp.NewBapp()
	if err != nil {
		log.Fatal(err)
	}

	// set handlerfunc on the http.DefaultServeMux, just like you would with a normal http application
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello bapp")
	})

	// open the application in browser
	b.Open()

	select {}
}
