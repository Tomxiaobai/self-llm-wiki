package gateway

func BuildGraphPlan() []string {
	return []string{
		"normalize_query",
		"call_rag_agent",
		"judge_need_search",
		"call_search_agent",
		"call_report_agent",
		"finalize_answer",
	}
}
