package main

import (
	"fmt"
	"net/http"
	"os"
)

func home(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	fmt.Fprint(res, "The homepage.")
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
