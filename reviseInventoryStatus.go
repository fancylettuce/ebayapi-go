package ebayapi

import "encoding/xml"

// ReviseInventoryStatusRequest struct
type ReviseInventoryStatusRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentials
	InventoryStatus      []*InventoryStatus `xml:"InventoryStatus"`
	ErrorLanguage        string             `xml:"ErrorLanguage,omitempty"`
	MessageID            string             `xml:"MessageID,omitempty"`
	Version              string             `xml:"Version,omitempty"`
	WarningLevel         string             `xml:"WarningLevel,omitempty"`
}

// InventoryStatus struct
type InventoryStatus struct {
	ItemID     string  `xml:"ItemID"`
	Quantity   int     `xml:"Quantity"`
	SKU        string  `xml:"SKU"`
	StartPrice float64 `xml:"StartPrice"`
}

// CallName returns name of call
func (rq ReviseInventoryStatusRequest) CallName() string {
	return "ReviseInventoryStatus"
}

// Body ataches credential and returns XML body
func (rq ReviseInventoryStatusRequest) Body(creds *Credentials) interface{} {
	rq.XMLName = xml.Name{
		Space: "urn:ebay:apis:eBLBaseComponents",
		Local: rq.CallName(),
	}
	rq.RequesterCredentials = &RequesterCredentials{EBayAuthToken: creds.AuthToken}
	return rq
}

// ParseResponse retruns response data as EbayResponse object
func (rq ReviseInventoryStatusRequest) ParseResponse(resp []byte) (EbayResponse, error) {
	var xmlResponse ReviseInventoryStatusResponse
	err := xml.Unmarshal(resp, &xmlResponse)

	return xmlResponse, err
}

// ResponseErrors returns errors
func (rs ReviseInventoryStatusResponse) ResponseErrors() EbayErrors {
	return rs.ebayResponse.Errors
}

// ReviseInventoryStatusResponse struct
type ReviseInventoryStatusResponse struct {
	ebayResponse
	InventoryStatus []InventoryStatus `xml:"InventoryStatus"`
}
