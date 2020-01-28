package handler

type Word struct {
	Category string   `json:"category"`
	Synonyms []string `json:"synonyms"`
}

type Entry struct {
	Word    string `json:"word"`
	Entries []Word `json:"entries"`
}
