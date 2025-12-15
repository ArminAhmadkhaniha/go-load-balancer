package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/ArminAhmadkhaniha/go-load-balancer/proxy"
)

func main() {
	target := "https://www.google.com"
	handle , err := proxy.NewProxy(target)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", handle)
	fmt.Printf("Load Balancer started at :8000 -> forwarding to %s\n", target)
	log.Fatal(http.ListenAndServe(":8000", nil))
}


