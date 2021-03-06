package resources_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"server/resources"
)

var _ = Describe("Donresource", func() {
	Context("test", func() {
		It("returns", func() {
			Expect(resources.DoSomething("value")).To(Equal("hello value"))
		})
	})
})
