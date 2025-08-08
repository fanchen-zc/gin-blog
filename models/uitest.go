package models

type UIAnalysis struct {
	Problems    []Problem    `json:"problems"`
	Suggestions []Suggestion `json:"suggestions"`
}

type Problem struct {
	Element     string `json:"element"`
	Description string `json:"description"`
	Suggestion  string `json:"suggestion"`
}

type Suggestion struct {
	Area        string `json:"area"`
	Description string `json:"description"`
	Suggestion  string `json:"suggestion"`
}
