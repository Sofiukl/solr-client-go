package solrqry

import (
	"strconv"
	"strings"

	solr "github.com/sofiukl/solr-client-go/solr/common"
)

// ResultPostprocess - Defines the post processor on raw result returned by the Solr
type ResultPostprocess interface {
	postProcess(solrRes string) string
}

// ConnectionOption - This information needs to be passed during connection creation
type ConnectionOption struct {
	Host string
	Port string
	Root string
	Core string
}

// Requestor - This is the interface to the Solr Core
// From here the functionality will be exposed to the caller of this (solr-client-go) project
type Requestor struct {
	conn     *solr.Connection
	loglevel string
}

// EdismaxOption - This is structure for Edismax
type EdismaxOption struct {
	Q  string
	Qf []string
}

// FacetOption - This is the facet object
type FacetOption struct {
	On     bool
	Field  string
	Prefix string
	Query  string
}

// SearchOption - Caller of the search function will provide this datas
type SearchOption struct {
	Q       []string
	Edismax EdismaxOption
	Facet   FacetOption
	Fq      []string
	Fl      []string
	Sort    []string
	Start   int
	Rows    int
}

// NewQueryInterface - This function will be called to create a interface
func NewQueryInterface(connOpt ConnectionOption) Requestor {
	conn := solr.NewConnection(solr.ConnectionOption{
		Host: connOpt.Host,
		Port: connOpt.Port,
		Root: connOpt.Root,
		Core: connOpt.Core})
	solrintf := Requestor{conn: conn, loglevel: "INFO"}
	solrintf.SetLogLevel("INFO")
	return solrintf
}

// SetLogLevel - This sets the log level
func (solrintf Requestor) SetLogLevel(loglevel string) Requestor {
	solrintf.loglevel = loglevel
	solr.InitLogger(loglevel)
	return solrintf
}

// Search - This function is exposed for search
func (solrintf Requestor) Search(searchOption SearchOption, postProcess func(string) string) string {

	queryCriteria := NewQueryCrtiteriaObject()
	filterCriteria := NewFilterCrtiteriaObject()
	flCriteria := NewFlCriteriaObject()
	paginationCriteria := NewPaginationCriteriaObject()
	sortCriteria := NewSortCriteriaObject()
	edismaxCriteria := NewEdismaxQueryCriteriaObject()
	facetCriteria := NewFacetCrtiteriaObject()

	prepareQ(searchOption.Q, queryCriteria)
	prepareFq(searchOption.Fq, filterCriteria)
	prepareFl(searchOption.Fl, flCriteria)
	preparePagination(searchOption.Start, searchOption.Rows, paginationCriteria)
	prepareSort(searchOption.Sort, sortCriteria)
	prepareEdismax(searchOption.Edismax, edismaxCriteria)
	prepareFacet(searchOption.Facet, facetCriteria)

	reposne := NewSolrQueryClient(*solrintf.conn).
		SetQueryCriteria(*queryCriteria).
		SetEdismaxQueryCriteria(*edismaxCriteria).
		SetFilerCriteria(*filterCriteria).
		SetFlCriteria(*flCriteria).
		SetPaginationCriteria(*paginationCriteria).
		SetSortCriteria(*sortCriteria).
		SetFacetCriteria(*facetCriteria).
		Search(postProcess)
	return reposne
}

// Select - This executes the specified request as is
func (solrintf Requestor) Select(queryURL string) string {
	conn := *solrintf.conn
	requestPrefixURL := conn.MakeRequestURL()
	queryReq := requestPrefixURL + "?" + queryURL
	solr.GetDebugLogger().Println("Full Req: ", queryReq)
	rb := ReqBuilder{queryReq: queryReq}
	body := solr.HandleGetReq(rb.queryReq)
	return body
}
func prepareQ(Q []string, queryCriteria *QueryCriteria) {
	for _, qField := range Q {
		fieldwithvalue := strings.Split(qField, ":")
		queryCriteria.
			AddCriteria(QueryCriteriaOption{Fieldname: fieldwithvalue[0], Fieldvalue: fieldwithvalue[1]})
	}
}

func prepareFq(Fq []string, filterCriteria *FilterCriteria) {
	for _, fqField := range Fq {
		fieldwithvalue := strings.Split(fqField, ":")
		filterCriteria.
			AddCriteria(FilterCriteriaOption{Fieldname: fieldwithvalue[0], Fieldvalue: fieldwithvalue[1]})
	}
}

func prepareFl(Fl []string, flCriteria *FlCriteria) {
	flCriteria.
		AddCriteria(FlOption{Fields: Fl})
}

func preparePagination(start int, rows int, paginationCriteria *PaginationCriteria) {
	paginationCriteria.
		AddCriteria(PaginationOption{Start: start, Rows: rows})
}

func prepareSort(sort []string, sortCriteria *SortCriteria) {
	for _, sortField := range sort {
		fieldwithvalue := strings.Split(sortField, ":")
		sortCriteria.
			AddCriteria(SortCriteriaOption{Sortfield: fieldwithvalue[0], Sortorder: fieldwithvalue[1]})
	}
}

func prepareEdismax(edismaxOption EdismaxOption, edismaxCriteria *EdismaxQueryCriteria) {
	edismaxCriteria.
		AddQCriteria(edismaxOption.Q)
	for _, edismaxField := range edismaxOption.Qf {
		fieldwithvalue := strings.Split(edismaxField, ":")
		exp, _ := strconv.Atoi(fieldwithvalue[1])
		edismaxCriteria.
			AddQfCriteria(QfCriteriaOption{Field: fieldwithvalue[0], Exp: exp})
	}
}

func prepareFacet(facetOption FacetOption, facetCriteria *FacetCriteria) {
	if !facetOption.On {
		return
	}
	facetCriteria.
		AddCriteria(FacetCriteriaOption{
			Fieldname: facetOption.Field,
			Prefix:    facetOption.Prefix,
			Query:     facetOption.Query})
}
