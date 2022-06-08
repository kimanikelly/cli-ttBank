package app

import (
	"fmt"

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

	fmt.Println(nonce)
	fmt.Println(gasPrice)
	fmt.Println(chainId)
	fmt.Println(privateKey)

}
