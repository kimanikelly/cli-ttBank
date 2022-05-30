package utils

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/kimanikelly/cli-ttBank/client"
)

func Nonce(address common.Address) uint64 {

	// Returns the number of transactions the connected signer sent overtime
	// The nonce can only be used once per transaction and increments by 1 per transaction
	nonce, err := client.ClientConnection().PendingNonceAt(context.Background(), address)

	if err != nil {
		log.Fatalf("Failed to return the nonce %v", err)
	}

	return nonce
}

func GasPrice() *big.Int {

	// Estimates the cost needed to perform a transaction
	gasPrice, err := client.ClientConnection().SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatalf("Failed to return the gasPrice %v", err)
	}

	return gasPrice

}

func ChainID() *big.Int {
	chainId, err := client.ClientConnection().ChainID(context.Background())

	if err != nil {
		log.Fatalf("Failed to return the chainID %v", err)
	}

	return chainId
}
