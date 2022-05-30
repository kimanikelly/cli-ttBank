package main

import (
	"fmt"

	"github.com/kimanikelly/cli-ttBank/signer"
)

func main() {

	address, privateKey := signer.Wallet()

	fmt.Println(address, privateKey)
}
