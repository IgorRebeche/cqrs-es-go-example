package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/IgorRebeche/cqrs-es-go-example/internal/gateways/gatewayExample"
	"github.com/IgorRebeche/cqrs-es-go-example/internal/greet"
	"github.com/IgorRebeche/cqrs-es-go-example/internal/xau"
)

// func (w command.authorizeCommand)

func main() {

	fmt.Println(greet.TheWorld())
	fmt.Println(xau.Xau())
	fmt.Println(gatewayExample.Pega())

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
