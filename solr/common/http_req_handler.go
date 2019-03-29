package solr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	reqMethodGet    = "GET"
	reqMethodPost   = "POST"
	reqMethodPatch  = "PATCH"
	reqMethodPut    = "PUT"
	reqMethodDelete = "DELETE"
)

// Docs - This is the response structure
type Docs struct {
	Pinprojectid int
	Modifyby     int
	Createdate   string
}

// Response - This is resposne
type Response struct {
	NumFound int
	Start    int
	Docs     []Docs
}

// SResponse - This is the response strucure
type SResponse struct {
	Response Response
}

// HandleGetReq - This handles the GET request to Solr Core
func HandleGetReq(URL string) string {
	clientReq := &http.Client{}
	req, _ := http.NewRequest(reqMethodGet, URL, nil)
	req.Header.Set("accept", "application/json; charset=utf-8")
	resp, err := clientReq.Do(req)

	if err != nil {
		fmt.Println("Solr server err ", err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Solr server err ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		GetErrorLogger().Println("Error reading request.. ", err)
	}
	solrRes := SResponse{}
	if err := json.Unmarshal([]byte(string(body)), &solrRes); err != nil {
		panic(err)
	}
	fmt.Println(solrRes)
	return string(body)
}
