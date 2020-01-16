package ebayapi

import "encoding/xml"

// GetItemRequest obj
type GetItemRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentials
	ItemID               string
}

// CallName returns name of call
func (c GetItemRequest) CallName() string {
	return "GetItem"
}

// Body ataches credential and returns XML body
func (c GetItemRequest) Body(creds *Credentials) interface{} {
	c.XMLName = xml.Name{
		Space: "urn:ebay:apis:eBLBaseComponents",
		Local: c.CallName(),
	}
	c.RequesterCredentials = &RequesterCredentials{EBayAuthToken: creds.AuthToken}
	return c
}

// ParseResponse retruns response data as EbayResponse object
func (c GetItemRequest) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse GetItemResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

// GetItemResponse response object
type GetItemResponse struct {
	ebayResponse

	Item struct {
		ItemID        string
		Quantity      int64
		SellingStatus struct {
			ListingStatus string
			QuantitySold  int64
		}
	}
}

// ResponseErrors returns errors
func (r GetItemResponse) ResponseErrors() EbayErrors {
	return r.ebayResponse.Errors
}
