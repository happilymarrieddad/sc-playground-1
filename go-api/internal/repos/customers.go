package repos

import (
	"api/internal/utils"
	"api/types"
	"time"

	"xorm.io/xorm"
)

//go:generate mockgen -destination=./mocks/Customers.go -package=mock_repos api/internal/repos Customers
type CustomersRepo interface {
	FindOrCreate(name string) (*types.Customer, error)
	Create(cat ...*types.Customer) (err error)
	Delete(ids ...int64) error
	Get(id int64) (*types.Customer, error)
	Find(opts *CustomersFindOpts) (cats []*types.Customer, err error)
}

func NewCustomersRepo(db *xorm.Engine) CustomersRepo {
	return &customersRepo{db: db}
}

type customersRepo struct {
	db *xorm.Engine
}

func (r *customersRepo) FindOrCreate(name string) (*types.Customer, error) {
	existingObj, err := r.Find(&CustomersFindOpts{ByObj: &types.Customer{Name: name}})
	if err != nil {
		return nil, err
	}

	if len(existingObj) > 0 {
		return existingObj[0], nil
	}

	cust := &types.Customer{Name: name}
	if err = r.Create(cust); err != nil {
		return nil, err
	}

	return cust, nil
}

func (r *customersRepo) Create(cats ...*types.Customer) (err error) {
	sesh := r.db.NewSession()
	t := time.Now()

	for _, cat := range cats {
		if err = types.Validate(cat); err != nil {
			return handleRollback(sesh, err)
		}

		cat.CreatedAt = t
		cat.UpdatedAt = nil

		if _, err = r.db.Insert(cat); err != nil {
			return handleRollback(sesh, err)
		}
	}

	return sesh.Commit()
}

func (r *customersRepo) Delete(ids ...int64) error {
	if _, err := r.db.In("id", utils.Int64ArrToInterfaceArr(ids...)...).Delete(&types.Customer{}); err != nil {
		return err
	}

	return nil
}

func (r *customersRepo) Get(id int64) (*types.Customer, error) {
	cat := new(types.Customer)

	if has, err := r.db.ID(id).Get(cat); err != nil {
		return nil, err
	} else if !has {
		return nil, types.NewNotFoundError("unable to find customer by id")
	}

	return cat, nil
}

type CustomersFindOpts struct {
	Limit  int
	Offset int
	ByObj  *types.Customer
	IDs    []int64
}

func (r *customersRepo) Find(opts *CustomersFindOpts) (cats []*types.Customer, err error) {
	if opts == nil {
		opts = &CustomersFindOpts{}
	}

	sesh := r.db.OrderBy("name")

	if opts.Limit > 0 {
		if opts.Offset > 0 {
			sesh = sesh.Limit(opts.Limit, opts.Offset)
		} else {
			sesh = sesh.Limit(opts.Limit)
		}
	}

	if len(opts.IDs) > 0 {
		sesh = sesh.In("id", utils.Int64ArrToInterfaceArr(opts.IDs...)...)
	}

	if opts.ByObj != nil {
		if opts.ByObj.Name != "" {
			sesh = sesh.And("name = ?", opts.ByObj.Name)
		}

		if err = sesh.Find(&cats); err != nil {
			return nil, handleRollback(sesh, err)
		}
	} else {
		if err = sesh.Find(&cats); err != nil {
			return nil, handleRollback(sesh, err)
		}
	}

	return cats, sesh.Commit()
}
