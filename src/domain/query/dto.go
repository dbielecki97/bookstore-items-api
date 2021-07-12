package query

type EsQuery struct {
	Equals []FieldValue `json:"equals,omitempty"`
}
type FieldValue struct {
	Field string      `json:"field,omitempty"`
	Value interface{} `json:"value,omitempty"`
}
