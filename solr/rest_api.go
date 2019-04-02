package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sofiukl/solr-client-go/solr/solrqry"

	"github.com/gorilla/mux"
)

// SearchModel - This is the search model
type SearchModel struct {
	PINProjectID          string
	Type                  string
	SearchString          string
	ReportFilterFields    string
	ReportFilterValues    string
	ReportDate            string
	SortBy                string
	SortOrder             string
	Limit                 string
	Offset                string
	ContentSearchRequired int
}

const (
	port = ":8000"
)

var (
	searchFields       = []string{"reportnumber", "pdfreportcontent"}
	flFieldsConfigured = []string{"score", "reportnumber"}
)

func main() {
	StartServer()
}

// StartServer - This func starts the serever
func StartServer() {
	fmt.Println("Server is started at ", port)
	router := mux.NewRouter()
	router.HandleFunc("/api/search", DoSearch).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}

// DoSearch - This is the search handler function
func DoSearch(w http.ResponseWriter, req *http.Request) {

	searchType := getAttributeVal(req, "type")
	searchString := getAttributeVal(req, "searchString")
	limit, _ := strconv.Atoi(getAttributeVal(req, "limit"))
	offset, _ := strconv.Atoi(getAttributeVal(req, "offset"))
	sortBys := strings.Split(getAttributeVal(req, "sortBy"), ",")
	sortOrders := strings.Split(getAttributeVal(req, "sortOrder"), ",")
	fqFields := strings.Split(getAttributeVal(req, "fqFields"), ",")
	fqValues := strings.Split(getAttributeVal(req, "fqValues"), ",")
	flFields := strings.Split(getAttributeVal(req, "flFields"), ",")

	// validations
	errorIfBlank("type", searchType)
	errorIfBlank("searchString", searchString)

	// prepare engine parameters
	Q := prepareQ(searchString, searchFields)
	Fq := prepareFq(searchType, fqFields, fqValues)
	Sort := prepareSort(sortBys, sortOrders)
	Fl := prepareFl(flFields)

	// Call Search Engine

	body := solrqry.NewQueryInterface(solrqry.ConnectionOption{
		Host: "192.168.50.241",
		Port: "4983",
		Root: "solr",
		Core: "BetaCSCollection"}).
		SetLogLevel("INFO").
		Search(solrqry.SearchOption{
			Q:     Q,
			Fq:    Fq,
			Fl:    Fl,
			Sort:  Sort,
			Start: offset,
			Rows:  limit})
	_ = body
}
func errorIfBlank(field string, val string) {
	if val == "" {
		log.Fatal("please specify the ", field)
	}
}
func prepareFl(flFields []string) []string {
	if len(flFields) == 0 {
		return flFieldsConfigured
	}
	return flFields
}
func prepareSort(sortBys []string, sortOrders []string) []string {
	// prepare Sort   []string{"score:desc", "reportnumber:desc"}
	var Sort []string
	for index, sortField := range sortBys {
		temp := sortField + ":" + sortOrders[index]
		Sort = append(Sort, temp)
	}
	return Sort
}
func prepareQ(searchString string, searchFields []string) []string {
	// prepare Q
	var Q []string

	for _, searchField := range searchFields {
		temp := ""
		if searchField == "reportnumber" {
			_, err := strconv.Atoi(searchString)
			if err == nil {
				temp = searchField + ":" + searchString
				Q = append(Q, temp)
			}
		} else {
			temp = searchField + ":" + searchString
			Q = append(Q, temp)
		}

	}
	return Q
}

func prepareFq(searchType string, fqFields []string, fqValues []string) []string {
	// prepare Fq  []string{"type:PDFReport", "pdfreporttemplateid:2330"}
	var Fq []string
	Fq = append(Fq, "type:"+searchType)
	for index, fqField := range fqFields {
		temp := fqField + ":" + fqValues[index]
		Fq = append(Fq, temp)
	}
	return Fq
}

func getAttributeVal(req *http.Request, key string) string {
	keys, _ := req.URL.Query()[key]
	if len(keys) == 0 {
		fmt.Println(key, ":", "")
		return ""
	}
	fmt.Println(key, ":", keys[0])
	return keys[0]
}
