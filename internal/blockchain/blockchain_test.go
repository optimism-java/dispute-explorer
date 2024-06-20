package blockchain

import (
	"fmt"
	"testing"
)

func TestRemoveContract(t *testing.T) {
	contracts = GetContracts()
	fmt.Println(GetContracts())
	AddContract("0x05F9613aDB30026FFd634f38e5C4dFd30a197Ba1")
	AddContract("CC")
	fmt.Println(GetContracts())
	RemoveContract("0x05f9613adB30026FFd634f38e5C4dFd30a197ba1")
	RemoveContract("BB")
	RemoveContract("CC")
	fmt.Println(GetContracts())
}
