package entities

import "time"

// String entity
type String struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Depth     float64    `json:"depth"`
	CaseID    string     `json:"case_id"`
	CreatedAt time.Time  `json:"created_at"`
	Sections  []*Section `json:"sections"`
}

// SectionType entity
type SectionType struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// SectionAttribute entity
type SectionAttribute struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Unit          string    `json:"unit"`
	ValueTypeID   string    `json:"value_type_id"`
	SectionTypeID string    `json:"section_type_id"`
	CreatedAt     time.Time `json:"created_at"`
}

// Section entity
type Section struct {
	ID            string          `json:"id"`
	StringID      string          `json:"string_id"`
	SectionTypeID string          `json:"section_type_id"`
	CreatedAt     time.Time       `json:"created_at"`
	Values        []*SectionValue `json:"values"`
}

// SectionValue entity
type SectionValue struct {
	ID          string    `json:"id"`
	SectionID   string    `json:"section_id"`
	AttributeID string    `json:"attribute_id"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
}

// SectionValueType entity
type SectionValueType struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// SectionAttributeTranslation entity
type SectionAttributeTranslation struct {
	ID          string `json:"id"`
	AttributeID string `json:"attribute_id"`
	LanguageID  string `json:"language_id"`
	Name        string `json:"name"`
	Unit        string `json:"unit"`
}
