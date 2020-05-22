package request

type PageInfo struct {
	Page     int `json:"page" from:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
}
