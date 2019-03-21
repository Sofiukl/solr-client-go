package main

import (
	"fmt"

	"github.com/sofiukl/solr-client-go/solr"
)

func main() {
	fmt.Println("Solr Client go started its work..")

	conn := solr.NewConnection("192.168.99.100", "8983", "solr", "gettingstarted")

	// queryCriteria := solr.NewQueryCrtiteriaObject().
	// 	AddCriteria(solr.QueryCriteriaOption{Fieldname: "address_s", Fieldvalue: "46221 Landing Parkway Fremont, CA 94538"})

	filterCriteria := solr.NewFilterCrtiteriaObject().
		//AddCriteria(solr.FilterCriteriaOption{Fieldname: "compName_s", Fieldvalue: "A-Data Technology"}).
		AddCriteria(solr.FilterCriteriaOption{Fieldname: "id", Fieldvalue: "a*"})

	paginationCriteria := solr.NewPaginationCriteriaObject().
		AddCriteria(solr.PaginationOption{Start: 0, Rows: 2})

	flCriteria := solr.NewFlCriteriaObject().
		AddCriteria(solr.FlOption{Fields: []string{"id", "score"}})

	sortCriteria := solr.NewSortCriteriaObject().
		AddCriteria(solr.SortCriteriaOption{Sortfield: "score", Sortorder: "desc"}).
		AddCriteria(solr.SortCriteriaOption{Sortfield: "id", Sortorder: "desc"})

	solr.NewSolrClient(*conn).
		// SetQueryCriteria(*queryCriteria).
		SetFilerCriteria(*filterCriteria).
		SetPaginationCriteria(paginationCriteria).
		SetFlCriteria(flCriteria).
		SetSortCriteria(sortCriteria).
		Search()
	fmt.Println("Solr Client go completed its work.")
}
