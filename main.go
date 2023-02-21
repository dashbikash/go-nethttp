package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fmt.Printf("Rest API\n")
	mux.HandleFunc("/", handlerIndex)

	mux.Handle("/mw", middleware(http.HandlerFunc(handlerMiddleware)))
	log.Fatal(http.ListenAndServe(":8090", mux))
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func handlerMiddleware(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is logged page")
}

func middleware(next http.Handler) http.Handler {

	nh := func(w http.ResponseWriter, r *http.Request) {
		log.Println("Inside middleware")
		// Code that run before
		next.ServeHTTP(w, r) // Breaks the flow if not used
		// Code that run after
	}

	return http.HandlerFunc(nh)
}
