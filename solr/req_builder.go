package solr

import "fmt"

// QueryReqBuilder - This is he request builder structure
type QueryReqBuilder struct {
	queryReq string
}

// NewQueryReqBuilder - This is the rquest builder
func NewQueryReqBuilder(client Client) QueryReqBuilder {

	queryCrtiteria := client.queryCriteria
	filterCriteria := client.filterCriteria
	paginationCriteria := client.paginationCriteria
	flCriteria := client.flCriteria
	sortCriteria := client.sortCriteria
	conn := client.connection

	builders := []CriteriaBuilder{queryCrtiteria, filterCriteria, paginationCriteria, flCriteria, sortCriteria}
	quryCriterias := NewCriteriaBuilderEngine(builders).Build()

	requestPrefixURL := conn.MakeRequestURL()
	queryReq := requestPrefixURL + quryCriterias
	fmt.Println("queryReq: " + queryReq)
	rb := QueryReqBuilder{queryReq: queryReq}
	return rb
}

// Execute the request
func (rb QueryReqBuilder) Execute() string {
	body := HandleGetReq(rb.queryReq)
	return body
}
