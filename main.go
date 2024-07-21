package main

import (
	"log"
	"net/http"
	"web-page-analyser/pkg/handlers"
)

func main() {
	http.HandleFunc("/analysis", handlers.AnalysisHandler)
	http.HandleFunc("/analyse", handlers.AnalyseHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
