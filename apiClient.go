package goabove

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// raResponseErr represents an unmarshalled response from Runabove API in case od error
type raResponseErr struct {
	ErrorCode string `json:"errorCode"`
	HttpCode  string `json:"httpCode"`
	Message   string `json:"message"`
}

// apiresponse represent a ruabove API response
type apiResponse struct {
	StatusCode int
	Status     string
	Body       []byte
}

// handleCommon return error on unexpected HTTP code
func (r *apiResponse) HandleErr(err error, expectedHttpCode []int) error {
	if err != nil {
		return err
	}
	for _, code := range expectedHttpCode {
		if r.StatusCode == code {
			return nil
		}
	}
	// Try to get API response
	if r.Body != nil {
		var raResponseErr raResponseErr
		err := json.Unmarshal(r.Body, &raResponseErr)
		if err == nil {
			if len(raResponseErr.ErrorCode) != 0 {
				if raResponseErr.ErrorCode == "INVALID_SIGNATURE" {
					raResponseErr.ErrorCode = "INVALID_SIGNATURE_OR_CONSUMERKEY_HAS_EXPIRATED"
				}
				return errors.New(raResponseErr.ErrorCode)
			} else {
				return errors.New(raResponseErr.Message)
			}
		}
	}
	return errors.New(fmt.Sprintf("%d - %s", r.StatusCode, r.Status))
}

// ApiClient represent an runabove HTTP API client
type ApiClient struct {
	*http.Client
	ak string // Application key
	as string // Application secret
	ck string // Consumer key
}

// newRaClient returns a new ApiClient
func NewClient(ak, as, ck string) (c *ApiClient) {
	return &ApiClient{&http.Client{}, ak, as, ck}
}

func (c *ApiClient) Call(method string, ressource string, payload string) (response apiResponse, err error) {
	query := fmt.Sprintf("%s/%s/%s", RA_API_ENDPOINT, RA_API_VERSION, ressource)
	req, err := http.NewRequest(method, query, strings.NewReader(payload))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")

	timestamp := fmt.Sprintf("%d", int32(time.Now().Unix()))
	req.Header.Add("X-Ra-Timestamp", timestamp)
	req.Header.Add("X-Ra-Application", c.ak)
	req.Header.Add("X-Ra-Consumer", c.ck)
	req.Header.Add("User-Agent", "goabove (https://github.com/Toorop/goabove)")
	p := strings.Split(ressource, "?")
	req.URL.Opaque = fmt.Sprintf("/%s/%s", RA_API_VERSION, p[0])

	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("%s+%s+%s+%s+%s+%s", c.as, c.ck, method, query, payload, timestamp)))
	req.Header.Add("X-Ra-Signature", fmt.Sprintf("$1$%x", h.Sum(nil)))

	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	response.StatusCode = resp.StatusCode
	response.Status = resp.Status
	response.Body, err = ioutil.ReadAll(resp.Body)
	//fmt.Println(resp.StatusCode, string(response.Body))

	return
}
