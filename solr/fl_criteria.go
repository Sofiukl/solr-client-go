package solr

import (
	"strings"
)

// FlOption - Fl options
type FlOption struct {
	Fields []string
}

// FlCriteria - structure
type FlCriteria struct {
	separator string
	flKeyword string
	criterias []string
}

// NewFlCriteriaObject - This function creates the Fl object
func NewFlCriteriaObject() FlCriteria {
	Fl := FlCriteria{separator: ",", flKeyword: "fl", criterias: []string{}}
	return Fl
}

// AddCriteria - This adds the Fl criteria
func (Fl FlCriteria) AddCriteria(FlOption FlOption) FlCriteria {
	Fl.criterias = FlOption.Fields
	return Fl
}

// BuildCriteria - This builds the whole criteria
func (Fl FlCriteria) BuildCriteria() string {
	if len(Fl.criterias) == 0 {
		return ""
	}
	FlStr := strings.Join(Fl.criterias, Fl.separator)
	return Fl.flKeyword + "=" + FlStr
}
