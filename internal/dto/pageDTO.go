package dto

type PageDTO struct {
	URL    string `json:"url"`
	Filter string `json:"filter,omitempty"`
}
