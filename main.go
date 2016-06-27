package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Fetching data from another service!\n")
	var sourceURL = os.Getenv("DATA_SOURCE")
	if sourceURL == "" {
		sourceURL = "http://web.cloud66.local:3000/counts"
	}
	resp, err := http.Get(sourceURL)
	if err != nil {
		// handle error
		io.WriteString(w, fmt.Sprintf("Count from %s is \"<<unavailable>>\"", sourceURL))
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		io.WriteString(w, fmt.Sprintf("Count from %s is \"<<unavailable>>\"", sourceURL))
		return
	}
	io.WriteString(w, fmt.Sprintf("Count from %s is \"%s\"", sourceURL, string(body)))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
