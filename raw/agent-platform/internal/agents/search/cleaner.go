package search

import "strings"

type Cleaner interface {
	Clean(input string) string
}

type BasicCleaner struct{}

func (BasicCleaner) Clean(input string) string {
	return strings.TrimSpace(input)
}
