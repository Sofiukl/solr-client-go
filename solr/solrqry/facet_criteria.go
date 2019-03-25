package solrqry

import (
	"fmt"
	"strings"
)

// Facet - This is the facet object
type Facet struct {
	on     bool
	field  string
	prefix string
	query  string
}

// FacetCriteria - Structure for Facet Input Details
type FacetCriteria struct {
	criterias []string
}

// FacetCriteriaOption - This is the option for facet criteria
type FacetCriteriaOption struct {
	Fieldname string
	Prefix    string
	Query     string
}

// NewFacetCrtiteriaObject - This creates new facet criteria object
func NewFacetCrtiteriaObject() *FacetCriteria {
	facetCriteria := FacetCriteria{criterias: []string{}}
	return &facetCriteria
}

// AddCriteria - This function add the criterias
func (facetCriteria *FacetCriteria) AddCriteria(criteria FacetCriteriaOption) *FacetCriteria {
	facet := Facet{}
	if criteria.Fieldname != "" {
		facet.on = true
		facet.field = criteria.Fieldname
		criteriaFacetOn := "facet=on"
		criteriaFacetField := "facet.field=" + facet.field
		facetCriteria.criterias = append(facetCriteria.criterias, criteriaFacetOn)
		facetCriteria.criterias = append(facetCriteria.criterias, criteriaFacetField)
	}

	if criteria.Prefix != "" {
		facet.prefix = criteria.Prefix
		criteriaFacetPrefix := "facet.prefix=" + facet.prefix
		facetCriteria.criterias = append(facetCriteria.criterias, criteriaFacetPrefix)
	}
	if criteria.Query != "" {
		facet.query = criteria.Query
		criteriaFacetQuery := "facet.query=" + facet.query
		facetCriteria.criterias = append(facetCriteria.criterias, criteriaFacetQuery)
	}
	fmt.Println("Facet: ", facet)
	return facetCriteria
}

// BuildCriteria - This builds the whole criteria
func (facetCriteria FacetCriteria) BuildCriteria() string {
	if len(facetCriteria.criterias) == 0 {
		return ""
	}

	facetStr := strings.Join(facetCriteria.criterias, "&")
	return facetStr
}
