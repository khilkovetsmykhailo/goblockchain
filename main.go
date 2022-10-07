package main

import (
	"fmt"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	w1 := wallet.NewWallet()
	t := wallet.NewTransaction(w1.PrivateKey(), w1.PublicKey(), w1.BlockchainAddress(), "B", 1.0)
	fmt.Println(t.GenerateSignature().String())
}
