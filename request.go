package ebayapi

import (
	"encoding/xml"
)

// Call - interface for API calls
type Call interface {
	Body(*Credentials) interface{}
	CallName() string
	ParseResponse([]byte) (EbayResponse, error)
}

// RequesterCredentials holds auth token for API calls
type RequesterCredentials struct {
	EBayAuthToken string `xml:"eBayAuthToken"`
}

type ebayRequest struct {
	creds   Credentials
	command Call
}

func (c ebayRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	err := e.Encode(c.command.Body(&c.creds))
	if err != nil {
		return err
	}

	return nil
}
