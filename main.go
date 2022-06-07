package main

import (
	"fmt"

	"github.com/kimanikelly/cli-ttBank/contract"
)

func main() {

	ttBank := contract.ContractInstance()

	fmt.Println(ttBank)

}
