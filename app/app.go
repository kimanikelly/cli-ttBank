package app

import (
	"fmt"

	"github.com/kimanikelly/cli-ttBank/contract"
	"github.com/kimanikelly/cli-ttBank/signer"
)

func StartApp() {

	// Returns the TTBank.sol contract instance
	ttBank := contract.ContractInstance()

	fmt.Println(ttBank)

	signerAddress, privateKey := signer.Wallet()

	fmt.Println(signerAddress, privateKey)

}
