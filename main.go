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
	fmt.Println(w1.PrivateKeyStr())
	fmt.Println(w1.PublicKeyStr())
}
