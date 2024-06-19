package blockchain

import (
	"fmt"
	"testing"
)

func TestRemoveContract(t *testing.T) {
	contracts = GetContracts()
	fmt.Println(GetContracts())
	AddContract("BB")
	AddContract("CC")
	fmt.Println(GetContracts())
	RemoveContract("AA")
	RemoveContract("BB")
	RemoveContract("CC")
	fmt.Println(GetContracts())
}
