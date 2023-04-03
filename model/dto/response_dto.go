package dto

type ResponseStatus struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Status ResponseStatus `json:"status"`
	Data   interface{}    `json:"data,omitempty"`
}

type PagedResponse struct {
	Status ResponseStatus `json:"status"`
	Data   []interface{}  `json:"data,omitempty"`
	Paging Paging         `json:"paging,omitempty"`
}
