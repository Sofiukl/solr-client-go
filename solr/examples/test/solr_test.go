package main

import (
	"testing"
	"github.com/sofiukl/solr-client-go/solr/solrqry"
)

var (
	host = "localhost"
	port = "4983"
	root = "solr"
	core = "gettingstarted"
)
func TestSearch(t *testing.T) {
	body := solrqry.NewQueryInterface(solrqry.ConnectionOption{
		Host: host,
		Port: port,
		Root: root,
		Core: core}).
		SetLogLevel("INFO").
		Search(solrqry.SearchOption{
			Edismax: solrqry.EdismaxOption{
				Q:  "2/55",
				Qf: []string{"reportnumber:100", "pdfreportcontent:1"},
			},
			// Q:     []string{"*:*"},
			Fq: []string{"type:PDFReport", "pdfreporttemplateid:2330"},
			Fl: []string{"reportnumber", "score"},
			Facet: solrqry.FacetOption{
				On:    true,
				Field: "reportnumber",
			},
			Sort:  []string{"score:desc", "reportnumber:desc"},
			Start: 0,
			Rows:  12}, func postProcess(solrRes string) string {
				return solrRes
			})
	_ = body
}

func TestSelect(t *testing.T) {
	body := solrqry.NewQueryInterface(solrqry.ConnectionOption{
		Host: "****",
		Port: "****",
		Root: "solr",
		Core: "****"}).
		SetLogLevel("ERROR").
		Select("facet.field=reportnumber&facet=on&fl=reportnumber,score,pdfreportcontent&fq=pdfreporttemplateid:2330&fq=type:PDFReport&q=(pdfreportcontent:2/55)^100")
	_ = body
}
