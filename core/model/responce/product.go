package responce



type ProductResult struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Title  string  `json:"title"`
	Img    string  `json:"img"`
	Detail string  `json:"detail"`
	Price  float32 `json:"price"`
	Stock  int     `json:"stock"`
}
