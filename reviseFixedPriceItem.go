package ebayapi

import "encoding/xml"

// ReviseFixedPriceItemRequest struct
type ReviseFixedPriceItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentials
	Item                 *Item
	ErrorLanguage        string `xml:",omitempty"`
	MessageID            string `xml:",omitempty"`
	Version              string `xml:",omitempty"`
	WarningLevel         string `xml:",omitempty"`
}

// CallName returns name of call
func (rq ReviseFixedPriceItemRequest) CallName() string {
	return "ReviseFixedPriceItem"
}

// Body ataches credential and returns XML body
func (rq ReviseFixedPriceItemRequest) Body(creds *Credentials) interface{} {
	rq.XMLName = xml.Name{
		Space: "urn:ebay:apis:eBLBaseComponents",
		Local: rq.CallName(),
	}
	rq.RequesterCredentials = &RequesterCredentials{EBayAuthToken: creds.AuthToken}
	return rq
}

// ParseResponse retruns response data as EbayResponse object
func (rq ReviseFixedPriceItemRequest) ParseResponse(resp []byte) (EbayResponse, error) {
	var xmlResponse ReviseFixedPriceItemResponse
	err := xml.Unmarshal(resp, &xmlResponse)

	return xmlResponse, err
}

// ResponseErrors returns errors
func (rs ReviseFixedPriceItemResponse) ResponseErrors() EbayErrors {
	return rs.ebayResponse.Errors
}

// ReviseFixedPriceItemResponse struct
type ReviseFixedPriceItemResponse struct {
	ebayResponse
}
