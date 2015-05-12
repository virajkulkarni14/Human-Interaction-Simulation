package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	logToServer("this will be logged :D")
}

func logToServer(s string) {
	encodedString, err := UrlEncoded(s)
	fmt.Println("url encoded:", encodedString)
	url := "http://localhost:4444?log=" + encodedString
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
