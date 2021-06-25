package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
		fmt.Println(r.Proto)

		if pusher, ok := w.(http.Pusher); ok {
			// Push is supported.
			options := &http.PushOptions{
				Header: http.Header{
					"Accept-Encoding": r.Header["Accept-Encoding"],
				},
			}
			// Push is supported.
			if err := pusher.Push("/app.js", options); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
	})

	h2s := &http2.Server{
	}
	h1s := &http.Server{
		Addr:    ":8008",
		Handler: h2c.NewHandler(handler, h2s),
	}
	http2.ConfigureServer(h1s, h2s)
	log.Fatal(h1s.ListenAndServe())
}
