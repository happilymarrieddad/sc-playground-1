package repos_test

import (
	"api/types"
	"fmt"
	"testing"
	"time"

	"api/internal/postgres"
	. "api/internal/repos"
	"api/internal/utils"

	_ "github.com/jackc/pgx/v5/stdlib"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"xorm.io/xorm"
)

var (
	db      *xorm.Engine
	gr      GlobalRepo
	newCus1 *types.Customer
	newCus2 *types.Customer
)

var _ = BeforeSuite(func() {
	defer GinkgoRecover()

	var (
		err error
	)

	db, err = postgres.NewDB(
		utils.GetEnv("SYMBIOSIS_DB_USER_TEST", "postgres"),
		utils.GetEnv("SYMBIOSIS_DB_PASS_TEST", "postgres"),
		utils.GetEnv("SYMBIOSIS_DB_HOST_TEST", "localhost"),
		utils.GetEnv("SYMBIOSIS_DB_PORT_TEST", "5433"),
		utils.GetEnv("SYMBIOSIS_DB_DATABASE_TEST", "symbiosis"),
	)
	Expect(err).To(BeNil())

	gr, err = NewGlobalRepo(db)
	Expect(err).To(BeNil())

	clearDatabase("customers")
	newCus1 = &types.Customer{
		Name: "Customer TEST",
	}
	newCus2 = &types.Customer{
		Name: "Customer TEST 2",
	}

	Expect(gr.Customers().Create(newCus1)).To(Succeed())
	Expect(newCus1.ID).To(BeNumerically(">", 0))
	cy, cm, cd := newCus1.CreatedAt.Date()
	ty, tm, td := time.Now().Date()
	Expect(cy).To(Equal(ty))
	Expect(cm).To(Equal(tm))
	Expect(cd).To(Equal(td))
	Expect(newCus1.UpdatedAt).To(BeNil())

	Expect(gr.Customers().Create(newCus2)).To(Succeed())
	Expect(newCus2.ID).To(BeNumerically(">", 0))
	cy, cm, cd = newCus2.CreatedAt.Date()
	ty, tm, td = time.Now().Date()
	Expect(cy).To(Equal(ty))
	Expect(cm).To(Equal(tm))
	Expect(cd).To(Equal(td))
	Expect(newCus2.UpdatedAt).To(BeNil())
})

func clearDatabase(tables ...string) {
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", table))
		Expect(err).To(Succeed())
	}
}

func TestRepos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repos Suite")
}
