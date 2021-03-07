package model

type User struct {
	ID         string  `json:"id"         query:"id"`
	FirstName  string  `json:"firstName"  query:"firstName"`
	LastName   string  `json:"lastName"   query:"lastName"`
	Email      string  `json:"email"      query:"email"`
	UserName   string  `json:"userName"   query:"userName"`
	Status     string  `json:"status"     query:"status"`
	Department *string `json:"department" query:"department"`
}
