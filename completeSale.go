package ebayapi

import (
	"encoding/xml"
	"time"
)

// CompleteSaleRequest type
type CompleteSaleRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentials
	ItemID               string    `xml:"ItemID,omitempty"`
	OrderID              string    `xml:"OrderID,omitempty"`
	OrderLineItemID      string    `xml:"OrderLineItemID,omitempty"`
	Paid                 bool      `xml:"Paid,omitempty"`
	Shipment             *Shipment `xml:"Shipment,omitempty"`
	Shipped              bool      `xml:"Shipped,omitempty"`
	TransactionID        string    `xml:"TransactionID,omitempty"`
	ErrorHandling        string    `xml:"ErrorHandling,omitempty"`
	ErrorLanguage        string    `xml:"ErrorLanguage,omitempty"`
	MessageID            string    `xml:"MessageID,omitempty"`
	Version              string    `xml:"Version,omitempty"`
	WarningLevel         string    `xml:"WarningLevel,omitempty"`
	// FeedbackInfo struct {
	// 	CommentText string `xml:"CommentText,omitempty"`
	// 	CommentType string `xml:"CommentType,omitempty"`
	// 	TargetUser  string `xml:"TargetUser,omitempty"`
	// } `xml:"FeedbackInfo,omitempty"`
}

// Shipment type
type Shipment struct {
	ShipmentTrackingDetails []*ShipmentTrackingDetails `xml:"ShipmentTrackingDetails,omitempty"`
	ShippedTime             *time.Time                 `xml:"ShippedTime,omitempty"`
}

// ShipmentTrackingDetails type
type ShipmentTrackingDetails struct {
	ShipmentLineItem       *ShipmentLineItem `xml:"ShipmentLineItem,omitempty"`
	ShipmentTrackingNumber string            `xml:"ShipmentTrackingNumber,omitempty"`
	ShippingCarrierUsed    string            `xml:"ShippingCarrierUsed,omitempty"`
}

// ShipmentLineItem type
type ShipmentLineItem struct {
	LineItem []*LineItem `xml:"LineItem,omitempty"`
}

// LineItem type
type LineItem struct {
	CountryOfOrigin string `xml:"CountryOfOrigin,omitempty"`
	Description     string `xml:"Description,omitempty"`
	ItemID          string `xml:"ItemID,omitempty"`
	Quantity        int32  `xml:"Quantity,omitempty"`
	TransactionID   string `xml:"TransactionID,omitempty"`
}

// CallName returns name of call
func (c CompleteSaleRequest) CallName() string {
	return "CompleteSale"
}

// Body ataches credential and returns XML body
func (c CompleteSaleRequest) Body(creds *Credentials) interface{} {
	c.XMLName = xml.Name{
		Space: "urn:ebay:apis:eBLBaseComponents",
		Local: c.CallName(),
	}
	c.RequesterCredentials = &RequesterCredentials{EBayAuthToken: creds.AuthToken}
	return c
}

// ParseResponse retruns response data as EbayResponse object
func (c CompleteSaleRequest) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse CompleteSaleResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

// ResponseErrors returns errors
func (r CompleteSaleResponse) ResponseErrors() EbayErrors {
	return r.ebayResponse.Errors
}

// CompleteSaleResponse type
type CompleteSaleResponse struct {
	ebayResponse
}
