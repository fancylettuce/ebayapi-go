package ebayapi

import (
	"encoding/json"
	"net/url"
	"strconv"
)

//
// The 'Login' call from the ebay ClientAlerts API uses a REST Querystring request and recieves a JSON response
//

// LoginRequest holds credential data for Client Alerts API login
type LoginRequest struct {
	ClientAlertsAuthToken string
}

// CallName retruns the string name of this call
func (r LoginRequest) CallName() string {
	return "Login"
}

// Body ataches credential and returns url querystring values
func (r LoginRequest) Body(creds *Credentials) interface{} {
	return url.Values{
		"version":               {"957"},
		"appid":                 {creds.AppID},
		"callname":              {r.CallName()},
		"ClientAlertsAuthToken": {r.ClientAlertsAuthToken},
	}
}

// ParseResponse unmarshals response bytes into obj
func (r LoginRequest) ParseResponse(resp []byte) (EbayResponse, error) {
	var response LoginResponse
	err := json.Unmarshal(resp, &response)
	return response, err
}

// LoginResponse is Unmarshalled from JSON
type LoginResponse struct {
	Ack           string `json:"Ack"`
	SessionData   string `json:"SessionData"`
	SessionID     string `json:"SessionID"`
	Build         string `json:"Build"`
	CorrelationID string `json:"CorrelationID"`
	Errors        []struct {
		ErrorClassification string `json:"ErrorClassification"`
		ErrorCode           string `json:"ErrorCode"`
		ErrorParameters     struct {
			Value string `json:"Value"`
		} `json:"ErrorParameters"`
		LongMessage  string `json:"LongMessage"`
		SeverityCode string `json:"SeverityCode"`
		ShortMessage string `json:"ShortMessage"`
	} `json:"Errors"`
	Timestamp string `json:"Timestamp"`
	Version   string `json:"Version"`
}

// Failure checks if call failed
func (r LoginResponse) Failure() bool {
	return r.Ack == "Failure"
}

// ResponseErrors returns error array
func (r LoginResponse) ResponseErrors() EbayErrors {
	var errors EbayErrors
	for _, err := range r.Errors {
		code, _ := strconv.Atoi(err.ErrorCode)
		errors = append(errors, ebayResponseError{
			ShortMessage:        err.ShortMessage,
			LongMessage:         err.LongMessage,
			ErrorCode:           code,
			SeverityCode:        err.SeverityCode,
			ErrorClassification: err.ErrorClassification,
		})
	}

	return errors
}
