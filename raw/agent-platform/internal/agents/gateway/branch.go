package gateway

import (
	"strings"

	"agent-platform/internal/shared/schema"
)

func NeedExternalSearch(query string, internal []schema.Evidence) bool {
	if len(internal) < 2 {
		return true
	}

	lowerQuery := strings.ToLower(query)
	keywords := []string{"latest", "current", "recent", "web", "internet", "网上", "最新"}
	for _, keyword := range keywords {
		if strings.Contains(lowerQuery, keyword) {
			return true
		}
	}
	return false
}
