package main

import (
	example "github.com/sofiukl/solr-client-go/examples"
	"github.com/sofiukl/solr-client-go/solr/solrqry"
)

func main() {
	solrqry.NewQueryInterface(solrqry.ConnectionOption{
		Host: example.Host,
		Port: example.Port,
		Root: example.Root,
		Core: example.Core}).
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
