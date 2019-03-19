package solr

import (
	"strings"
)

// FilterCriteria - This conains all the filter criterias
type FilterCriteria struct {
	filterKeyword string
	critarias     []string
}

// CriteriaOption - This is the option for filter criteria
type FilterCriteriaOption struct {
	Fieldname  string
	Fieldvalue string
}

// NewFilterCrtiteriaObject - This creates new filter criteria object
func NewFilterCrtiteriaObject() *FilterCriteria {
	filterCriteria := FilterCriteria{filterKeyword: "fq", critarias: []string{}}
	return &filterCriteria
}

// AddCriteria - This function add the criterias
func (filterCriteria *FilterCriteria) AddCriteria(criteria FilterCriteriaOption) *FilterCriteria {
	criteriaStr := filterCriteria.filterKeyword + "=" + criteria.Fieldname + ":" + criteria.Fieldvalue
	filterCriteria.critarias = append(filterCriteria.critarias, criteriaStr)
	return filterCriteria
}

// BuildCriteria - This builds the whole criteria
func (filterCriteria *FilterCriteria) BuildCriteria() string {
	if len(filterCriteria.critarias) == 0 {
		return ""
	}

	filerStr := strings.Join(filterCriteria.critarias, "&")
	return filerStr
}
