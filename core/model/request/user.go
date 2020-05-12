package request

type RegisterStruct struct {
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Head     string `json:"head"`
}

type LoginStruct struct {
	Phone    string `json:"phone"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}
