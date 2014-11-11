package region

import (
	"encoding/json"
	"github.com/Toorop/goabove"
)

// regionRessource represents a "me" ressources
type regionRessource struct {
	*goabove.ApiClient
}

// New return a new regionRessource
func New(client *goabove.ApiClient) (*regionRessource, error) {
	if client == nil {
		return nil, goabove.ErrNoRaApiClient
	}
	return &regionRessource{client}, nil
}

// GetRegions  return available regions
func (r *regionRessource) GetAll() (regions []string, err error) {
	resp, err := r.Call("GET", "region", "")
	if err = resp.HandleErr(err, []int{200}); err != nil {
		return
	}
	err = json.Unmarshal(resp.Body, &regions)
	return
}
