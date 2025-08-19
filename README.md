# overlay-go

A Go library for **Web3 development** that makes it easy to implement wallet creation, token and NFT issuance, minting, burning, and transfers across multiple blockchains.

---

## 📖 Overview

**Overlay** enables simple Web3 development across multiple blockchains, including:
- **Solana** 
- **Ethereum**
- **BSC (Binance Smart Chain)**
- **Base**
- **Polygon** 
- **Bera Chain**

It eliminates the need for signing transactions with wallets, which is typically necessary in blockchain development. Overlay implements high-level security features and securely encrypts wallet information to safely execute transactions without storing any database locally.

---

## 🚀 Installation

### Via Go Modules

```bash
go mod init your-project
go get github.com/overlaydotfun/overlay-go
```

---

## ⚙️ Configuration

Before using the library, you need to configure your API credentials:

```go
package main

import (
    "fmt"
    "github.com/overlaydotfun/overlay-go/overlay"
)

func main() {
    client := overlay.NewClient(overlay.Config{
        APIKey:  "your_api_key",
        AuthKey: "your_auth_key",
        Env:     overlay.EnvDevnet, // or overlay.EnvMainnet
    })
}
```

### 🌐 Environment Options

| Environment | Constant | URL | Description |
|-------------|----------|-----|-------------|
| `devnet` | `overlay.EnvDevnet` | `https://devnet.overlay.fun` | Development environment |
| `mainnet` | `overlay.EnvMainnet` | `https://mainnet.overlay.fun` | Production environment |

---

## 📘 Usage

### 👛 Wallet

#### Create a Wallet

```go
// Create a new wallet
response, err := client.Wallet.Create(map[string]interface{}{
    "network": "solana", // or "evm"
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("Wallet created: %v\n", response)
```

### 🪙 Token

#### Create a Token

```go
// Create a new token
response, err := client.Token.Create(map[string]interface{}{
    "network":  "solana", // or "eth", "base", "bsc", "blast", "polygon", "bera", "monad", etc
    "name":     "My Token",
    "symbol":   "MTK",
    "image":    "image url",
    "supply":   1000000000,
    "decimals": 9,
    // Add other token parameters
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("Token created: %v\n", response)
```

#### Mint Tokens

```go
// Mint tokens to an address
mintAddress := "your_token_mint_address"
response, err := client.Token.Mint(mintAddress, map[string]interface{}{
    "network": "solana",
    "amount":  1000,
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("Success: %v\n", response["success"])
```

#### Transfer Tokens

```go
// Transfer tokens between addresses
response, err := client.Token.Transfer(mintAddress, map[string]interface{}{
    "network":   "solana",
    "amount":    1000,
    "recipient": "recipient_address",
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("Transfer result: %v\n", response)
```

#### Burn Tokens

```go
// Burn tokens
response, err := client.Token.Burn(mintAddress, map[string]interface{}{
    "network": "solana",
    "amount":  1000,
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("Burn result: %v\n", response)
```

### 🖼️ NFT

#### Create an NFT

```go
// Create a new NFT collection or individual NFT
response, err := client.NFT.Create(map[string]interface{}{
    "network":     "solana",
    "name":        "My NFT",
    "symbol":      "MNFT",
    "description": "A unique NFT",
    "images":      []string{"https://example.com/image.png"},
    // Add other NFT metadata
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("NFT created: %v\n", response["success"])
```

#### Mint NFTs

```go
// Mint an NFT
mintAddress := "your_nft_mint_address"
response, err := client.NFT.Mint(mintAddress, map[string]interface{}{
    "network": "solana",
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("NFT minted: %v\n", response)
```

#### Transfer NFTs

```go
// Transfer an NFT
response, err := client.NFT.Transfer(mintAddress, map[string]interface{}{
    "network":   "solana",
    "recipient": "new_owner",
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("NFT transferred: %v\n", response)
```

#### Burn NFTs

```go
// Burn an NFT
response, err := client.NFT.Burn(mintAddress, map[string]interface{}{
    "network": "solana",
})

if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}

fmt.Printf("NFT burned: %v\n", response)
```

---

## 📦 Response Format

All methods return a flattened `map[string]interface{}` with the following structure:

```go
response, err := client.Token.Create(params)

if err != nil {
    // Handle error
    return
}

// Access response properties
success := response["success"].(bool)
message := response["message"].(string)

// Data properties are flattened to top level
if mintAddress, ok := response["mint_address"]; ok {
    fmt.Printf("Mint Address: %v\n", mintAddress)
}

if transactionID, ok := response["transaction_id"]; ok {
    fmt.Printf("Transaction ID: %v\n", transactionID)
}
```

---

## ⚠️ Error Handling

The library includes built-in error handling. Always check for errors and the `success` property:

```go
response, err := client.Token.Mint(mintAddress, map[string]interface{}{
    "amount": 1000,
})

if err != nil {
    fmt.Printf("❌ Request failed: %v\n", err)
    return
}

if success, ok := response["success"].(bool); ok && success {
    fmt.Println("✅ Token minted successfully!")
    if txID, exists := response["transaction_id"]; exists {
        fmt.Printf("Transaction ID: %v\n", txID)
    }
} else {
    if message, exists := response["message"]; exists {
        fmt.Printf("❌ Error: %v\n", message)
    }
}
```

---

## 🌍 Supported Blockchains

- 🔥 **Solana**
- ⚡ **Ethereum**
- 💛 **Binance Smart Chain (BSC)**
- 🔵 **Base**
- 🟣 **Polygon**
- 🐻 **Bera Chain**

---

## 🔐 Security Features

- ✅ **No local database storage**
- 🔒 **Secure wallet encryption**
- 🛡️ **High-level security implementation**
- 🔑 **Safe transaction execution without exposing private keys**

---

## 📋 Requirements

- **Go** >= 1.21
- Internet connection for API requests

---

## 🛠️ Development

After cloning the repository:

```bash
git clone https://github.com/overlaydotfun/overlay-go.git
cd overlay-go
go mod tidy
```

### Running Tests

```bash
go test ./...
```

---

## 🤝 Contributing

Bug reports and pull requests are welcome on GitHub at [https://github.com/overlaydotfun/overlay-go](https://github.com/overlaydotfun/overlay-go).

---

## 📚 Documentation

For more detailed documentation, visit: [https://docs.overlay.fun](https://docs.overlay.fun)

---

## 📄 License

This library is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).

---

## 📞 Contact

- 📧 **Email**: contact@overlay.fun
- 🌐 **Website**: https://overlay.fun
- 📖 **Documentation**: https://docs.overlay.fun

---

## Examples

### Complete Example

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/overlaydotfun/overlay-go/overlay"
)

func main() {
    // Initialize client
    client := overlay.NewClient(overlay.Config{
        APIKey:  "your_api_key",
        AuthKey: "your_auth_key",
        Env:     overlay.EnvDevnet,
    })

    // Create a wallet
    walletResponse, err := client.Wallet.Create(map[string]interface{}{
        "network": "solana",
    })
    
    if err != nil {
        log.Fatalf("Failed to create wallet: %v", err)
    }
    
    if success, ok := walletResponse["success"].(bool); ok && success {
        if walletAddr, exists := walletResponse["wallet_address"]; exists {
            fmt.Printf("Wallet created: %v\n", walletAddr)
        }
    }
    
    // Create a token
    tokenResponse, err := client.Token.Create(map[string]interface{}{
        "network":  "solana",
        "name":     "Test Token",
        "symbol":   "TEST",
        "supply":   1000000,
        "decimals": 9,
    })
    
    if err != nil {
        log.Fatalf("Failed to create token: %v", err)
    }
    
    if success, ok := tokenResponse["success"].(bool); ok && success {
        if mintAddr, exists := tokenResponse["mint_address"]; exists {
            fmt.Printf("Token created: %v\n", mintAddr)
            
            // Mint some tokens
            mintResponse, err := client.Token.Mint(mintAddr.(string), map[string]interface{}{
                "network": "solana",
                "amount":  1000,
            })
            
            if err != nil {
                log.Printf("Failed to mint tokens: %v", err)
                return
            }
            
            if mintSuccess, ok := mintResponse["success"].(bool); ok && mintSuccess {
                fmt.Println("Tokens minted successfully!")
            }
        }
    }
}
```

### Error Handling Example

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/overlaydotfun/overlay-go/overlay"
)

func handleTokenOperation() {
    client := overlay.NewClient(overlay.Config{
        APIKey:  "your_api_key",
        AuthKey: "your_auth_key", 
        Env:     overlay.EnvDevnet,
    })

    response, err := client.Token.Create(map[string]interface{}{
        "network": "solana",
        "name":    "My Token",
        "symbol":  "MTK",
        "supply":  1000000,
    })

    if err != nil {
        // Handle network or API errors
        log.Printf("API request failed: %v", err)
        return
    }

    // Check if the operation was successful
    success, exists := response["success"]
    if !exists {
        log.Println("Invalid response format")
        return
    }

    if successBool, ok := success.(bool); ok && successBool {
        fmt.Println("✅ Token created successfully!")
        
        // Access response data
        if mintAddress, exists := response["mint_address"]; exists {
            fmt.Printf("Mint Address: %v\n", mintAddress)
        }
        
        if txID, exists := response["transaction_id"]; exists {
            fmt.Printf("Transaction ID: %v\n", txID)
        }
    } else {
        // Handle business logic errors
        if message, exists := response["message"]; exists {
            fmt.Printf("❌ Operation failed: %v\n", message)
        } else {
            fmt.Println("❌ Unknown error occurred")
        }
    }
}

func main() {
    handleTokenOperation()
}
```