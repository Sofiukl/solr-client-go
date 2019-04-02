package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/sofiukl/solr-client-go/solr/solrqry"

	"github.com/gorilla/mux"
)

// Host - This is solr server port
// Port - This is solr connection port
// Root - This is the root path
// Core - This is the name of the solr collection
var (
	Host = "192.168.50.241"
	Port = "4983"
	Root = "solr"
	Core = "BetaCSCollection"
)

const (
	port = ":8000"
)

var (
	searchFields = []string{"reportnumber", "pdfreportcontent"}
	// searchFields = []string{"id", "compName_s"}
	flFieldsConfigured = []string{"score", "pinprojectid", "reportnumber", "reportname",
		"createby", "createdate", "pdfreportcontent"}
	// flFieldsConfigured = []string{"score", "id", "compName_s"}
)

func postProcess(solrRes string) string {
	transformedRes := SResponse{}
	if err := json.Unmarshal([]byte(string(solrRes)), &transformedRes); err != nil {
		panic(err)
	}
	fmt.Printf("Response: %+v\n", transformedRes)
	b, err := json.Marshal(transformedRes)
	if err != nil {
		fmt.Println("Failed to post process the solr response")
	}
	return string(b)
}

// Docs - This is the response structure
type Docs struct {
	// ID        string `json:"id,omitempty"`
	// CompNameS string `json:"compName_s,omitempty"`
	// AddressS  string `json:"address_s,omitempty"`
	// Version   int    `json:"__version_,omitempty"`

	PinprojectID          int    `json:"pinprojectid,omitempty"`
	ModifyBy              int    `json:"modifyby,omitempty"`
	ReportDate            string `json:"reportdate,omitempty"`
	CreateDate            string `json:"createdate,omitempty"`
	PdfReportModifyByname string `json:"pdfreportmodifybyname,omitempty"`
	Type                  string `json:"type,omitempty"`
	PdfreportID           int    `json:"pdfreportid,omitempty"`
	Createby              int    `json:"createby,omitempty"`
	Modifydate            string `json:"modifydate,omitempty"`
	Approvaltype          int    `json:"approvaltype,omitempty"`
	Reportname            string `json:"reportname,omitempty"`
	Reportnumber          int    `json:"reportnumber,omitempty"`
	Pdfreporttemplateid   int    `json:"pdfreporttemplateid,omitempty"`
	Pdfreportcreatebyname string `json:"pdfreportcreatebyname,omitempty"`
	Status                int    `json:"status,omitempty"`
	Pdfreportcontent      string `json:"pdfreportcontent,omitempty"`
}

// Response - This is resposne
type Response struct {
	NumFound int    `json:"numFound,omitempty"`
	Start    int    `json:"start,omitempty"`
	Docs     []Docs `json:"docs,omitempty"`
}

// SResponse - This is the response strucure
type SResponse struct {
	Response Response `json:"response,omitempty"`
}

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
	fmt.Println("User Inputs: ")
	searchType := getAttributeVal(req, "type")
	searchString := getAttributeVal(req, "searchString")
	limit, _ := strconv.Atoi(getAttributeVal(req, "limit"))
	offset, _ := strconv.Atoi(getAttributeVal(req, "offset"))
	sortBys := splitAttributeVal(getAttributeVal(req, "sortBy"))
	sortOrders := splitAttributeVal(getAttributeVal(req, "sortOrder"))
	fqFields := splitAttributeVal(getAttributeVal(req, "fqFields"))
	fqValues := splitAttributeVal(getAttributeVal(req, "fqValues"))
	flFields := splitAttributeVal(getAttributeVal(req, "flFields"))

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
		Host: Host,
		Port: Port,
		Root: Root,
		Core: Core}).
		SetLogLevel("INFO").
		Search(solrqry.SearchOption{
			Q:     Q,
			Fq:    Fq,
			Fl:    Fl,
			Sort:  Sort,
			Start: offset,
			Rows:  limit}, postProcess)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bodyJSON, _ := json.Marshal(body)
	w.Write(bodyJSON)
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
	// Fq = append(Fq, "type:"+searchType)
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

func splitAttributeVal(val string) []string {
	values := []string{}
	if val != "" {
		values = strings.Split(val, ",")
	}
	return values
}
