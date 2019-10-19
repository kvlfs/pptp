package model

type SendInfo struct {
	Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsUsed   bool   `json:"isUsed"`
}
