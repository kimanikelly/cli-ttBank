package app

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/kimanikelly/cli-ttBank/contract"
	"github.com/kimanikelly/cli-ttBank/signer"
	"github.com/kimanikelly/cli-ttBank/utils"
)

func StartApp() {

	// Returns the TTBank.sol contract instance
	ttBank := contract.ContractInstance()

	fmt.Println(ttBank)

	signerAddress, privateKey := signer.Wallet()

	nonce := utils.Nonce(signerAddress)

	gasPrice := utils.GasPrice()

	chainId := utils.ChainID()

	// Binds the connected signer to the transaction options
	auth, authErr := bind.NewKeyedTransactorWithChainID(privateKey, chainId)

	// If authErr does not equal nil (zero value) throw an error
	if authErr != nil {
		log.Fatalf("Failed to build the NewKeyedTransactorWithChainID %v", authErr)
	}

	fmt.Println(nonce)
	fmt.Println(gasPrice)
	fmt.Println(chainId)
	fmt.Println(privateKey)
	fmt.Println(auth)

}
