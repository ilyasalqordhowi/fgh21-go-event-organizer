package lib

type Message struct {
	Success     bool   `json:"success"`
	Message     string `json:"message"`
	ResultsInfo any    `json:"resultsInfo,omitempty"`
	Results     any    `json:"results,omitempty"`
}
type TotalInfo struct {
	TotalData int `json:"totalData"`
	TotalPage int `json:"totalPage"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Next      int `json:"next"`
	Prev      int `json:"prev"`
}