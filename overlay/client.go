package overlay

import (
	"net/http"
)

const (
	EnvDevnet  = "devnet"
	EnvMainnet = "mainnet"
)

// Response is the standardized response from the API.
type Response struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

// Config holds the configuration for the API client.
type Config struct {
	APIKey  string
	AuthKey string
	Env     string // Use EnvDevnet or EnvMainnet
}

// Client is the main client for the Overlay API.
type Client struct {
	config     Config
	baseURL    string
	httpClient *http.Client

	Wallet *WalletResource
	Token  *TokenResource
	NFT    *NFTResource
}

// NewClient creates a new Overlay API client from a config.
func NewClient(config Config) *Client {
	env := config.Env
	if env == "" {
		env = EnvDevnet // Default to devnet
	}

	baseURL := ""
	switch env {
	case EnvMainnet:
		baseURL = "https://mainnet.overlay.fun"
	default:
		baseURL = "https://devnet.overlay.fun"
	}

	c := &Client{
		config:     config,
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}

	c.Wallet = &WalletResource{client: c}
	c.Token = &TokenResource{client: c}
	c.NFT = &NFTResource{client: c}

	return c
}

// WalletResource handles wallet-related operations.
type WalletResource struct {
	client *Client
}

// TokenResource handles token-related operations.
type TokenResource struct {
	client *Client
}

// NFTResource handles NFT-related operations.
type NFTResource struct {
	client *Client
}
