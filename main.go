package main

import (
	"fmt"
	"goblockchain/block"
	"goblockchain/wallet"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	bc := block.NewBlockchain(walletM.BlockchainAddress())
	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)
	isAdded := bc.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())

	fmt.Println("Added  ", isAdded)
	bc.Mining()
	bc.Print()
	fmt.Println(bc.BalanceOf(walletM.BlockchainAddress()))
	fmt.Println(bc.BalanceOf(walletA.BlockchainAddress()))
	fmt.Println(bc.BalanceOf(walletB.BlockchainAddress()))
}
