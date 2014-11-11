package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Toorop/goabove"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetApiCredential Return credentials (ck)
func GetApiCredential(ak string) (credentials *raApiCredentials, err error) {
	client := &http.Client{}
	credentials = &raApiCredentials{}
	body := "{\"accessRules\":[{\"method\":\"GET\",\"path\":\"/*\"},{\"method\":\"POST\",\"path\":\"/*\"},{\"method\":\"DELETE\",\"path\":\"/*\"},{\"method\":\"PUT\",\"path\":\"/*\"},{\"method\":\"DELETE\",\"path\":\"/*\"} ]}"
	url := fmt.Sprintf("%s/%s/auth/credential", goabove.RA_API_ENDPOINT, goabove.RA_API_VERSION)
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
	req.Header.Add("User-Agent", "runabove-cli (https://github.com/Toorop/runabove-cli)")
	req.Header.Add("X-Ra-Application", ak)
	req.Header.Add("Content-type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return
	}
	// Bad HTTP status
	if resp.StatusCode > 399 {
		err = errors.New(resp.Status)
		return
	}
	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(r, credentials)
	return
}
