package request

type OrderInfo struct {
	UserId    int64 `json:"userid" form:"userId"`
	ProductId int64 `json:"productId" form:"productId"`
	Count     int   `json:"count" form:"count"`
}
