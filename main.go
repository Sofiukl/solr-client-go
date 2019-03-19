package main

import (
	"fmt"

	"github.com/sofiukl/solr-client-go/solr"
)

func main() {
	fmt.Println("Solr Client go started its work..")

	conn := solr.NewConnection("192.168.50.241", "4983", "solr", "BetaCSCollection")

	queryCriteria := solr.NewQueryCrtiteriaObject().
		AddCriteria(solr.QueryCriteriaOption{Fieldname: "rfiid", Fieldvalue: "10173069"}).
		AddCriteria(solr.QueryCriteriaOption{Fieldname: "rficreatebyid", Fieldvalue: "10173069"})

	filterCriteria := solr.NewFilterCrtiteriaObject().
		AddCriteria(solr.FilterCriteriaOption{Fieldname: "type", Fieldvalue: "RFI"}).
		AddCriteria(solr.FilterCriteriaOption{Fieldname: "pwaccountid", Fieldvalue: "11610579"})

	solr.NewSolrClient(*conn).
		SetQueryCriteria(*queryCriteria).
		SetFilerCriteria(*filterCriteria).
		Search()
	fmt.Println("Solr Client go completed its work.")
}
