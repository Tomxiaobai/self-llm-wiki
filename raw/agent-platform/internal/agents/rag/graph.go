package rag

func BuildGraphPlan() []string {
	return []string{
		"normalize_query",
		"retrieve_internal_documents",
		"return_candidate_evidence",
	}
}
