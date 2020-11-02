package main

import (
	"io"
	"io/ioutil"
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
	w.Header().Set("Content-Length", strconv.Itoa(int(r.ContentLength)))
	_, err := io.Copy(w, r.Body)
	if err != nil {
		http.Error(w, "Internal Server Error ", http.StatusInternalServerError)
		return
	}
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	byteSizeStr := vars["bytesize"]
	logger.Printf("Download handler called with bytesize: %s", byteSizeStr)
	byteSize, err := strconv.Atoi(byteSizeStr)
	if err != nil {
		logger.Printf("Failed to convert bytesize: %v", err)
		http.Error(w, "BatRequest", http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/octet-stream")
	w.Header().Add("Content-Length", byteSizeStr)
	reader := newRandReader(byteSize)
	io.Copy(w, reader)
}

func discardHandler(w http.ResponseWriter, r *http.Request) {
	logger.Printf("Discard handler called with Content Length: %d[byte]", r.ContentLength)
	io.Copy(ioutil.Discard, r.Body)
}
