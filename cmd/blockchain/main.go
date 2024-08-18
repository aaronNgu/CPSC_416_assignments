package main

import (
	"fmt"
	"sandbox/blockchain"
)

func main() {
	difficulty := 1
	chain := blockchain.NewBlockchain(difficulty)
	chain.CreateTransaction(blockchain.Transaction{CharCredit: 11, Addr: blockchain.Addr{Key: "node1"}, Post: ""})
	chain.CreateTransaction(blockchain.Transaction{CharCredit: -5, Addr: blockchain.Addr{Key: "node1"}, Post: ""})
	chain.MinePendingTransactions("node2")
	fmt.Println("Balance of node1 is ", chain.GetBalanceOfAddress("node1"))
	fmt.Println("Balance of node1 is ", chain.GetBalanceOfAddress("node2"))
	chain.MinePendingTransactions("node1")
	fmt.Println("Balance of node1 is ", chain.GetBalanceOfAddress("node1"))
	fmt.Println("Balance of node1 is ", chain.GetBalanceOfAddress("node2"))
}
