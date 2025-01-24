package types

type Pagination struct {
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}

type PaginatedResponse struct {
	Pagination Pagination  `json:"pagination"`
	Data       interface{} `json:"data"`
}
