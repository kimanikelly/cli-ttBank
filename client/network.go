package client

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadNetwork() string {

	var network string

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed to load the environment variables %v", err)
	}

	getNetwork := os.Getenv("NETWORK")

	if getNetwork == "localhost" {
		network = "http://localhost:8545"
	} else if getNetwork == "rinkeby" {
		network = os.Getenv("RINKEBY_URL")
	}

	return network
}
