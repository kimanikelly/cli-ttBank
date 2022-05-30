package client

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func ClientConnection() *ethclient.Client {

	// Connects to the provider and creates the client instance
	client, err := ethclient.Dial(LoadNetwork())

	// If err does not equal nil(zero value) throw an error
	if err != nil {

		log.Fatalf("Failed to connect to the Ethereum client %v", err)
	}

	// Returns the client instance
	return client
}
