package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", logMessage)
	log.Fatal(http.ListenAndServe("localhost:4444", nil))
}

func logMessage(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		panic(err)
	}
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(":", m["log"][0])
}
