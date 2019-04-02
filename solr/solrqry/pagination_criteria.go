package solrqry

import (
	"strconv"
	"strings"
)

// PaginationOption - Pagination options
type PaginationOption struct {
	Start int
	Rows  int
}

// PaginationCriteria - structure
type PaginationCriteria struct {
	startKeyword string
	rowsKeyword  string
	critarias    []string
}

// NewPaginationCriteriaObject - This function creates the Pagination object
func NewPaginationCriteriaObject() *PaginationCriteria {
	pagination := PaginationCriteria{startKeyword: "start", rowsKeyword: "rows", critarias: []string{}}
	return &pagination
}

// AddCriteria - This adds the pagination criteria
func (pagination *PaginationCriteria) AddCriteria(paginationOption PaginationOption) *PaginationCriteria {
	criteriaStrStart := pagination.startKeyword + "=" + strconv.Itoa(paginationOption.Start)
	pagination.critarias = append(pagination.critarias, criteriaStrStart)
	if paginationOption.Rows != 0 {
		criteriaStrRows := pagination.rowsKeyword + "=" + strconv.Itoa(paginationOption.Rows)
		pagination.critarias = append(pagination.critarias, criteriaStrRows)
	}
	return pagination
}

// BuildCriteria - This builds the whole criteria
func (pagination PaginationCriteria) BuildCriteria() string {
	if len(pagination.critarias) == 0 {
		return ""
	}
	paginationStr := strings.Join(pagination.critarias, "&")
	return paginationStr
}
