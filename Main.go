package main

import (
	"./blockchain"
	"fmt"
)

func main() {
	blockchain.NewBlock("1st Block")
	blockchain.NewBlock("2nd Block")
	blockchain.DisplayAllBlocks()

	blockchain.ModifyBlock(1, "Modified Second Block")
	fmt.Println("After modification:")
	blockchain.DisplayAllBlocks()
}