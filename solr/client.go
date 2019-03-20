package solr

import "fmt"

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

// doSearch - This makes get request to the Solr
func doSearch(client Client) string {
	queryCrtiteria := client.queryCriteria
	filterCriteria := client.filterCriteria
	conn := client.connection

	body := NewQueryReqBuilder(conn, queryCrtiteria, filterCriteria).
		Execute()
	fmt.Println(body)
	return body
}

// Search - This is exposed API for search in Solr with the specified query
func (client *Client) Search() string {
	return doSearch(*client)
}
