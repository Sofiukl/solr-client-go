package solrqry

// CriteriaBuilder - This is the contract for the Criteria Builder
type CriteriaBuilder interface {
	BuildCriteria() string
}

// CriteriaBuilderEngine - This is the CriteriaBuilder
type CriteriaBuilderEngine struct {
	criteriaBuilders []CriteriaBuilder
}

// NewCriteriaBuilderEngine - This is return the instance of CriteriaBuilder
func NewCriteriaBuilderEngine(builders []CriteriaBuilder) CriteriaBuilderEngine {

	cbe := CriteriaBuilderEngine{criteriaBuilders: builders}
	return cbe

}

// Build - This builds the whole criteria
func (cbe CriteriaBuilderEngine) Build() string {
	URLParam := ""
	for _, builder := range cbe.criteriaBuilders {
		criteria := builder.BuildCriteria()
		if criteria != "" && URLParam == "" {
			URLParam = URLParam + criteria
		} else if criteria != "" {
			URLParam = URLParam + "&" + criteria
		}
	}
	if URLParam != "" {
		URLParam = "?" + URLParam
	}
	return URLParam
}
