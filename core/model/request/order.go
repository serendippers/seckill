package request

type OrderInfo struct {
	UserId         int64 `json:"userId" form:"userId"`
	ProductId      int64 `json:"productId" form:"productId"`
	DeliveryAddrId int64 `json:"deliveryAddrId" form:"deliveryAddrId"`
}
