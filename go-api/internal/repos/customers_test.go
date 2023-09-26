package repos_test

import (
	. "api/internal/repos"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CustomersRepo test", func() {
	var (
		repo CustomersRepo
	)

	BeforeEach(func() {
		repo = gr.Customers()
		Expect(repo).NotTo(BeNil())
	})

	Context("Get", func() {
		It("should successfully get a customer", func() {
			existingCustomer, err := repo.Get(newCus1.ID)
			Expect(err).To(BeNil())
			Expect(existingCustomer).NotTo(BeNil())
			Expect(existingCustomer.Name).To(Equal(newCus1.Name))
		})
	})

	Context("Find", func() {
		It("should successfully get the passed in opts by id", func() {
			existingCusts, err := repo.Find(&CustomersFindOpts{IDs: []int64{newCus1.ID}})
			Expect(err).To(BeNil())
			Expect(existingCusts).To(HaveLen(1))

			for _, cat := range existingCusts {
				switch cat.ID {
				case newCus1.ID:
					Expect(cat.Name).To(Equal(newCus1.Name))
				default:
					Expect(false).To(BeTrue(), "should never get here")
				}
			}
		})
	})
})
