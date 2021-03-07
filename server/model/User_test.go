package model_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"server/model"
)

var _ = Describe("User", func() {
	It("unmarshals from expected json", func() {
		var user model.User
		jsonData := []byte(`{ "id": "1", "firstName": "fn", "lastName": "ln", "email": "em", "userName": "un", "status": "A", "department": "dp" }`)
		err := json.Unmarshal(jsonData, &user)
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

})
