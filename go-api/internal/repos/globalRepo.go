package repos

import (
	"sync"

	"xorm.io/xorm"
)

var singular GlobalRepo

//go:generate mockgen -destination=./mocks/GlobalRepo.go -package=mock_repos api/internal/repos GlobalRepo
type GlobalRepo interface {
	Customers() CustomersRepo
	Users() UsersRepo
}

func NewGlobalRepo(db *xorm.Engine) (GlobalRepo, error) {
	if singular == nil {
		singular = &globalRepo{
			db:    db,
			mutex: &sync.RWMutex{},
			repos: make(map[string]interface{}),
		}
	}

	return singular, nil
}

type globalRepo struct {
	db    *xorm.Engine
	repos map[string]interface{}
	mutex *sync.RWMutex
}

func (gr *globalRepo) factory(key string, fn func(db *xorm.Engine) interface{}) interface{} {
	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	val, exists := gr.repos[key]
	if exists {
		return val
	}

	newFac := fn(gr.db)
	gr.repos[key] = newFac

	return newFac
}

func (gr *globalRepo) Customers() CustomersRepo {
	return gr.factory("Customers", func(db *xorm.Engine) interface{} { return NewCustomersRepo(db) }).(CustomersRepo)
}

func (gr *globalRepo) Users() UsersRepo {
	return gr.factory("Users", func(db *xorm.Engine) interface{} { return NewUsersRepo(db) }).(UsersRepo)
}
