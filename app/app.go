package app

import (
	"fmt"

	"github.com/kimanikelly/cli-ttBank/contract"
)

func StartApp() {

	ttBank := contract.ContractInstance()

	fmt.Println(ttBank)
}
