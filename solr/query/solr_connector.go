package solrqry

import (
	"strconv"
	"strings"

	"github.com/sofiukl/solr-client-go/solr/common"
)

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
	conn *solr.Connection
}

// EdismaxOption - This is structure for Edismax
type EdismaxOption struct {
	Q  string
	Qf []string
}

// SearchOption - Caller of the search function will provide this datas
type SearchOption struct {
	Q       []string
	Edismax EdismaxOption
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
	solrintf := Requestor{conn: conn}
	return solrintf
}

// Search - This function is exposed for search
func (solrintf Requestor) Search(searchOption SearchOption) string {

	queryCriteria := NewQueryCrtiteriaObject()
	filterCriteria := NewFilterCrtiteriaObject()
	flCriteria := NewFlCriteriaObject()
	paginationCriteria := NewPaginationCriteriaObject()
	sortCriteria := NewSortCriteriaObject()
	edismaxCriteria := NewEdismaxQueryCriteriaObject()

	prepareQ(searchOption.Q, queryCriteria)
	prepareFq(searchOption.Fq, filterCriteria)
	prepareFl(searchOption.Fl, flCriteria)
	preparePagination(searchOption.Start, searchOption.Rows, paginationCriteria)
	prepareSort(searchOption.Sort, sortCriteria)
	prepareEdismax(searchOption.Edismax, edismaxCriteria)

	reposne := NewSolrQueryClient(*solrintf.conn).
		SetQueryCriteria(*queryCriteria).
		SetEdismaxQueryCriteria(*edismaxCriteria).
		SetFilerCriteria(*filterCriteria).
		SetFlCriteria(*flCriteria).
		SetPaginationCriteria(*paginationCriteria).
		SetSortCriteria(*sortCriteria).
		Search()
	return reposne
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