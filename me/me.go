package me

import (
	"encoding/json"
	"github.com/Toorop/goabove"
)

// meRessource represents a "me" ressources
type meRessource struct {
	*apiClient
}

// New return a new meRessource
func New(client *apiClient) (*meRessource, error) {
	if client == nil {
		return nil, goabove.ErrNoRaApiClient
	}
	return &meRessource{client}, nil
}

//  GetUserInfo returns user info
func (r *meRessource) GetUserInfo() (info userInfo, err error) {
	resp, err := r.Call("GET", "me", "")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &info)
	return
}
