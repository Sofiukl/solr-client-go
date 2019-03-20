package solr

import "fmt"

// QueryReqBuilder - This is he request builder structure
type QueryReqBuilder struct {
	queryReq string
}

// NewQueryReqBuilder - This is the rquest builder
func NewQueryReqBuilder(conn Connection, queryCrtiteria QueryCriteria, filterCriteria FilterCriteria) QueryReqBuilder {

	builders := []CriteriaBuilder{queryCrtiteria, filterCriteria}
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
