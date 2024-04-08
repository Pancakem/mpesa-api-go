package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

// Env is the environment type
type Env string

const (
	// DEV is the development env tag

	// SANDBOX is the sandbox env tag
	SANDBOX = iota
	// PRODUCTION is the production env tag
	PRODUCTION
)

// Service is an Mpesa Service
type Service struct {
	AppKey    string
	AppSecret string
	Env       int
}

// New return a new Mpesa Service
func New(appKey, appSecret string, env int) (Service, error) {
	return Service{appKey, appSecret, env}, nil
}

// Generate Mpesa Daraja Access Token
func (s Service) auth() (string, error) {
	url := s.baseURL() + "oauth/v1/generate?grant_type=client_credentials"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(s.AppKey, s.AppSecret)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Connection", "keep-alive")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "failed to send request")
	}

	if res.StatusCode != 200 {
		return "", errors.New(fmt.Sprintf("request failed with status %d", res.StatusCode))
	}

	defer res.Body.Close()
	var authResp authResponse
	err = json.NewDecoder(res.Body).Decode(&authResp)
	if err != nil {
		return "", errors.Wrap(err, "could not decode auth response")
	}

	accessToken := authResp.AccessToken
	return accessToken, nil
}

// Simulation requests user device for payment
func (s Service) MpesaExpress(express Express) (*MpesaResponse, error) {
	body, err := json.Marshal(express)
	if err != nil {
		return nil, err
	}
	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["content-type"] = "application/json"
	headers["authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := s.baseURL() + "mpesa/stkpush/v1/processrequest"
	return s.newReq(url, body, headers)
}

// TransactionStatus gets status of a transaction
func (s Service) TransactionStatus(express Express) (*MpesaResponse, error) {
	body, err := json.Marshal(express)
	if err != nil {
		return nil, err
	}

	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth

	url := s.baseURL() + "mpesa/stkpushquery/v1/query"
	return s.newReq(url, body, headers)
}

// C2BRegisterURL requests
func (s Service) C2BRegisterURL(c2bRegisterURL C2BRegisterURL) (*MpesaResponse, error) {
	body, err := json.Marshal(c2bRegisterURL)
	if err != nil {
		return nil, err
	}

	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["Cache-Control"] = "no-cache"

	url := s.baseURL() + "mpesa/c2b/v1/registerurl"
	return s.newReq(url, body, headers)
}

// C2BSimulation sends a new request
func (s Service) C2BSimulation(c2b C2B) (*MpesaResponse, error) {
	body, err := json.Marshal(c2b)
	if err != nil {
		return nil, err
	}

	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := s.baseURL() + "mpesa/c2b/v1/simulate"
	return s.newReq(url, body, headers)
}

// B2CRequest sends a new request
func (s Service) B2CRequest(b2c B2C) (*MpesaResponse, error) {
	body, err := json.Marshal(b2c)
	if err != nil {
		return nil, err
	}

	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := s.baseURL() + "mpesa/b2c/v1/paymentrequest"
	return s.newReq(url, body, headers)
}

// B2BRequest sends a new request
func (s Service) B2BRequest(b2b B2B) (*MpesaResponse, error) {
	body, err := json.Marshal(b2b)
	if err != nil {
		return nil, err
	}
	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := s.baseURL() + "mpesa/b2b/v1/paymentrequest"
	return s.newReq(url, body, headers)
}

// Reversal requests a reversal?
func (s Service) Reversal(reversal Reversal) (*MpesaResponse, error) {
	body, err := json.Marshal(reversal)
	if err != nil {
		return nil, err
	}

	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := s.baseURL() + "safaricom/reversal/v1/request" //TODO :: CONFIRM THIS URL/ENDPOINT???
	return s.newReq(url, body, headers)
}

// BalanceInquiry sends a balance inquiry
func (s Service) BalanceInquiry(balanceInquiry *BalanceInquiry) (*MpesaResponse, error) {
	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(balanceInquiry)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"
	headers["postman-token"] = "2aa448be-7d56-a796-065f-b378ede8b136"

	url := s.baseURL() + "mpesa/accountbalance/v1/query"
	return s.newReq(url, body, headers)
}

// BalanceInquiry sends a balance inquiry
func (s Service) PullTransactions(pull Pull) (*MpesaResponse, error) {
	auth, err := s.auth()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(pull)
	if err != nil {
		return nil, err
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = "Bearer " + auth
	headers["cache-control"] = "no-cache"

	url := s.baseURL() + "pulltransactions/v1/query"
	return s.newReq(url, body, headers)
}

func (s Service) newReq(url string, body []byte, headers map[string]string) (*MpesaResponse, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	client := &http.Client{Timeout: 60 * time.Second}
	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	var mpesaResp MpesaResponse

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&mpesaResp)
	if err != nil {
		return nil, err
	}

	return &mpesaResp, nil
}

func (s Service) baseURL() string {
	if s.Env == PRODUCTION {
		return "https://api.safaricom.co.ke/"
	}
	return "https://sandbox.safaricom.co.ke/"
}
