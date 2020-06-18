package request

type OrderInfo struct {
	UserId    int64 `json:"userId" form:"userId"`
	ProductId int64 `json:"productId" form:"productId"`
	//Count     int   `json:"count" form:"count"`
}
