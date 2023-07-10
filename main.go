package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var logger = log.New(os.Stderr, "[TESTSERVER]", log.LUTC|log.LstdFlags)

func main() {
	var logFile string
	var serverAddr string
	var isTLS bool
	flag.StringVar(&logFile, "logfile", "", "file name for log, output to stdout if empty")
	flag.StringVar(&serverAddr, "server", "0.0.0.0:10000", "bind addr for server")
	flag.BoolVar(&isTLS, "tls", false, "use tls if true")
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
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/echo", echoHandler).Methods("POST")
	r.HandleFunc("/status/{statusCode}", statusHandler)
	r.HandleFunc("/download/{bytesize}", downloadHandler).Methods("GET")
	r.HandleFunc("/discard", discardHandler).Methods("POST")

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

	if isTLS {
		logger.Println("Start Server with TLS")
		if err := http.ListenAndServeTLS(serverAddr, "certs/localhost.crt", "certs/localhost.key", corsOpts.Handler(r)); err != nil {
			logger.Printf("%+v", err)
			os.Exit(1)
		}
	} else {
		logger.Println("Start Server")
		if err := http.ListenAndServe(serverAddr, corsOpts.Handler(r)); err != nil {
			logger.Printf("%+v", err)
			os.Exit(1)
		}
	}
}
