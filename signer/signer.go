package signer

import (
	"crypto/ecdsa"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kimanikelly/cli-ttBank/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Wallet() (common.Address, *ecdsa.PrivateKey) {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Failed to load the .env file %v", err)
	}

	// The privateKey for the connected signer
	var privateKey *ecdsa.PrivateKey

	// The public address for the connected signer
	var signerAddress common.Address

	// Checks if the 'NETWORK' env variable is localhost
	if os.Getenv("NETWORK") == "localhost" {

		// Returns the parsed ganache private key
		privateKey, err = crypto.HexToECDSA(utils.PrivateKeys()[0])

		// If err does not equal nil(zero value) throw an error
		if err != nil {
			log.Fatalf("Failed to parse the private key %v", err)
		}

		localPublicKey := privateKey.Public()

		publicKeyECDSA, ok := localPublicKey.(*ecdsa.PublicKey)

		if !ok {
			log.Fatal("Failed assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		signerAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	}

	// Checks if the 'NETWORK' env variable is rinkeby
	if os.Getenv("NETWORK") == "rinkeby" {

		privateKey, err = crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))

		if err != nil {
			log.Fatalf("Failed to parse the private key %v", err)
		}

		publicKey := privateKey.Public()

		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

		if !ok {
			log.Fatal("Failed assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		signerAddress = crypto.PubkeyToAddress(*publicKeyECDSA)

	}

	return signerAddress, privateKey
}
