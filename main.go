package main

import "fmt"

func main() {
	fmt.Println("Solr Client go started its work..")

	conn := NewConnection("192.168.50.241", "4983", "solr", "BetaCSCollection")

	queryCriteria := NewQueryCrtiteriaObject()
	queryCriteria.AddCriteria(QueryCriteriaOption{fieldname: "rfiid", fieldvalue: "10173069"})
	queryCriteria.AddCriteria(QueryCriteriaOption{fieldname: "rficreatebyid", fieldvalue: "10173069"})

	filterCriteria := NewFilterCrtiteriaObject()
	filterCriteria.AddCriteria(FilterCriteriaOption{fieldname: "type", fieldvalue: "RFI"})
	filterCriteria.AddCriteria(FilterCriteriaOption{fieldname: "pwaccountid", fieldvalue: "11610579"})

	solrClient := NewSolrClient(*conn)
	solrClient.SetQueryCriteria(*queryCriteria)
	solrClient.SetFilerCriteria(*filterCriteria)
	solrClient.Search()
}
