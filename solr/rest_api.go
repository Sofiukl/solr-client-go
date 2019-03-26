package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SearchModel - This is the search model
type SearchModel struct {
	SearchString       string
	ReportFilterFields []string
	ReportDate         []string
	SortBy             []string
	SortOrder          []string
	Limit              string
	Offset             string
}

func main() {
	StartServer()
}

// StartServer - This func starts the serever
func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/api/search", DoSearch).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// DoSearch - This is the search handler function
func DoSearch(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	_ = params
	var searchModel SearchModel
	_ = json.NewDecoder(req.Body).Decode(&searchModel)

	fmt.Printf("SearchInput %+v\n", searchModel)
	// retrieve information from search model
}
