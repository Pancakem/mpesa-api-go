package mpesa

import (
	"time"
)

type authResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}

// C2B is a model
type C2B struct {
	ShortCode     string
	CommandID     string
	Amount        int
	Msisdn        string
	BillRefNumber string
}

// B2C is a model
type B2C struct {
	InitiatorName      string
	SecurityCredential string
	CommandID          string
	Amount             int
	PartyA             string
	PartyB             string
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
	Occassion          string
}

// B2B is a model
type B2B struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	SenderIdentifierType   string
	RecieverIdentifierType string
	Amount                 int
	PartyA                 string
	PartyB                 string
	Remarks                string
	AccountReference       string
	QueueTimeOutURL        string
	ResultURL              string
}

// Express is a model
type Express struct {
	BusinessShortCode string
	Password          string
	Timestamp         string
	TransactionType   string
	Amount            int
	PartyA            string
	PartyB            string
	PhoneNumber       string
	CallBackURL       string
	AccountReference  string
	TransactionDesc   string
}

// Mpesa Response Callback model
type MpesaCallbackResponse struct {
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  struct {
				Item []struct {
					Name  string `json:"Name"`
					Value string `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata"`
		} `json:"stkCallback"`
	} `json:"Body"`
}

type MpesaResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}

var defaultTransactionType = "CustomerPayBillOnline"

// NewExpress creates an express request object. Does the password generation and timestamp
func NewExpress(shortCode string, amount int, phoneNumber, callbackURL, ref, desc, passkey string) *Express {
	timestamp := time.Now().Format("20060102030405")
	password := GeneratePassword(shortCode, passkey, timestamp)

	return &Express{
		BusinessShortCode: shortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   defaultTransactionType,
		Amount:            amount,
		PartyA:            phoneNumber,
		PartyB:            shortCode,
		PhoneNumber:       phoneNumber,
		CallBackURL:       callbackURL,
		AccountReference:  ref,
		TransactionDesc:   desc,
	}
}

// Reversal is a model
type Reversal struct {
	Initiator              string
	SecurityCredential     string
	CommandID              string
	TransactionID          string
	Amount                 int
	ReceiverParty          string
	ReceiverIdentifierType string
	QueueTimeOutURL        string
	ResultURL              string
	Remarks                string
	Occassion              string
}

// BalanceInquiry is a model
type BalanceInquiry struct {
	Initiator          string
	SecurityCredential string
	CommandID          string
	PartyA             string
	IdentifierType     string
	Remarks            string
	QueueTimeOutURL    string
	ResultURL          string
}

// Pull is a model
type Pull struct {
	ShortCode  string
	StartDate  string
	EndDate    string
	PageNumber string
}

// RegisterURL is a model
type C2BRegisterURL struct {
	ShortCode       string
	ResponseType    string
	ConfirmationURL string
	ValidationURL   string
}
