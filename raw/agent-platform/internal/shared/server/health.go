package server

import (
	"fmt"
	"log"
	"net/http"
)

func ServeHealth(addr, name string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		_, _ = fmt.Fprintf(w, "%s ok", name)
	})

	log.Printf("%s listening on %s", name, addr)
	return http.ListenAndServe(addr, mux)
}
