package solrqry

import (
	"strconv"
	"strings"

	"github.com/sofiukl/solr-client-go/solr/common"
)

// QfCriteriaOption - This is the structure for each criteria
type QfCriteriaOption struct {
	Field string
	Exp   int
}

// EdismaxQueryCriteria - This is the edix max structure
type EdismaxQueryCriteria struct {
	queryKeyword string
	qString      string
	qfCriterias  []string
}

// NewEdismaxQueryCriteriaObject - This returns the critria object
func NewEdismaxQueryCriteriaObject() *EdismaxQueryCriteria {
	edismaxQueryCriteria := EdismaxQueryCriteria{queryKeyword: "edismax", qString: "", qfCriterias: []string{}}
	return &edismaxQueryCriteria
}

// AddQCriteria - This func adds one criteria
func (edismaxCriteria *EdismaxQueryCriteria) AddQCriteria(q string) *EdismaxQueryCriteria {
	edismaxCriteria.qString = solr.URLEncoded(q)
	return edismaxCriteria
}

// AddQfCriteria - This func adds one criteria
func (edismaxCriteria *EdismaxQueryCriteria) AddQfCriteria(qfCriteria QfCriteriaOption) *EdismaxQueryCriteria {
	criteria := qfCriteria.Field + "^" + strconv.Itoa(qfCriteria.Exp)
	edismaxCriteria.qfCriterias = append(edismaxCriteria.qfCriterias, criteria)
	return edismaxCriteria
}

// BuildCriteria - This func builds the whole criteria
func (edismaxCriteria EdismaxQueryCriteria) BuildCriteria() string {
	if len(edismaxCriteria.qfCriterias) == 0 {
		return ""
	}
	criteriaStr := strings.Join(edismaxCriteria.qfCriterias, solr.URLEncoded(" "))
	edismaxq := []string{"q=" + edismaxCriteria.qString, "defType=" + edismaxCriteria.queryKeyword, "qf=" + criteriaStr}

	return strings.Join(edismaxq, "&")
}
