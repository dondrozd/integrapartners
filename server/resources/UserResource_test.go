package resources_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"server/resources"
)

var _ = Describe("UserResource", func() {
	It("RegisterNewUserResource", func() {
		resources.RegisterNewUserResource(nil, nil)
	})
	It("thing", func() {
		//get
	})
})
