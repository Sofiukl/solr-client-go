[![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/Sofiukl/solr-client-go/solr)

# solr-client-go [Work In Progress]

Solr client for  Golang developers.
At least Solr v4 is required if you want to use all the supported features.
Apache Solr documentation link https://wiki.apache.org/solr/

# Install
go get github.com/Sofiukl/solr-client-go

# Features
1.1 Search
1.2 Add [WIP]
1.3 Update [WIP]
1.4 Delete [WIP]
1.5 Commit [WIP]
1.6 Rollback [WIP]
1.7 Optimize [WIP]

# Usages

	package main

	import (
		"github.com/sofiukl/solr-client-go/solr/solrqry"
	)

	func main() {
		solrqry.NewQueryInterface(solrqry.ConnectionOption{
			Host: "192.168.99.100",
			Port: "8983",
			Root: "solr",
			Core: "gettingstarted"}).
			Search(solrqry.SearchOption{
				Edismax: solrqry.EdismaxOption{
					Q:  "a*",
					Qf: []string{"id:100"},
				},
				// Q:     []string{"*:*"},
				Fq:    []string{"id:a*"},
				Fl:    []string{"id", "score"},
				Sort:  []string{"id:asc"},
				Start: 0,
				Rows:  12})
	}


See example folder for details

# Documentation
godoc -http=:6060

{{server}}/pkg/github.com/sofiukl/solr-client-go/

# License
MIT

Feel free to raise any issue. This package is under development and will be highly maintainded.
