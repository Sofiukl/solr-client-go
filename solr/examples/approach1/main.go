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
