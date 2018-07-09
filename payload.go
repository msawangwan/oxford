package oxford

type GrammaticalFeature struct {
	Text string `json:"text,omitempty"`
	Type string `json:"type,omitempty"`
}

type InflectionInfo struct {
	ID   string `json:"id,omitempty"`
	Text string `json:"text,omitempty"`
}

type LexicalEntry struct {
	GrammaticalFeatures GrammaticalFeature `json:"grammatical_features,omitempty"`
	InflectionOf        InflectionInfo     `json:"inflection_of,omitempty"`
	Language            string             `json:"language,omitempty"`
	LexicalCategory     string             `json:"lexical_category,omitempty"`
	Text                string             `json:"text,omitempty"`
}

type Result struct {
	ID             string         `json:"id,omitempty"`
	Language       string         `json:"language,omitempty"`
	LexicalEntries []LexicalEntry `json:"lexical_entries,omitempty"`
	Type           string         `json:"type,omitempty"`
	Word           string         `json:"word,omitempty"`
}

type WordExistsResponse struct {
	Metadata struct{} `json:"metadata,omitempty"`
	Results  []Result `json:"results,omitempty"`
}
