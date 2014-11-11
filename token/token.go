package token

import (
	"encoding/json"
	"github.com/Toorop/goabove"
	"github.com/Toorop/gopenstack"
)

// tokenRessource  Getepresents a "auth" ressources
type tokenRessource struct {
	*goabove.ApiClient
}

// New return a new tokenRessource
func New(client *goabove.ApiClient) (*tokenRessource, error) {
	if client == nil {
		return nil, goabove.ErrNoRaApiClient
	}
	return &tokenRessource{client}, nil
}

// Get return a valid gopenstack keyring to use with gopenstack
func (r *tokenRessource) GetGosKeyring() (*gopenstack.Keyring, error) {
	resp, err := r.Call("GET", "token", "")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return nil, err
	}
	var keyring gopenstack.Keyring
	err = json.Unmarshal(resp.Body, &keyring)
	if err != nil {
		return nil, err
	}
	return &keyring, nil
}
