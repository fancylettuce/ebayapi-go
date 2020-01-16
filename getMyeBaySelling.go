package ebayapi

import (
	"encoding/xml"
)

// GetMyeBaySellingRequest type
type GetMyeBaySellingRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentials
	ActiveList           *ActiveListRequest `xml:",omitempty"`
	// SellingSummary struct {
	// 	Include bool `xml:"Include"`
	// } `xml:"SellingSummary"`
	DetailLevel    []string `xml:",omitempty"`
	ErrorLanguage  string   `xml:",omitempty"`
	MessageID      string   `xml:",omitempty"`
	OutputSelector []string `xml:",omitempty"`
	Version        string   `xml:",omitempty"`
	WarningLevel   string   `xml:",omitempty"`
}

//ActiveListRequest struct
type ActiveListRequest struct {
	Include      bool        `xml:",omitempty"`
	IncludeNotes string      `xml:",omitempty"`
	ListingType  string      `xml:",omitempty"`
	Pagination   *Pagination `xml:",omitempty"`
	Sort         string      `xml:",omitempty"`
}

// DefaultOutputSelection selects only order data fields required for order entry
func (r GetMyeBaySellingRequest) DefaultOutputSelection() *GetMyeBaySellingRequest {
	r.OutputSelector = []string{
		"ActiveList.ItemArray.Item.SKU",
		"ActiveList.ItemArray.Item.ItemID",
		"ActiveList.ItemArray.Item.Title",
		"ActiveList.ItemArray.Item.QuantityAvailable",
		"ActiveList.ItemArray.Item.SellingStatus.CurrentPrice",
		"ActiveList.PaginationResult.TotalNumberOfEntries",
		"ActiveList.PaginationResult.TotalNumberOfPages",
	}
	return &r
}

// CallName returns name of call
func (r GetMyeBaySellingRequest) CallName() string {
	return "GetMyeBaySelling"
}

// Body ataches credential and returns XML body
func (r GetMyeBaySellingRequest) Body(creds *Credentials) interface{} {
	r.XMLName = xml.Name{
		Space: "urn:ebay:apis:eBLBaseComponents",
		Local: r.CallName(),
	}
	r.RequesterCredentials = &RequesterCredentials{EBayAuthToken: creds.AuthToken}
	return r
}

// ParseResponse retruns response data as EbayResponse object
func (r GetMyeBaySellingRequest) ParseResponse(resp []byte) (EbayResponse, error) {
	var xmlResponse GetMyeBaySellingResponse
	err := xml.Unmarshal(resp, &xmlResponse)

	return xmlResponse, err
}

// ResponseErrors returns errors
func (r GetMyeBaySellingResponse) ResponseErrors() EbayErrors {
	return r.ebayResponse.Errors
}

// Item type
type Item struct {
	SKU               string `xml:"SKU,omitempty"`
	StartPrice        *Price `xml:"StartPrice,omitempty"`
	ItemID            string `xml:"ItemID,omitempty"`
	Title             string `xml:"Title,omitempty"`
	QuantityAvailable int    `xml:"QuantityAvailable,omitempty"`
	SellingStatus     struct {
		CurrentPrice *Price `xml:"CurrentPrice,omitempty"`
	} `xml:"SellingStatus,omitempty"`
}

// Price struct
type Price struct {
	Amount     float64 `xml:",chardata"`
	CurrencyID string  `xml:"currencyID,attr,omitempty"`
}

// GetMyeBaySellingResponse type
type GetMyeBaySellingResponse struct {
	// MANY ADDITIONAL FIELDS AS REQUIRED: https://developer.ebay.com/devzone/xml/docs/reference/ebay/getmyebayselling.html
	ebayResponse
	XMLName    xml.Name           `xml:"GetMyeBaySellingResponse"`
	Xmlns      string             `xml:"xmlns,attr"`
	ActiveList ActiveListResponse `xml:"ActiveList"`

	// Summary struct {
	// 	ActiveAuctionCount   int `xml:"ActiveAuctionCount"`
	// 	AmountLimitRemaining struct {
	// 		Amount     float64 `xml:",chardata"`
	// 		CurrencyID string  `xml:"currencyID,attr"`
	// 	} `xml:"AmountLimitRemaining"`
	// 	AuctionBidCount          string `xml:"AuctionBidCount"`
	// 	AuctionSellingCount      string `xml:"AuctionSellingCount"`
	// 	ClassifiedAdCount        string `xml:"ClassifiedAdCount"`
	// 	ClassifiedAdOfferCount   string `xml:"ClassifiedAdOfferCount"`
	// 	QuantityLimitRemaining   string `xml:"QuantityLimitRemaining"`
	// 	SoldDurationInDays       string `xml:"SoldDurationInDays"`
	// 	TotalAuctionSellingValue struct {
	// 		Amount     float64 `xml:",chardata"`
	// 		CurrencyID string  `xml:"currencyID,attr"`
	// 	} `xml:"TotalAuctionSellingValue"`
	// 	TotalLeadCount         string `xml:"TotalLeadCount"`
	// 	TotalListingsWithLeads string `xml:"TotalListingsWithLeads"`
	// 	TotalSoldCount         string `xml:"TotalSoldCount"`
	// 	TotalSoldValue         struct {
	// 		Amount     float64 `xml:",chardata"`
	// 		CurrencyID string  `xml:"currencyID,attr"`
	// 	} `xml:"TotalSoldValue"`
	// } `xml:"Summary"`

	Build                 string `xml:"Build"`
	CorrelationID         string `xml:"CorrelationID"`
	HardExpirationWarning string `xml:"HardExpirationWarning"`
	Version               string `xml:"Version"`
}

// ActiveListResponse struct
type ActiveListResponse struct {
	ItemArray struct {
		Items []Item `xml:"Item"`
	} `xml:"ItemArray"`
	PaginationResult PaginationResult `xml:"PaginationResult"`
}
