package main

import (
	"fmt"

	"github.com/sofiukl/solr-client-go/solr"
	"github.com/sofiukl/solr-client-go/solr/caller_interface"
)

func main() {
	fmt.Println("Solr Client go started its work..")

	// Approach 1:
	fmt.Println("Approach 1 ..........................................")

	solrconnector.NewInterface().
		Search(solrconnector.SearchOption{
			Edismax: solrconnector.EdismaxOption{
				Q:  "a*",
				Qf: []string{"id:100"},
			},
			// Q:     []string{"*:*"},
			Fq:    []string{"id:a*"},
			Fl:    []string{"id", "score"},
			Sort:  []string{"id:asc"},
			Start: 0,
			Rows:  12})

	// Approach 2:
	fmt.Println("Approach 2 ..........................................")
	conn := solr.NewConnection("192.168.99.100", "8983", "solr", "gettingstarted")

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
	fmt.Println("Solr Client go completed its work.")
}
