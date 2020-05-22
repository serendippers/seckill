package responce

type PageResult struct {
	List     interface{} `json:"list"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Total    int           `json:"total"`
}
