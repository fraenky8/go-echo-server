package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	defaultPort = "80"
)

func main() {
	log.Println("starting up ...")

	port := defaultPort
	if len(os.Args) > 1 {
		port = strings.TrimSpace(os.Args[1])
	}

	if port == "" {
		log.Fatalf("invalid port specified: %v\n", os.Args[1])
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("could not get hostname: %v\n", err)
	}

	log.Println("register handler /")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("<%v> [%v]\t%v", hostname, r.Method, r.URL)

		contentType := r.Header.Get("Content-Type")
		if contentType == "" {
			contentType = "text/plain"
		}
		w.Header().Set("Content-Type", contentType)

		defer r.Body.Close()
		io.Copy(w, r.Body)
	})

	log.Println("register handler /hostname")
	http.HandleFunc("/hostname", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("<%v> [%v]\t%v", hostname, r.Method, r.URL)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(hostname))
		w.Write([]byte("\n"))
	})

	log.Println("register handler /formvalues")
	http.HandleFunc("/formvalues", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("<%v> [%v]\t%v", hostname, r.Method, r.URL)
		w.Header().Set("Content-Type", "text/plain")

		err := r.ParseForm()
		if err != nil {
			w.Write([]byte(err.Error()))
			w.Write([]byte("\n"))
			return
		}

		for key, values := range r.Form { // range over map
			w.Write([]byte(fmt.Sprintf("[%s] => %v\n", key, values)))
		}
	})

	log.Printf("<%v> listening at :%v ...\n", hostname, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
