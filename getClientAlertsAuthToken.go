package ebayapi

import "encoding/xml"

// GetClientAlertsAuthTokenRequest ...
type GetClientAlertsAuthTokenRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentials
}

// CallName returns name of call
func (r GetClientAlertsAuthTokenRequest) CallName() string {
	return "GetClientAlertsAuthToken"
}

// Body ataches credential and returns XML body
func (r GetClientAlertsAuthTokenRequest) Body(creds *Credentials) interface{} {
	r.XMLName = xml.Name{
		Space: "urn:ebay:apis:eBLBaseComponents",
		Local: r.CallName(),
	}
	r.RequesterCredentials = &RequesterCredentials{EBayAuthToken: creds.AuthToken}
	return r
}

// ParseResponse retruns response data as EbayResponse object
func (r GetClientAlertsAuthTokenRequest) ParseResponse(resp []byte) (EbayResponse, error) {
	var xmlResponse GetClientAlertsAuthTokenResponse
	err := xml.Unmarshal(resp, &xmlResponse)
	return xmlResponse, err
}

// GetClientAlertsAuthTokenResponse response object
type GetClientAlertsAuthTokenResponse struct {
	ebayResponse
	XMLName               xml.Name `xml:"GetClientAlertsAuthTokenResponse"`
	Xmlns                 string   `xml:"xmlns,attr"`
	ClientAlertsAuthToken string   `xml:"ClientAlertsAuthToken"`
	HardExpirationTime    string   `xml:"HardExpirationTime"`
	Build                 string   `xml:"Build"`
	CorrelationID         string   `xml:"CorrelationID"`
	HardExpirationWarning string   `xml:"HardExpirationWarning"`
	Version               string   `xml:"Version"`
}

// ResponseErrors returns errors
func (r GetClientAlertsAuthTokenResponse) ResponseErrors() EbayErrors {
	return r.ebayResponse.Errors
}
