package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world again!\n")
	resp, err := http.Get("http://web.cloud66.local:3001/counts")
	if err != nil {
		// handle error
		io.WriteString(w, "Count from web.cloud66.local is: <<unavailable>>")
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		io.WriteString(w, "Count from web.cloud66.local is: <<unavailable>>")
		return
	}
	io.WriteString(w, "Count from web.cloud66.local is: "+string(body))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8000", nil)
}
