package main

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	statusStr := vars["statusCode"]

	logger.Printf("Status handler called with staus code: %s", statusStr)

	statusCode, err := strconv.Atoi(statusStr)
	if err != nil {
		http.Error(w, "Failed to convert to integer", http.StatusBadRequest)
		return
	}

	http.Error(w, "Status Write", statusCode)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Hello handler called")
	w.Write([]byte("hello"))
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Echo handler called with Content Length: %d[byte]", r.ContentLength)
	w.Header().Set("Content-Length", string(r.ContentLength))
	_, err := io.Copy(w, r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
}
