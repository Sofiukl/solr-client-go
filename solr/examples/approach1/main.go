package main

import (
	"github.com/sofiukl/solr-client-go/solr/solrqry"
)

var (
	host = "localhost"
	port = "4983"
	root = "solr"
	core = "gettingstarted"
)

func postProcess(solrRes string) string {
	return solrRes
}
func main() {
	solrqry.NewQueryInterface(solrqry.ConnectionOption{
		Host: host,
		Port: port,
		Root: root,
		Core: core}).
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
			Rows:  12}, postProcess)
}
