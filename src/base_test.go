package Base_test

import (
	Base "github.com/golang-cop/base/src"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("Base", func() {
	var base Base.Interface
	var methods []string = []string{`Kind`, `Methods`, `RespondTo`}
	ginkgo.BeforeEach(func() {
		base = Base.New()
	})
	ginkgo.It("method Kind() can return the Interface kind", func() {
		gomega.Expect(
			base.Kind(),
		).To(gomega.BeIdenticalTo(`Base.Interface`))
	})
	ginkgo.It("method Methods() return the object methods", func() {
		gomega.Expect(
			base.Methods(),
		).To(gomega.BeComparableTo(methods))
	})
	ginkgo.It(`method RespondTo() return false as this is not a Base.Interface method`, func() {
		gomega.Expect(
			base.RespondTo(`unexistent`),
		).To(gomega.BeFalseBecause(`this is not a Base.Interface method`))
	})
	ginkgo.It(`method RespondTo() return true as this is a Base.Interface method`, func() {
		gomega.Expect(
			base.RespondTo(`RespondTo`),
		).To(gomega.BeTrueBecause(`this is a Base.Interface method`))
	})
})
