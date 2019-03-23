package main

import (
	"github.com/sofiukl/solr-client-go/solr"
)

func main() {
	conn := solr.NewConnection(solr.ConnectionOption{
		Host: "192.168.99.100",
		Port: "8983",
		Root: "solr",
		Core: "gettingstarted"})

	// edismaxCriteria := solr.NewEdismaxQueryCriteriaObject().
	// 	AddQCriteria("*CA*").
	// 	AddQfCriteria(solr.QfCriteriaOption{Field: "id", Exp: 1}).
	// 	AddQfCriteria(solr.QfCriteriaOption{Field: "address_s", Exp: 100})

	queryCriteria := solr.NewQueryCrtiteriaObject().
		AddCriteria(solr.QueryCriteriaOption{Fieldname: "*", Fieldvalue: "*"})

	filterCriteria := solr.NewFilterCrtiteriaObject().
		AddCriteria(solr.FilterCriteriaOption{Fieldname: "id", Fieldvalue: "a*"})

	paginationCriteria := solr.NewPaginationCriteriaObject().
		AddCriteria(solr.PaginationOption{Start: 0, Rows: 5})

	flCriteria := solr.NewFlCriteriaObject().
		AddCriteria(solr.FlOption{Fields: []string{"id", "score", "address_s"}})

	sortCriteria := solr.NewSortCriteriaObject().
		AddCriteria(solr.SortCriteriaOption{Sortfield: "score", Sortorder: "desc"}).
		AddCriteria(solr.SortCriteriaOption{Sortfield: "id", Sortorder: "desc"})

	solr.NewSolrClient(*conn).
		SetQueryCriteria(*queryCriteria).
		//SetEdismaxQueryCriteria(*edismaxCriteria).
		SetFilerCriteria(*filterCriteria).
		SetPaginationCriteria(*paginationCriteria).
		SetFlCriteria(*flCriteria).
		SetSortCriteria(*sortCriteria).
		Search()
}
