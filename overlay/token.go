package overlay

import (
	"fmt"
	"net/http"
)

// Create creates a new token and returns a flattened response map.
func (r *TokenResource) Create(params map[string]interface{}) (map[string]interface{}, error) {
	return r.client.sendRequest(http.MethodPost, "/tokens", params, true)
}

// Mint mints new tokens and returns a flattened response map.
func (r *TokenResource) Mint(mintAddress string, params map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/tokens/%s/mint", mintAddress)
	return r.client.sendRequest(http.MethodPost, path, params, true)
}

// Transfer transfers tokens and returns a flattened response map.
func (r *TokenResource) Transfer(mintAddress string, params map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/tokens/%s/transfer", mintAddress)
	return r.client.sendRequest(http.MethodPost, path, params, true)
}

// Burn burns tokens and returns a flattened response map.
func (r *TokenResource) Burn(mintAddress string, params map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/tokens/%s/burn", mintAddress)
	return r.client.sendRequest(http.MethodPost, path, params, true)
}
