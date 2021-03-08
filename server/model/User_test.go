package model_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"server/model"
)

var _ = Describe("User", func() {
	var longText = "85qbRpeeBKNm6J0Q0j4G9W9opda4aCrMy0bJ0VC65yg0PftUJiWiJTeLQTszOlTUYcC9e8YDHf1kqMHxEVgJ24ShAafnafCPIzPjfpa9VjQ7AlTjQxaamRYksjLiXOA9VzQiAr68pc2CN87FJebAPvwSANja9TLmbRgmBr5N6kmBTRwtYlVaZLkHUaHtA88AQM7u9GKWQ4J2VSytKwTMNPgBJoqJdIlLqzllqPHbPwkRPnbvqEYqv8f0OMA85BNe"
	var validJsonData = []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "A", "department" : "dp" }`)

	It("unmarshals from expected json", func() {
		var user model.User
		err := json.Unmarshal(validJsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.FirstName).To(Equal("fn"))
		Expect(user.LastName).To(Equal("ln"))
		Expect(user.Email).To(Equal("em"))
		Expect(user.UserName).To(Equal("un"))
		Expect(user.Status).To(Equal("A"))
		Expect(*user.Department).To(Equal("dp"))
	})

	It("can have nil Department", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "A" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.FirstName).To(Equal("fn"))
		Expect(user.LastName).To(Equal("ln"))
		Expect(user.Email).To(Equal("em"))
		Expect(user.UserName).To(Equal("un"))
		Expect(user.Department).To(BeNil())
	})

	It("validation passes in happy path", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "A" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeTrue())
	})

	It("validation fails when first name is empty", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "", "lastName": "ln", "email": "em", "userName": "un", "status": "A" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when first name is too long", func() {
		var user model.User
		err := json.Unmarshal(validJsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		user.FirstName = longText
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when last name is empty", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "", "email": "em", "userName": "un", "status": "A" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when last name is too long", func() {
		var user model.User
		err := json.Unmarshal(validJsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		user.LastName = longText
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when email is empty", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "", "userName": "un", "status": "A" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when email is too long", func() {
		var user model.User
		err := json.Unmarshal(validJsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		user.LastName = longText
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when user name is empty", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "", "status": "A" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when user name is too long", func() {
		var user model.User
		err := json.Unmarshal(validJsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		user.UserName = longText
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when status is empty", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation passes when status is A", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "A" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeTrue())
	})

	It("validation passes when status is I", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "I" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeTrue())
	})

	It("validation passes when status is T", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "T" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeTrue())
	})

	It("validation passes when status is not A / I / T", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "X" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation fails when department is empty", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "T", "department" : "" }`)
		err := json.Unmarshal(jsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		Expect(user.IsValid()).To(BeFalse())
	})

	It("validation passes when department is empty", func() {
		var user model.User
		err := json.Unmarshal(validJsonData, &user)
		if err != nil {
			Fail("couldnt unmarshal json to user")
		}
		user.Department = &longText
		Expect(user.IsValid()).To(BeFalse())
	})

})
