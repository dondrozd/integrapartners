package daos_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"server/daos"
)

var _ = Describe("User", func() {
	var testuser daos.User
	var longText = "85qbRpeeBKNm6J0Q0j4G9W9opda4aCrMy0bJ0VC65yg0PftUJiWiJTeLQTszOlTUYcC9e8YDHf1kqMHxEVgJ24ShAafnafCPIzPjfpa9VjQ7AlTjQxaamRYksjLiXOA9VzQiAr68pc2CN87FJebAPvwSANja9TLmbRgmBr5N6kmBTRwtYlVaZLkHUaHtA88AQM7u9GKWQ4J2VSytKwTMNPgBJoqJdIlLqzllqPHbPwkRPnbvqEYqv8f0OMA85BNe"

	BeforeEach(func() {
		testuser = daos.User{
			ID:        "1",
			FirstName: "fn",
			LastName:  "ln",
			Email:     "em",
			UserName:  "un",
			Status:    "A",
		}
	})
	Context("validation", func() {
		It("all fields should be valid", func() {
			Expect(testuser.IsValid()).Should(BeTrue())
		})
		It("empty first name not valid", func() {
			testuser.FirstName = ""
			Expect(testuser.IsValid()).Should(BeFalse())
		})
		It("long first name not valid", func() {
			testuser.FirstName = longText
			Expect(testuser.IsValid()).Should(BeFalse())
		})
		It("empty last name not valid", func() {
			testuser.LastName = ""
			Expect(testuser.IsValid()).Should(BeFalse())
		})
		It("long last name not valid", func() {
			testuser.LastName = longText
			Expect(testuser.IsValid()).Should(BeFalse())
		})
		It("empty email not valid", func() {
			testuser.Email = ""
			Expect(testuser.IsValid()).Should(BeFalse())
		})
		It("long email not valid", func() {
			testuser.Email = longText
			Expect(testuser.IsValid()).Should(BeFalse())
		})
		It("empty user name not valid", func() {
			testuser.UserName = ""
			Expect(testuser.IsValid()).Should(BeFalse())
		})
		It("long user name not valid", func() {
			testuser.UserName = longText
			Expect(testuser.IsValid()).Should(BeFalse())
		})
		It("status of I is valid", func() {
			testuser.Status = "I"
			Expect(testuser.IsValid()).Should(BeTrue())
		})
		It("status of A is valid", func() {
			testuser.Status = "A"
			Expect(testuser.IsValid()).Should(BeTrue())
		})
		It("status of T is valid", func() {
			testuser.Status = "T"
			Expect(testuser.IsValid()).Should(BeTrue())
		})
		It("other status are not valid", func() {
			testuser.Status = "X"
			Expect(testuser.IsValid()).Should(BeFalse())
		})
	})

})
