package main

import (
	"testing"

	example "github.com/sofiukl/solr-client-go/examples"
	"github.com/sofiukl/solr-client-go/solr/solrqry"
)

func TestSearch(t *testing.T) {
	body := solrqry.NewQueryInterface(solrqry.ConnectionOption{
		Host: example.Host,
		Port: example.Port,
		Root: example.Root,
		Core: example.Core}).
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
			Rows:  12})
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