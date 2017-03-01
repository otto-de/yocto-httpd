package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func logRequest(req *http.Request, statusCode int) {
	log.Printf("%s %s %s %d", req.RemoteAddr, req.Method, req.URL.Path, statusCode)
}

func serveRequest(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		logRequest(req, http.StatusMethodNotAllowed)
		w.Header().Add("Allow", "GET")
		http.Error(w, "Only GET supported.", http.StatusMethodNotAllowed)
		return
	}

	logRequest(req, http.StatusOK)
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "OK")
}

func listen(port uint) error {
	log.Printf("starting httpd on port %d", port)
	addr := fmt.Sprintf(":%d", port)
	http.Handle("/", http.HandlerFunc(serveRequest))
	if err := http.ListenAndServe(addr, nil); err != nil {
		return err
	}
	log.Printf("stopped httpd")
	return nil
}

func main() {
	var port = flag.Uint("port", 8080, "Port to listen for connections")
	flag.Parse()

	err := listen(*port)
	if err != nil {
		log.Printf("error starting httpd: %s", err)
		os.Exit(1)
	}
}
