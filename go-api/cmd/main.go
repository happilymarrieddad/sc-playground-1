package main

import (
	apipkg "api/internal/api"
	"api/internal/postgres"
	"api/internal/repos"
	"api/internal/utils"
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

	apipkg.Start(gr, 8000)
}
