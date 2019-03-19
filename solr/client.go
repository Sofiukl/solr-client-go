package solr

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client - This is solr client expose all the funcanalities that
// required by the client
type Client struct {
	connection     Connection
	queryCriteria  QueryCriteria
	filterCriteria FilterCriteria
}

// SearchResponse - This will be returned to the caller
// type SearchResponse struct {
// 	error bool
// }

// NewSolrClient - Creates new solr client
func NewSolrClient(connection Connection) *Client {
	client := Client{connection: connection, queryCriteria: QueryCriteria{}, filterCriteria: FilterCriteria{}}
	return &client
}

// SetQueryCriteria - sets the query criteria
func (client *Client) SetQueryCriteria(queryCriteria QueryCriteria) *Client {
	client.queryCriteria = queryCriteria
	return client
}

// SetFilerCriteria - sets the filter criteria
func (client *Client) SetFilerCriteria(filterCriteria FilterCriteria) *Client {
	client.filterCriteria = filterCriteria
	return client
}

// requestGet - This makes get request to the Solr
func requestGet(client Client) string {
	queryCrtiteria := client.queryCriteria
	filterCriteria := client.filterCriteria
	conn := client.connection

	queryStr := queryCrtiteria.BuildCriteria()
	filterStr := filterCriteria.BuildCriteria()
	requestPrefixURL := conn.MakeRequestURL()
	URLParam := ""

	if filterStr != "" && queryStr != "" {
		URLParam = filterStr + "&" + queryStr
	} else if filterStr != "" {
		URLParam = filterStr
	} else if queryStr != "" {
		URLParam = queryStr
	}
	if URLParam != "" {
		URLParam = "?" + URLParam
	}
	requestFullPath := requestPrefixURL + URLParam
	fmt.Println("requestFullPath: " + requestFullPath)

	clientReq := &http.Client{}
	req, _ := http.NewRequest("GET", requestFullPath, nil)
	req.Header.Set("accept", "application/json; charset=utf-8")
	resp, err := clientReq.Do(req)

	// resp, err := http.Get(requestFullPath)
	if err != nil {
		fmt.Println("Solr query err ", err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Solr server err ", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("body: " + string(body))
	return string(body)
}

// Search - This is exposed API for search in Solr with the specified query
func (client *Client) Search() string {
	return requestGet(*client)
}
