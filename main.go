package main

import (
	"fmt"
	"log"
)

// main is the entry point for our blockchain node.
// Currently, it just boots up, but soon it will initialize
// the ledger and start the network server.
func main() {
	fmt.Println("=====================================")
	fmt.Println("   🚀 GO-CHAIN NODE INITIALIZED 🚀   ")
	fmt.Println("=====================================")
	
	log.Println("Booting up cryptographic ledger...")
	
	// TODO in Commit 3: Initialize the Blockchain
	// TODO in Commit 5: Start the P2P Network Server
	
	log.Println("Node is running (Press CTRL+C to exit)")
}
