package repos_test

import (
	. "api/internal/repos"
	"api/types"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UsersRepo test", func() {
	var repo UsersRepo

	BeforeEach(func() {
		clearDatabase("users")

		repo = gr.Users()
		Expect(repo).NotTo(BeNil())
	})

	Context("Create", func() {
		It("should successfully create a user", func() {
			newUsr := types.NewUser("Nick", "Kotenberg", "nick@mail.com", "1234", newCus1.ID)
			Expect(repo.Create(newUsr)).To(Succeed())

			Expect(newUsr.ID).NotTo(BeEquivalentTo(0))
			Expect(newUsr.FirstName).To(Equal("Nick"))
			Expect(newUsr.LastName).To(Equal("Kotenberg"))
			Expect(newUsr.Email).To(Equal("nick@mail.com"))
			Expect(newUsr.PasswordMatches("1234")).To(BeTrue())
			cy, cm, cd := newUsr.CreatedAt.Date()
			ty, tm, td := time.Now().Date()
			Expect(cy).To(Equal(ty))
			Expect(cm).To(Equal(tm))
			Expect(cd).To(Equal(td))
			Expect(newUsr.UpdatedAt).To(BeNil())
		})
	})

	Context("GetByEmail", func() {
		var usr *types.User

		BeforeEach(func() {
			usr = types.NewUser("Nick", "Kotenberg", "nick@mail.com", "1234", newCus1.ID)
			Expect(repo.Create(usr)).To(Succeed())
			Expect(usr.ID).NotTo(BeEquivalentTo(0))
		})

		It("should not get an invalid user", func() {
			usr, err := repo.GetByEmail("garbage")
			Expect(types.IsNotFoundError(err)).To(BeTrue())
			Expect(usr).To(BeNil())
		})

		It("should successfully get the correct user", func() {
			existingUsr, err := repo.GetByEmail(usr.Email)
			Expect(err).To(BeNil())
			Expect(existingUsr.ID).To(Equal(usr.ID))
			Expect(existingUsr.FirstName).To(Equal(usr.FirstName))
			Expect(existingUsr.LastName).To(Equal(usr.LastName))
			Expect(existingUsr.Email).To(Equal(usr.Email))
		})
	})

	Context("GetByID", func() {
		var usr *types.User

		BeforeEach(func() {
			usr = types.NewUser("Nick", "Kotenberg", "nick@mail.com", "1234", newCus1.ID)
			Expect(repo.Create(usr)).To(Succeed())
			Expect(usr.ID).NotTo(BeEquivalentTo(0))
		})

		It("should not get an invalid user", func() {
			usr, err := repo.GetByID(999999)
			Expect(types.IsNotFoundError(err)).To(BeTrue())
			Expect(usr).To(BeNil())
		})

		It("should successfully get the correct user", func() {
			existingUsr, err := repo.GetByID(usr.ID)
			Expect(err).To(BeNil())
			Expect(existingUsr.ID).To(Equal(usr.ID))
			Expect(existingUsr.FirstName).To(Equal(usr.FirstName))
			Expect(existingUsr.LastName).To(Equal(usr.LastName))
			Expect(existingUsr.Email).To(Equal(usr.Email))
		})
	})
})
