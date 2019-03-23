package solrconnector

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sofiukl/solr-client-go/solr"
)

// Requestor - This is the interface to the Solr Core
// From here the functionality will be exposed to the caller of this (solr-client-go) project
type Requestor struct {
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

// NewInterface - This function will be called to create a interface
func NewInterface() Requestor {
	solrintf := Requestor{}
	return solrintf
}

// Search - This function is exposed for search
func (solrintf Requestor) Search(searchOption SearchOption) string {
	conn := solr.NewConnection("192.168.99.100", "8983", "solr", "gettingstarted")
	queryCriteria := solr.NewQueryCrtiteriaObject()
	filterCriteria := solr.NewFilterCrtiteriaObject()
	flCriteria := solr.NewFlCriteriaObject()
	paginationCriteria := solr.NewPaginationCriteriaObject()
	sortCriteria := solr.NewSortCriteriaObject()
	edismaxCriteria := solr.NewEdismaxQueryCriteriaObject()
	prepareQ(searchOption.Q, queryCriteria)
	prepareFq(searchOption.Fq, filterCriteria)
	prepareFl(searchOption.Fl, flCriteria)
	preparePagination(searchOption.Start, searchOption.Rows, paginationCriteria)
	prepareSort(searchOption.Sort, sortCriteria)
	prepareEdismax(searchOption.Edismax, edismaxCriteria)
	fmt.Println(flCriteria)
	reposne := solr.NewSolrClient(*conn).
		SetQueryCriteria(*queryCriteria).
		SetEdismaxQueryCriteria(*edismaxCriteria).
		SetFilerCriteria(*filterCriteria).
		SetFlCriteria(*flCriteria).
		SetPaginationCriteria(*paginationCriteria).
		SetSortCriteria(*sortCriteria).
		Search()
	return reposne
}

func prepareQ(Q []string, queryCriteria *solr.QueryCriteria) {
	for _, qField := range Q {
		fieldwithvalue := strings.Split(qField, ":")
		queryCriteria.
			AddCriteria(solr.QueryCriteriaOption{Fieldname: fieldwithvalue[0], Fieldvalue: fieldwithvalue[1]})
	}
}

func prepareFq(Fq []string, filterCriteria *solr.FilterCriteria) {
	for _, fqField := range Fq {
		fieldwithvalue := strings.Split(fqField, ":")
		filterCriteria.
			AddCriteria(solr.FilterCriteriaOption{Fieldname: fieldwithvalue[0], Fieldvalue: fieldwithvalue[1]})
	}
}

func prepareFl(Fl []string, flCriteria *solr.FlCriteria) {
	flCriteria.
		AddCriteria(solr.FlOption{Fields: Fl})
}

func preparePagination(start int, rows int, paginationCriteria *solr.PaginationCriteria) {
	paginationCriteria.
		AddCriteria(solr.PaginationOption{Start: start, Rows: rows})
}

func prepareSort(sort []string, sortCriteria *solr.SortCriteria) {

	for _, sortField := range sort {
		fieldwithvalue := strings.Split(sortField, ":")
		sortCriteria.
			AddCriteria(solr.SortCriteriaOption{Sortfield: fieldwithvalue[0], Sortorder: fieldwithvalue[1]})
	}
}

func prepareEdismax(edismaxOption EdismaxOption, edismaxCriteria *solr.EdismaxQueryCriteria) {
	edismaxCriteria.
		AddQCriteria(edismaxOption.Q)
	for _, edismaxField := range edismaxOption.Qf {
		fieldwithvalue := strings.Split(edismaxField, ":")
		exp, _ := strconv.Atoi(fieldwithvalue[1])
		edismaxCriteria.
			AddQfCriteria(solr.QfCriteriaOption{Field: fieldwithvalue[0], Exp: exp})
	}
}
