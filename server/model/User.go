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

func (u *User) IsValid() bool {
	return validateStandardSize(u.FirstName) &&
		validateStandardSize(u.LastName) &&
		validateStandardSize(u.Email) &&
		validateStandardSize(u.UserName) &&
		validateStatus(u.Status) &&
		validateNullableWithSize(u.Department)
}

func validateStandardSize(value string) bool {
	size := len(value)
	return size > 0 && size <= 255
}

func validateNullableWithSize(value *string) bool {
	if value != nil {
		size := len(*value)
		return size > 0 && size <= 255
	}
	return true
}

func validateStatus(value string) bool {
	return value == "I" || value == "A" || value == "T"
}
