package overlay

import "net/http"

// Create creates a new wallet and returns a flattened response map.
// This endpoint does not require authentication.
func (r *WalletResource) Create(params map[string]interface{}) (map[string]interface{}, error) {
	return r.client.sendRequest(http.MethodPost, "/wallets", params, false)
}
