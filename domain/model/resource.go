package model

type Resource struct {
	Name       string       `json:"name"`
	Attributes AttributeSet `json:"attributes"`
}
