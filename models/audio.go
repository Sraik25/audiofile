package models

type Audio struct {
	Id       string   `json:"id,omitempty"`
	Path     string   `json:"path,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
	Status   string   `json:"status,omitempty"`
	Error    []error  `json:"error,omitempty"`
}
