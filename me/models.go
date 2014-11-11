package me

// User represent the result of getuserInfo
type userInfo struct {
	Firstname         string `json:"firstname"`
	Country           string `json:"country"`
	Commander         string `json:"commander"`
	Area              string `json:"area"`
	Twitter           string `json:"twitter"`
	Status            string `json:"status"`
	CellNumber        string `json:"cellNumber"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	AccountIdentifier string `json:"accountIdentifier"`
	City              string `json:"city"`
	Address           string `json:"address"`
	PostalCode        string `json:"postalCode"`
}
