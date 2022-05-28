package main

import (
	"fmt"
	"log"
	"net/http"
	// "os"
)

// func main() {
// 	ln, err := net.Listen("tcp", ":2020")

// 	if err != nil {
// 		log.Panicln(err)
// 	}
// 	defer ln.Close()

// 	controller.HandleIncomingSockets(ln)
// }

func main() {
	http.HandleFunc("/", indexHandler)

	port := "8080"
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello, Docker!")
}
