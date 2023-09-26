package users

import (
	"api/internal/repos"
	"api/types"
	"encoding/json"
)

func create(gr repos.GlobalRepo, sessionUser *types.User, data []byte) (interface{}, error) {
	type CreateUser struct {
		FirstName string `validate:"required" json:"firstName"`
		LastName  string `validate:"required" json:"lastName"`
		Email     string `validate:"required" json:"email"`
		// TODO: make sure these two match
		Password        string `validate:"required" json:"password"`
		ConfirmPassword string `validate:"required" json:"confirmPassword"`
	}

	req := &CreateUser{}
	if err := json.Unmarshal(data, req); err != nil {
		return nil, err
	}

	if err := types.Validate(req); err != nil {
		return nil, err
	}

	// TODO: allow admin users to pass in the customer_id
	newUsr := types.NewUser(
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
		sessionUser.CustomerID,
	)

	if err := gr.Users().Create(newUsr); err != nil {
		return nil, err
	}

	return newUsr, nil
}
