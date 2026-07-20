package main

import (
	"fmt"
	"log"

	// Import our custom core package
	"github.com/tharanesh/go-chain/internal/core"
)

func main() {
	fmt.Println("=====================================")
	fmt.Println("   🚀 GO-CHAIN NODE INITIALIZED 🚀   ")
	fmt.Println("=====================================")

	log.Println("Booting up cryptographic ledger...")

	// 1. Initialize the Blockchain (This creates the Genesis Block automatically)
	blockchain := core.NewBlockchain()

	// 2. Add some test transactions to the chain
	log.Println("Mining Block 1...")
	blockchain.AddBlock("Alice sends 5 BTC to Bob")

	log.Println("Mining Block 2...")
	blockchain.AddBlock("Bob sends 2 BTC to Charlie")

	// 3. Print out the entire ledger to verify the hashes
	fmt.Println("\n--- 📖 CURRENT LEDGER 📖 ---")
	for i, block := range blockchain.Blocks {
		fmt.Printf("Block ID : %d\n", i)
		fmt.Printf("Data     : %s\n", string(block.Data))
		fmt.Printf("Prev Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash     : %x\n", block.Hash)
		fmt.Println("-------------------------------------")
	}

	log.Println("Node execution complete.")
}
