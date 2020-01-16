package ebayapi

import (
	"encoding/xml"
	"time"
)

// GetOrdersRequest type
type GetOrdersRequest struct {
	XMLName              xml.Name
	RequesterCredentials *RequesterCredentials
	CreateTimeFrom       *time.Time    `xml:",omitempty"`
	CreateTimeTo         *time.Time    `xml:",omitempty"`
	IncludeFinalValueFee bool          `xml:",omitempty"`
	ModTimeFrom          *time.Time    `xml:",omitempty"`
	ModTimeTo            *time.Time    `xml:",omitempty"`
	NumberOfDays         int           `xml:",omitempty"`
	OrderIDArray         *OrderIDArray `xml:",omitempty"`
	OrderStatus          string        `xml:",omitempty"`
	Pagination           *Pagination   `xml:",omitempty"`
	SortingOrder         string        `xml:",omitempty"`
	DetailLevel          []string      `xml:",omitempty"`
	ErrorLanguage        string        `xml:",omitempty"`
	MessageID            string        `xml:",omitempty"`
	OutputSelector       []string      `xml:",omitempty"`
	Version              string        `xml:",omitempty"`
	WarningLevel         string        `xml:",omitempty"`
}

// OrderIDArray struct
type OrderIDArray struct {
	OrderIDs []string `xml:"OrderID,omitempty"`
}

// DefaultOutputSelection selects only order data fields required for order entry
func (c GetOrdersRequest) DefaultOutputSelection() *GetOrdersRequest {
	c.DetailLevel = []string{"ReturnAll"}
	c.OutputSelector = []string{
		"HasMoreOrders",
		"PageNumber",
		"PaginationResult",
		"OrderArray.Order.OrderID",
		"OrderArray.Order.OrderStatus",
		"OrderArray.Order.ShippedTime",
		"OrderArray.Order.TransactionArray.Transaction.Buyer.UserFirstName",
		"OrderArray.Order.TransactionArray.Transaction.Buyer.UserLastName",
		"OrderArray.Order.TransactionArray.Transaction.Buyer.Email",
		"OrderArray.Order.MonetaryDetails.Payments.Payment.ReferenceID",
		"OrderArray.Order.ShippingDetails.SellingManagerSalesRecordNumber",
		"OrderArray.Order.Total",
		"OrderArray.Order.ShippingAddress.Name",
		"OrderArray.Order.ShippingAddress.Street1",
		"OrderArray.Order.ShippingAddress.Street2",
		"OrderArray.Order.ShippingAddress.CityName",
		"OrderArray.Order.ShippingAddress.StateOrProvince",
		"OrderArray.Order.ShippingAddress.PostalCode",
		"OrderArray.Order.ShippingAddress.Phone",
		"OrderArray.Order.TransactionArray.Transaction.Item.SKU",
		"OrderArray.Order.TransactionArray.Transaction.TransactionID",
		"OrderArray.Order.TransactionArray.Transaction.QuantityPurchased",
		"OrderArray.Order.MonetaryDetails.Payments.Payment.ReferenceID",
	}
	return &c
}

// CallName returns name of call
func (c GetOrdersRequest) CallName() string {
	return "GetOrders"
}

// Body ataches credential and returns XML body
func (c GetOrdersRequest) Body(creds *Credentials) interface{} {
	c.XMLName = xml.Name{
		Space: "urn:ebay:apis:eBLBaseComponents",
		Local: c.CallName(),
	}
	c.RequesterCredentials = &RequesterCredentials{EBayAuthToken: creds.AuthToken}
	return c
}

// ParseResponse retruns response data as EbayResponse object
func (c GetOrdersRequest) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse GetOrdersResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

// ResponseErrors returns errors
func (r GetOrdersResponse) ResponseErrors() EbayErrors {
	return r.ebayResponse.Errors
}

// Pagination struct
type Pagination struct {
	EntriesPerPage int `xml:",omitempty"`
	PageNumber     int `xml:",omitempty"`
}

// PaginationResult holds page totals
type PaginationResult struct {
	TotalNumberOfPages   int
	TotalNumberOfEntries int
}

// GetOrdersResponse obj
type GetOrdersResponse struct {
	ebayResponse
	XMLName          xml.Name `xml:"GetOrdersResponse"`
	Xmlns            string   `xml:"xmlns,attr"`
	Version          string
	Build            string
	PaginationResult PaginationResult
	HasMoreOrders    bool
	OrderArray       struct {
		Orders []Order `xml:"Order"`
	} `xml:"OrderArray"`
	OrdersPerPage            int
	PageNumber               int
	ReturnedOrderCountActual int
}

// Order ebay api type - uncomment fields as needed
type Order struct {
	OrderID         string
	OrderStatus     string
	BuyerUserID     string
	CreatedTime     time.Time
	ShippedTime     *time.Time `xml:",omitempty"`
	MonetaryDetails struct {
		Payments struct {
			Payment struct {
				ReferenceID string `xml:"ReferenceID"`
				// FeeOrCreditAmount struct {
				// 	Amount     float64 `xml:",chardata"`
				// 	CurrencyID string  `xml:"currencyID,attr"`
				// } `xml:"FeeOrCreditAmount"`
				// Payee         string `xml:"Payee"`
				// Payer         string `xml:"Payer"`
				// PaymentAmount struct {
				// 	Amount     float64 `xml:",chardata"`
				// 	CurrencyID string  `xml:"currencyID,attr"`
				// } `xml:"PaymentAmount"`
				// PaymentReferenceID string    `xml:"PaymentReferenceID"`
				// PaymentStatus      string    `xml:"PaymentStatus"`
				// PaymentTime        time.Time `xml:"PaymentTime"`
			} `xml:"Payment"`
		} `xml:"Payments"`
		// Refunds struct {
		// 	Refund struct {
		// 		FeeOrCreditAmount struct {
		// 			Amount     float64 `xml:",chardata"`
		// 			CurrencyID string  `xml:"currencyID,attr"`
		// 		} `xml:"FeeOrCreditAmount"`
		// 		ReferenceID  string `xml:"ReferenceID"`
		// 		RefundAmount struct {
		// 			Amount     float64 `xml:",chardata"`
		// 			CurrencyID string  `xml:"currencyID,attr"`
		// 		} `xml:"RefundAmount"`
		// 		RefundStatus string `xml:"RefundStatus"`
		// 		RefundTime   string `xml:"RefundTime"`
		// 		RefundTo     string `xml:"RefundTo"`
		// 		RefundType   string `xml:"RefundType"`
		// 	} `xml:"Refund"`
		// } `xml:"Refunds"`
	} `xml:"MonetaryDetails"`
	ShippingDetails struct {
		SellingManagerSalesRecordNumber string
		// SalesTax struct {
		// 	ShippingIncludedInTax bool
		// 	SalesTaxPercent       float64
		// 	SalesTaxState         string
		// 	SalesTaxAmount        struct {
		// 		Amount     float64 `xml:",chardata"`
		// 		CurrencyID string  `xml:"currencyID,attr"`
		// 	} `xml:"SalesTaxAmount"`
		// } `xml:"SalesTax"`
		// ShippingServiceOptions struct {
		// 	ShippingService     string
		// 	ShippingServiceCost struct {
		// 		Amount     float64 `xml:",chardata"`
		// 		CurrencyID string  `xml:"currencyID,attr"`
		// 	} `xml:"ShippingServiceCost"`
		// 	ShippingServicePriority int
		// 	ExpeditedService        bool
		// 	ShippingTimeMin         int
		// 	ShippingTimeMax         int
		// } `xml:"ShippingServiceOptions"`
		//GetItFast                       bool
	} `xml:"ShippingDetails"`
	ShippingAddress struct {
		Name              string
		Street1           string
		Street2           string
		CityName          string
		StateOrProvince   string
		Country           string
		CountryName       string
		Phone             string
		PostalCode        string
		AddressID         string
		AddressOwner      string
		ExternalAddressID string
	} `xml:"ShippingAddress"`
	Total struct {
		Amount     float64 `xml:",chardata"`
		CurrencyID string  `xml:"currencyID,attr"`
	} `xml:"Total"`
	TransactionArray struct {
		Transaction []struct {
			TransactionID string
			Buyer         struct {
				Email         string
				UserFirstName string
				UserLastName  string
				// VATStatus     string
			} `xml:"Buyer"`
			ShippingDetails struct {
				SellingManagerSalesRecordNumber int
			} `xml:"ShippingDetails"`
			CreatedDate time.Time
			Item        struct {
				ItemID               string
				Site                 string
				Title                string
				SKU                  string
				ConditionID          int
				ConditionDisplayName string
			} `xml:"Item"`
			QuantityPurchased int32
			OrderLineItemID   string
			// Status            struct {
			// 	PaymentHoldStatus string
			// 	InquiryStatus     string
			// 	ReturnStatus      string
			// } `xml:"Status"`
			// TransactionPrice struct {
			// 	Amount     float64 `xml:",chardata"`
			// 	CurrencyID string  `xml:"currencyID,attr"`
			// } `xml:"TransactionPrice"`
			// EBayCollectAndRemitTax  bool `xml:"eBayCollectAndRemitTax"`
			// ShippingServiceSelected struct {
			// 	ShippingPackageInfo struct {
			// 		EstimatedDeliveryTimeMin time.Time
			// 		EstimatedDeliveryTimeMax time.Time
			// 		HandleByTime             time.Time
			// 	} `xml:"ShippingPackageInfo"`
			// } `xml:"ShippingServiceSelected"`
			// TransactionSiteID        string
			// Platform                 string
			// EBayCollectAndRemitTaxes struct {
			// 	TotalTaxAmount struct {
			// 		Amount     float64 `xml:",chardata"`
			// 		CurrencyID string  `xml:"currencyID,attr"`
			// 	} `xml:"TotalTaxAmount"`
			// 	TaxDetails []struct {
			// 		Imposition     string
			// 		TaxDescription string
			// 		TaxAmount      struct {
			// 			Amount     float64 `xml:",chardata"`
			// 			CurrencyID string  `xml:"currencyID,attr"`
			// 		} `xml:"TaxAmount"`
			// 		TaxOnSubtotalAmount struct {
			// 			Amount     float64 `xml:",chardata"`
			// 			CurrencyID string  `xml:"currencyID,attr"`
			// 		} `xml:"TaxOnSubtotalAmount"`
			// 		TaxOnShippingAmount struct {
			// 			Amount     float64 `xml:",chardata"`
			// 			CurrencyID string  `xml:"currencyID,attr"`
			// 		} `xml:"TaxOnShippingAmount"`
			// 		TaxOnHandlingAmount struct {
			// 			Amount     float64 `xml:",chardata"`
			// 			CurrencyID string  `xml:"currencyID,attr"`
			// 		} `xml:"TaxOnHandlingAmount"`
			// 	} `xml:"TaxDetails"`
			// } `xml:"eBayCollectAndRemitTaxes"`
			// ActualShippingCost struct {
			// 	Amount     float64 `xml:",chardata"`
			// 	CurrencyID string  `xml:"currencyID,attr"`
			// } `xml:"ActualShippingCost"`
			// ActualHandlingCost struct {
			// 	Amount     float64 `xml:",chardata"`
			// 	CurrencyID string  `xml:"currencyID,attr"`
			// } `xml:"ActualHandlingCost"`
			// ExtendedOrderID     string
			ExternalTransaction struct {
				ExternalTransactionID     string `xml:"ExternalTransactionID"`
				ExternalTransactionStatus string `xml:"ExternalTransactionStatus"`
				ExternalTransactionTime   string `xml:"ExternalTransactionTime"`
				FeeOrCreditAmount         struct {
					Amount     string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"FeeOrCreditAmount"`
				PaymentOrRefundAmount struct {
					Amount     string `xml:",chardata"`
					CurrencyID string `xml:"currencyID,attr"`
				} `xml:"PaymentOrRefundAmount"`
			} `xml:"ExternalTransaction"`
			// EBayPlusTransaction bool `xml:"eBayPlusTransaction"`
			// GuaranteedShipping  bool
			// GuaranteedDelivery  bool
			// Taxes               struct {
			// 	TotalTaxAmount struct {
			// 		Amount     float64 `xml:",chardata"`
			// 		CurrencyID string  `xml:"currencyID,attr"`
			// 	} `xml:"TotalTaxAmount"`
			// 	TaxDetails []struct {
			// 		Imposition     string
			// 		TaxDescription string
			// 		TaxAmount      struct {
			// 			Amount     float64 `xml:",chardata"`
			// 			CurrencyID string  `xml:"currencyID,attr"`
			// 		} `xml:"TaxAmount"`
			// 		TaxOnSubtotalAmount struct {
			// 			Amount     float64 `xml:",chardata"`
			// 			CurrencyID string  `xml:"currencyID,attr"`
			// 		} `xml:"TaxOnSubtotalAmount"`
			// 		TaxOnShippingAmount struct {
			// 			Amount     float64 `xml:",chardata"`
			// 			CurrencyID string  `xml:"currencyID,attr"`
			// 		} `xml:"TaxOnShippingAmount"`
			// 		TaxOnHandlingAmount struct {
			// 			Amount     float64 `xml:",chardata"`
			// 			CurrencyID string  `xml:"currencyID,attr"`
			// 		} `xml:"TaxOnHandlingAmount"`
			// 	} `xml:"TaxDetails"`
			// } `xml:"Taxes"`
		} `xml:"Transaction"`
	} `xml:"TransactionArray"`
	// PaymentMethods  string
	// SellerEmail     string
	// ShippingServiceSelected struct {
	// 	ShippingService     string
	// 	ShippingServiceCost struct {
	// 		Amount     float64 `xml:",chardata"`
	// 		CurrencyID string  `xml:"currencyID,attr"`
	// 	} `xml:"ShippingServiceCost"`
	// } `xml:"ShippingServiceSelected"`
	// Subtotal struct {
	// 	Amount     float64 `xml:",chardata"`
	// 	CurrencyID string  `xml:"currencyID,attr"`
	// } `xml:"Subtotal"`
	// EBayCollectAndRemitTax bool `xml:"eBayCollectAndRemitTax"`
	// PaidTime                            time.Time
	// IntegratedMerchantCreditCardEnabled bool
	// EIASToken                           string
	// PaymentHoldStatus                   string
	// IsMultiLegShipping                  bool
	// SellerUserID                        string
	// SellerEIASToken                     string
	// CancelStatus                        string
	// ExtendedOrderID                     string
	// ContainseBayPlusTransaction         bool
	// AdjustmentAmount struct {
	// 	Amount     float64 `xml:",chardata"`
	// 	CurrencyID string  `xml:"currencyID,attr"`
	// } `xml:"AdjustmentAmount"`
	// AmountPaid struct {
	// 	Amount     float64 `xml:",chardata"`
	// 	CurrencyID string  `xml:"currencyID,attr"`
	// } `xml:"AmountPaid"`
	// AmountSaved struct {
	// 	Amount     float64 `xml:",chardata"`
	// 	CurrencyID string  `xml:"currencyID,attr"`
	// } `xml:"AmountSaved"`
	// CheckoutStatus struct {
	// 	EBayPaymentStatus                   string `xml:"eBayPaymentStatus"`
	// 	LastModifiedTime                    time.Time
	// 	PaymentMethod                       string
	// 	Status                              string
	// 	IntegratedMerchantCreditCardEnabled bool
	// 	PaymentInstrument                   string
	// }
}
