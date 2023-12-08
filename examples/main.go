package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	vaultReader "github.com/lee-aaron/gmx-go-sdk/VaultReader"
)

func main() {

	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	reader, err := vaultReader.NewVaultReader(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reader)
}
