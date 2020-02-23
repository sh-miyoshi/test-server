package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"strconv"
)

var logger = log.New(os.Stderr, "[TESTSERVER]", log.LUTC|log.LstdFlags)

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

func main() {
	var logFile string
	var serverAddr string
	flag.StringVar(&logFile, "logfile", "", "file name for log, output to stdout if empty")
	flag.StringVar(&serverAddr, "server", "0.0.0.0:10000", "bind addr for server")
	flag.Parse()

	// output log to file
	if logFile != "" {
		file, err := os.Create(logFile)
		if err != nil {
			logger.Printf("failed to create logFile %v", logFile)
			os.Exit(1)
		}
		//do not call file.Close() because logger write log through file.Writer

		logger.SetOutput(file)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", defaultHandler).Methods("GET")
	r.HandleFunc("/status/{statusCode}", statusHandler)

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
	})

	logger.Println("Start Server")
	if err := http.ListenAndServe(serverAddr, corsOpts.Handler(r)); err != nil {
		logger.Printf("%+v", err)
		os.Exit(1)
	}
}
