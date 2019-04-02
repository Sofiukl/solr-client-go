package main

import (
	example "github.com/sofiukl/solr-client-go/examples"
	solr "github.com/sofiukl/solr-client-go/solr/common"
	"github.com/sofiukl/solr-client-go/solr/solrqry"
)

func main() {
	conn := solr.NewConnection(solr.ConnectionOption{
		Host: example.Host,
		Port: example.Port,
		Root: example.Root,
		Core: example.Core})

	// edismaxCriteria := solr.NewEdismaxQueryCriteriaObject().
	// 	AddQCriteria("*CA*").
	// 	AddQfCriteria(solr.QfCriteriaOption{Field: "id", Exp: 1}).
	// 	AddQfCriteria(solr.QfCriteriaOption{Field: "address_s", Exp: 100})

	queryCriteria := solrqry.NewQueryCrtiteriaObject().
		AddCriteria(solrqry.QueryCriteriaOption{Fieldname: "*", Fieldvalue: "*"})

	filterCriteria := solrqry.NewFilterCrtiteriaObject().
		AddCriteria(solrqry.FilterCriteriaOption{Fieldname: "id", Fieldvalue: "a*"})

	paginationCriteria := solrqry.NewPaginationCriteriaObject().
		AddCriteria(solrqry.PaginationOption{Start: 0, Rows: 5})

	flCriteria := solrqry.NewFlCriteriaObject().
		AddCriteria(solrqry.FlOption{Fields: []string{"id", "score", "address_s"}})

	sortCriteria := solrqry.NewSortCriteriaObject().
		AddCriteria(solrqry.SortCriteriaOption{Sortfield: "score", Sortorder: "desc"}).
		AddCriteria(solrqry.SortCriteriaOption{Sortfield: "id", Sortorder: "desc"})

	solrqry.NewSolrQueryClient(*conn).
		SetQueryCriteria(*queryCriteria).
		//SetEdismaxQueryCriteria(*edismaxCriteria).
		SetFilerCriteria(*filterCriteria).
		SetPaginationCriteria(*paginationCriteria).
		SetFlCriteria(*flCriteria).
		SetSortCriteria(*sortCriteria).
		Search()
}
