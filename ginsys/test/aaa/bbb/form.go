package bbb

type Demo001Form struct {
	Name string `json:"name" binding:"required,phone=r8"`
}
