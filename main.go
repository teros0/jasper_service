package main

import (
	"doReport/api"
	"log"
	"net/http"
)

func main() {
	address := ":7777"
	mux := http.NewServeMux()
	mux.HandleFunc("/upload_report", api.UploadReport)
	mux.HandleFunc("/export_report", api.ExportReport)
	log.Fatal(http.ListenAndServe(address, mux))
}
