package schema

type Citation struct {
	SourceType string `json:"source_type"`
	SourceURI  string `json:"source_uri"`
	ChunkID    string `json:"chunk_id"`
	SpanText   string `json:"span_text"`
}
