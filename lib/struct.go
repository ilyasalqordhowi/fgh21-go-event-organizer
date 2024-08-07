package lib

type Message struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Results any    `json:"results,omitempty"`
}