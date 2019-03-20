package solr

import "strings"

// Operators
const (
	OrOperator = "or"
)

// QueryCriteriaOption - This is the option for filter criteria
type QueryCriteriaOption struct {
	Fieldname  string
	Fieldvalue string
}

// QueryCriteria - This conains all the criterias
type QueryCriteria struct {
	queryKeyword string
	critarias    []string
	CriteriaBuilder
}

// NewQueryCrtiteriaObject - This creates new query criteria object
func NewQueryCrtiteriaObject() *QueryCriteria {
	queryCriteria := QueryCriteria{queryKeyword: "q", critarias: []string{}}
	return &queryCriteria
}

// AddCriteria - This function add the criterias
func (queryCriteria *QueryCriteria) AddCriteria(criteria QueryCriteriaOption) *QueryCriteria {
	criteriaStr := criteria.Fieldname + ":" + criteria.Fieldvalue
	queryCriteria.critarias = append(queryCriteria.critarias, criteriaStr)
	return queryCriteria
}

// BuildCriteria - This builds the whole criteria
func (queryCriteria QueryCriteria) BuildCriteria() string {
	if len(queryCriteria.critarias) == 0 {
		return ""
	}
	filterStr := strings.Join(queryCriteria.critarias, "%20"+OrOperator+"%20")
	return "q=" + filterStr
}