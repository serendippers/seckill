package request

type RegisterStruct struct {
	Nickname      string    `json:"nickname"`
	Password      string    `json:"password"`
	Salt          string    `json:"salt"`
	Head          string    `json:"head"`
}