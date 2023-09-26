package repos

import (
	"api/types"
	"fmt"
	"time"

	"xorm.io/xorm"
)

//go:generate mockgen -destination=./mocks/Users.go -package=mock_repos api/repos Users
type UsersRepo interface {
	Create(*types.User) error
	GetByID(id int64) (*types.User, error)
	GetByEmail(email string) (*types.User, error)
}

func NewUsersRepo(db *xorm.Engine) UsersRepo {
	return &usersRepo{db: db}
}

type usersRepo struct {
	db *xorm.Engine
}

func (r *usersRepo) GetByEmail(email string) (*types.User, error) {
	user := &types.User{Email: email}

	if exists, err := r.db.Get(user); err != nil {
		return nil, err
	} else if !exists {
		return nil, types.NewNotFoundError(fmt.Sprintf("user with email '%s' not found", email))
	}

	return user, nil
}

func (r *usersRepo) GetByID(id int64) (*types.User, error) {
	user := &types.User{}
	if exists, err := r.db.ID(id).Get(user); err != nil {
		return nil, err
	} else if !exists {
		return nil, types.NewNotFoundError(fmt.Sprintf("user with id '%d' not found", id))
	}

	return user, nil
}

func (r *usersRepo) Create(newUsr *types.User) error {
	if err := types.Validate(newUsr); err != nil {
		return err
	}

	t := time.Now()
	newUsr.CreatedAt = t
	newUsr.UpdatedAt = nil
	newUsr.UserType = types.StandardUserType

	if _, err := r.db.Insert(newUsr); err != nil {
		return err
	}

	return nil
}
