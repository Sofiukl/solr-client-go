package solr

import (
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
	return string(body)
}
