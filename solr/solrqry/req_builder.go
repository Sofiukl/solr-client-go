package solrqry

import (
	solr "github.com/sofiukl/solr-client-go/solr/common"
)

// ReqBuilder - This is he request builder structure
type ReqBuilder struct {
	queryReq string
}

// NewQueryReqBuilder - This is the rquest builder
func NewQueryReqBuilder(client Client) ReqBuilder {

	queryCrtiteria := client.queryCriteria
	edismaxCriteria := client.edismaxCriteria
	filterCriteria := client.filterCriteria
	paginationCriteria := client.paginationCriteria
	flCriteria := client.flCriteria
	sortCriteria := client.sortCriteria
	facetCriteria := client.facetCriteria
	conn := client.connection

	builders := []CriteriaBuilder{queryCrtiteria, edismaxCriteria, filterCriteria, paginationCriteria, flCriteria, sortCriteria, facetCriteria}
	quryCriterias := NewCriteriaBuilderEngine(builders).Build()

	requestPrefixURL := conn.MakeRequestURL()
	queryReq := requestPrefixURL + quryCriterias
	solr.GetDebugLogger().Println("Full Req: ", queryReq)
	rb := ReqBuilder{queryReq: queryReq}
	return rb
}

// Execute the request
func (rb ReqBuilder) Execute() string {
	body := solr.HandleGetReq(rb.queryReq)
	return body
}
