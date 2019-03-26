package main

import (
	"testing"

	"github.com/sofiukl/solr-client-go/solr/solrqry"
)

func TestSearch(t *testing.T) {
	body := solrqry.NewQueryInterface(solrqry.ConnectionOption{
		Host: "192.168.50.241",
		Port: "4983",
		Root: "solr",
		Core: "BetaCSCollection"}).
		SetLogLevel("ERROR").
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
		Host: "192.168.50.241",
		Port: "4983",
		Root: "solr",
		Core: "BetaCSCollection"}).
		SetLogLevel("ERROR").
		Select("facet.field=reportnumber&facet=on&fl=reportnumber,score,pdfreportcontent&fq=pdfreporttemplateid:2330&fq=type:PDFReport&q=(pdfreportcontent:2/55)^100")
	_ = body
}
