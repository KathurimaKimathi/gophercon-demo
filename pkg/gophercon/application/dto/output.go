package dto

// OKResponse is used to return OK response in http calls
type OKResponse struct {
	Status   string      `json:"status,omitempty"`
	Response interface{} `json:"response,omitempty"`
}
