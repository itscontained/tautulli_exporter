package main

const (
	TautulliURL string
	APIKey string
)
type library struct {
	Count       int    `json:"count"`
	ParentCount int    `json:"parent_count"`
	SectionId   int    `json:"section_id"`
	Agent       string `json:"agent"`
	SectionName string `json:"section_name"`
	SectionType string `json:"section_type"`
    ChildCount  int    `json:"child_count"`
}
