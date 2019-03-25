package solrqry

import (
	"fmt"

	solr "github.com/sofiukl/solr-client-go/solr/common"
)

// Client - This is solr client expose all the funcanalities that
// required by the client
type Client struct {
	connection         solr.Connection
	queryCriteria      QueryCriteria
	edismaxCriteria    EdismaxQueryCriteria
	filterCriteria     FilterCriteria
	paginationCriteria PaginationCriteria
	flCriteria         FlCriteria
	sortCriteria       SortCriteria
	facetCriteria      FacetCriteria
}

// SearchResponse - This will be returned to the caller
// type SearchResponse struct {
// 	error bool
// }

// NewSolrQueryClient - Creates new solr client
func NewSolrQueryClient(connection solr.Connection) *Client {
	client := Client{connection: connection, queryCriteria: QueryCriteria{}, filterCriteria: FilterCriteria{}}
	return &client
}

// SetQueryCriteria - sets the query criteria
func (client *Client) SetQueryCriteria(queryCriteria QueryCriteria) *Client {
	client.queryCriteria = queryCriteria
	return client
}

// SetEdismaxQueryCriteria - sets the edismax query criteria
func (client *Client) SetEdismaxQueryCriteria(edismaxCriteria EdismaxQueryCriteria) *Client {
	client.edismaxCriteria = edismaxCriteria
	return client
}

// SetFilerCriteria - sets the filter criteria
func (client *Client) SetFilerCriteria(filterCriteria FilterCriteria) *Client {
	client.filterCriteria = filterCriteria
	return client
}

// SetPaginationCriteria - sets the pagination criteria
func (client *Client) SetPaginationCriteria(paginationCriteria PaginationCriteria) *Client {
	client.paginationCriteria = paginationCriteria
	return client
}

// SetFlCriteria - sets the fl criteria
func (client *Client) SetFlCriteria(flCriteria FlCriteria) *Client {
	client.flCriteria = flCriteria
	return client
}

// SetSortCriteria - sets the sort criteria
func (client *Client) SetSortCriteria(sortCriteria SortCriteria) *Client {
	client.sortCriteria = sortCriteria
	return client
}

// SetFacetCriteria - sets the facet criteria
func (client *Client) SetFacetCriteria(facetCriteria FacetCriteria) *Client {
	client.facetCriteria = facetCriteria
	return client
}

// doSearch - This makes get request to the Solr
func doSearch(client Client) string {
	body := NewQueryReqBuilder(client).
		Execute()
	fmt.Println(body)
	return body
}

// Search - This is exposed API for search in Solr with the specified query
func (client *Client) Search() string {
	return doSearch(*client)
}
