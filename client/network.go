package client

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Handles the JsonRpcProvider URL based on the "NETWORK" env variable and returns the connection URL
func LoadNetwork() string {

	/*
		The value of 'network' is either http://localhost:8545 or https://rinkeby.infura.io/v3/
		based on the 'NETWORK' env variable.
	*/
	var network string

	// Loads the environment variables from the .env file
	err := godotenv.Load()

	// If err does not equal nil(zero value) throw an error
	if err != nil {
		log.Fatalf("Failed to load the environment variables %v", err)
	}

	// Checks if the 'NETWORK' env variable is localhost
	if os.Getenv("NETWORK") == "localhost" {

		// Sets the value of 'network' to a localhost connection
		network = "http://localhost:8545"
	}

	// Checks if the 'NETWORK' env variable is rinkeby
	if os.Getenv("NETWORK") == "rinkeby" {

		// Sets the value of 'network' to a rinkeby connection
		network = os.Getenv("RINKEBY_URL")
	}

	// Returns the value of 'network'
	return network
}
