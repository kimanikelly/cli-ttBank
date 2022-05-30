package utils

import (
	"context"
	"log"
	"math/big"

	"github.com/kimanikelly/cli-ttBank/client"
)

func ChainID() *big.Int {
	chainId, err := client.ClientConnection().ChainID(context.Background())

	if err != nil {
		log.Fatalf("Failed to chainID %v", err)
	}

	log.Println(chainId)

	return chainId
}
