# Go-Chain
A custom blockchain engine written in Go. Demonstrates applied cryptography, immutable data structures (Linked Lists), and Proof-of-Work mining.

# How to Run Your Blockchain in Ubuntu

**Step 1: Go to your project folder**
Type this command

```
cd ~/Go-Chain
```

**Step 2: Run the Blockchain Engine**
Execute your Go code by running this exact command:

```
go run cmd/node/main.go
```

# Project Structure

```
~/Go-Chain/
├── cmd/
│   └── node/
│       └── main.go           # The ignition key (App entry point)
├── internal/
│   └── core/
│       ├── block.go          # The SHA-256 logic and block data structure
│       └── blockchain.go     # The ledger management and linking logic
├── go.mod                    # The manifest/passport for your Go module
```
