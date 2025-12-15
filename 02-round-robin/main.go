package main

import (
	"fmt"
	"log"
	"net/http"
	"sync" 
	"github.com/ArminAhmadkhaniha/go-load-balancer/proxy"
)

func main() {
	servers := []string {
		"https://www.google.com",
		"https://www.yahoo.com",
		"https://www.bing.com",

	}
	var proxies []http.Handler
	for _, serversUrl := range servers {
		handler, err := proxy.NewProxy(serversUrl)
		if err != nil {
			log.Fatal(err)

		}
		proxies = append(proxies, handler)

	}
	var counter int = 0
	var mu sync.Mutex
	lbhandeler := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		counterIndex := counter % len(proxies)
		counter++
		targetProxy := proxies[counterIndex]
		targetUrl := servers[counterIndex]
		targetProxy.ServeHTTP(w, r)
		fmt.Printf("Search Request %d -> Forwarding to: %s\n", counter, targetUrl)
	}
	http.HandleFunc("/", lbhandeler)
	fmt.Println("Load Balancer running at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))





}
