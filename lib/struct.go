package lib

type Message struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results,omitempty"`
}
