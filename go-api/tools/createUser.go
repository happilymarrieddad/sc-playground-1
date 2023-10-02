package main

import (
	"api/internal/postgres"
	"api/internal/repos"
	"api/internal/utils"
	"api/types"
	"fmt"
)

func main() {
	db, err := postgres.NewDB(
		utils.GetEnv("SYMBIOSIS_DB_USER", "postgres"),
		utils.GetEnv("SYMBIOSIS_DB_PASS", "postgres"),
		utils.GetEnv("SYMBIOSIS_DB_HOST", "localhost"),
		utils.GetEnv("SYMBIOSIS_DB_PORT", "5433"),
		utils.GetEnv("SYMBIOSIS_DB_DATABASE", "symbiosis"),
	)
	if err != nil {
		panic(err)
	}

	gr, err := repos.NewGlobalRepo(db)
	if err != nil {
		panic(err)
	}

	newCust, err := gr.Customers().FindOrCreate("Admin")
	if err != nil {
		panic(err)
	}

	newUser := types.NewUser(
		"Nick",
		"Kotenberg",
		"nick@mail.com",
		"1234",
		newCust.ID,
	)

	if newUser, err = gr.Users().FindOrCreate(newUser); err != nil {
		panic(err)
	}

	fmt.Printf("New User ensured '%d' '%s", newUser.ID, newUser.Email)
}
