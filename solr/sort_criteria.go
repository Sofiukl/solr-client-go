package solr

import (
	"strings"
)

// SortCriteriaOption - This is the option for sort criteria
type SortCriteriaOption struct {
	Sortfield string
	Sortorder string
}

// SortCriteria - This conains all the criterias
type SortCriteria struct {
	sortKeyword string
	separator   string
	critarias   []string
}

// NewSortCriteriaObject - This creates new sort criteria object
func NewSortCriteriaObject() *SortCriteria {
	sortCriteria := SortCriteria{sortKeyword: "sort", separator: ",", critarias: []string{}}
	return &sortCriteria
}

// AddCriteria - This function add the criterias
func (sortCriteria *SortCriteria) AddCriteria(criteriaOption SortCriteriaOption) *SortCriteria {
	criteriaStr := urlEncoded(criteriaOption.Sortfield + " " + criteriaOption.Sortorder)
	sortCriteria.critarias = append(sortCriteria.critarias, criteriaStr)
	return sortCriteria
}

// BuildCriteria - This builds the whole criteria
func (sortCriteria SortCriteria) BuildCriteria() string {
	if len(sortCriteria.critarias) == 0 {
		return ""
	}
	sortStr := strings.Join(sortCriteria.critarias, sortCriteria.separator)
	return sortCriteria.sortKeyword + "=" + sortStr
}
