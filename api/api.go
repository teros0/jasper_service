package api

import (
	"doReport/jasper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UploadRequest struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ExportRequest struct {
	Name      string            `json:"name"`
	Format    string            `json:"format"`
	Arguments []jasper.Argument `json:"arguments"`
}

type Argument struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

func UploadReport(w http.ResponseWriter, r *http.Request) {
	var req UploadRequest
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Wrong body", http.StatusBadRequest)
	}
	if err := json.Unmarshal(bs, &req); err != nil {
		http.Error(w, "Wrong body", http.StatusBadRequest)
	}
	jasper.PostReport(req.Name, req.Content)
}

func ExportReport(w http.ResponseWriter, r *http.Request) {
	var req ExportRequest
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Wrong body", http.StatusBadRequest)
	}
	if err := json.Unmarshal(bs, &req); err != nil {
		fmt.Println(err)
		http.Error(w, "Wrong body", http.StatusBadRequest)
	}
	doc, err := jasper.ExportReport(req.Name, req.Format, req.Arguments)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = w.Write(doc)
	if err != nil {
		fmt.Println(err)
		return
	}
}
