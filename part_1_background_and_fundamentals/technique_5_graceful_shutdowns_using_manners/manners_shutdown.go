package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/braintree/manners"
)

type handler struct{}

func (*handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "Inigo Montoya"
	}
	fmt.Fprintf(res, "Hello, my name is %s!", name)
}

func shutDown() {
	shutDown := make(chan os.Signal)
	signal.Notify(shutDown, os.Interrupt, os.Kill)
	<-shutDown
	fmt.Println("Shutting down.")
	time.Sleep(time.Second)
	manners.Close()
}

func main() {
	go shutDown()
	fmt.Println("Serving at 8080.")
	manners.ListenAndServe(":8080", &handler{})
}
