package overlay

import (
	"fmt"
	"net/http"
)

// Create creates a new NFT and returns a flattened response map.
func (r *NFTResource) Create(params map[string]interface{}) (map[string]interface{}, error) {
	return r.client.sendRequest(http.MethodPost, "/nfts", params, true)
}

// Mint mints a new NFT and returns a flattened response map.
func (r *NFTResource) Mint(mintAddress string, params map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/nfts/%s/mint", mintAddress)
	return r.client.sendRequest(http.MethodPost, path, params, true)
}

// Transfer transfers an NFT and returns a flattened response map.
func (r *NFTResource) Transfer(mintAddress string, params map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/nfts/%s/transfer", mintAddress)
	return r.client.sendRequest(http.MethodPost, path, params, true)
}

// Burn burns an NFT and returns a flattened response map.
func (r *NFTResource) Burn(mintAddress string, params map[string]interface{}) (map[string]interface{}, error) {
	path := fmt.Sprintf("/nfts/%s/burn", mintAddress)
	return r.client.sendRequest(http.MethodPost, path, params, true)
}
