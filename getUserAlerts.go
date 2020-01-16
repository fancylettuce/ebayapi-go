package ebayapi

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

//
// The 'GetUserAlerts' call from the ebay ClientAlerts API uses a REST querystring request and recieves a JSON response
//

// GetUserAlertsRequest holds credential data for Client Alerts API GetUserAlerts
type GetUserAlertsRequest struct {
	SessionID   string
	SessionData string
}

// CallName retruns the string name of this call
func (r GetUserAlertsRequest) CallName() string {
	return "GetUserAlerts"
}

// Body ataches credential and returns url querystring values
func (r GetUserAlertsRequest) Body(creds *Credentials) interface{} {
	return url.Values{
		"callname":    {r.CallName()},
		"SessionID":   {r.SessionID},
		"SessionData": {r.SessionData},
	}
}

// ParseResponse unmarshals response bytes into obj
func (r GetUserAlertsRequest) ParseResponse(resp []byte) (EbayResponse, error) {
	var response GetUserAlertsResponse
	//fmt.Printf("\n%s", string(resp))
	err := json.Unmarshal(resp, &response)
	return response, err
}

// Failure checks if call failed
func (r GetUserAlertsResponse) Failure() bool {
	return r.Ack == "Failure"
}

// ResponseErrors returns error array
func (r GetUserAlertsResponse) ResponseErrors() EbayErrors {
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

// GetUserAlertsResponse is Unmarshalled from JSON
type GetUserAlertsResponse struct {
	Timestamp time.Time `json:"Timestamp"`
	Ack       string    `json:"Ack"`
	Build     string    `json:"Build"`
	Version   string    `json:"Version"`
	Errors    []struct {
		ErrorClassification string `json:"ErrorClassification"`
		ErrorCode           string `json:"ErrorCode"`
		ErrorParameters     struct {
			Value string `json:"Value"`
		} `json:"ErrorParameters"`
		LongMessage  string `json:"LongMessage"`
		SeverityCode string `json:"SeverityCode"`
		ShortMessage string `json:"ShortMessage"`
	} `json:"Errors"`
	ClientAlerts struct {
		ClientAlertEvent []struct {
			EventType             string `json:"EventType"`
			FixedPriceTransaction struct {
				EventType    string    `json:"EventType"`
				Timestamp    time.Time `json:"Timestamp"`
				ItemID       string    `json:"ItemID"`
				BidCount     int       `json:"BidCount"`
				SellerUserID string    `json:"SellerUserID"`
				EndTime      time.Time `json:"EndTime"`
				CurrentPrice struct {
					Value      float64 `json:"Value"`
					CurrencyID string  `json:"CurrencyID"`
				} `json:"CurrentPrice"`
				Title       string `json:"Title"`
				GalleryURL  string `json:"GalleryURL"`
				Quantity    int    `json:"Quantity"`
				Transaction []struct {
					AmountPaid struct {
						Value      float64 `json:"Value"`
						CurrencyID string  `json:"CurrencyID"`
					} `json:"AmountPaid"`
					QuantitySold    int       `json:"QuantitySold"`
					BuyerUserID     string    `json:"BuyerUserID"`
					TransactionID   string    `json:"TransactionID"`
					CreatedDate     time.Time `json:"CreatedDate"`
					OrderLineItemID string    `json:"OrderLineItemID"`
					ContainingOrder struct {
						OrderID string `json:"OrderID"`
					}
				} `json:"Transaction"`
			} `json:"FixedPriceTransaction"`
		} `json:"ClientAlertEvent"`
	} `json:"ClientAlerts"`
	SessionData string `json:"SessionData"`
}

// GetUserAlertsResponse2 is Unmarshalled from JSON
type GetUserAlertsResponse2 struct {
	SessionData   string `json:"SessionData"`
	Ack           string `json:"Ack"`
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
	Timestamp    string `json:"Timestamp"`
	Version      string `json:"Version"`
	ClientAlerts struct {
		ClientAlertEvents []struct {
			EventType string `json:"EventType"`

			FixedPriceTransaction struct {
				CurrentPrice struct {
					Value      float64 `json:"Value"`
					CurrencyID string  `json:"CurrencyID"`
				} `json:"CurrentPrice"`
				EndTime      string `json:"EndTime"`
				EventType    string `json:"EventType"`
				GalleryURL   string `json:"GalleryURL"`
				ItemID       string `json:"ItemID"`
				SellerUserID string `json:"SellerUserID"`
				Timestamp    string `json:"Timestamp"`
				Title        string `json:"Title"`
				Transaction  struct {
					AmountPaid struct {
						Value      float64 `json:"Value"`
						CurrencyID string  `json:"CurrencyID"`
					} `json:"AmountPaid"`
					BuyerUserID     string `json:"BuyerUserID"`
					CreatedDate     string `json:"CreatedDate"`
					OrderLineItemID string `json:"OrderLineItemID"`
					QuantitySold    int    `json:"QuantitySold"`
					TransactionID   string `json:"TransactionID"`
				} `json:"Transaction"`
			} `json:"FixedPriceTransaction"`

			FeedbackReceived struct {
				EventType      string `json:"EventType"`
				FeedbackDetail struct {
					CommentingUser string `json:"CommentingUser"`
					CommentText    string `json:"CommentText"`
					CommentType    string `json:"CommentType"`
					FeedbackID     string `json:"FeedbackID"`
					FeedbackScore  int    `json:"FeedbackScore"`
					ItemID         string `json:"ItemID"`
					ItemPrice      struct {
						Value      float64 `json:"Value"`
						CurrencyID string  `json:"CurrencyID"`
					} `json:"ItemPrice"`
					ItemTitle     string `json:"ItemTitle"`
					Role          string `json:"Role"`
					TransactionID string `json:"TransactionID"`
				} `json:"FeedbackDetail"`
				Timestamp string `json:"Timestamp"`
			} `json:"FeedbackReceived"`

			ItemListed struct {
				CurrentPrice struct {
					Value      float64 `json:"Value"`
					CurrencyID string  `json:"CurrencyID"`
				} `json:"CurrentPrice"`
				EndTime      string `json:"EndTime"`
				EventType    string `json:"EventType"`
				GalleryURL   string `json:"GalleryURL"`
				ItemID       string `json:"ItemID"`
				Quantity     int    `json:"Quantity"`
				SellerUserID string `json:"SellerUserID"`
				Timestamp    string `json:"Timestamp"`
				Title        string `json:"Title"`
			} `json:"ItemListed"`

			ItemMarkedShipped struct {
				EventType    string `json:"EventType"`
				ItemID       string `json:"ItemID"`
				OrderID      string `json:"OrderID"`
				SellerUserID string `json:"SellerUserID"`
				Shipment     struct {
					ShipmentTrackingNumber string `json:"ShipmentTrackingNumber"`
					ShippingCarrierUsed    string `json:"ShippingCarrierUsed"`
				} `json:"Shipment"`
				Timestamp     string `json:"Timestamp"`
				Title         string `json:"Title"`
				TransactionID string `json:"TransactionID"`
			} `json:"ItemMarkedShipped"`

			// ADDITIONAL ALERTS THAT MUST BE SUBCRIBED TO RECEIVE:

			// FixedPriceTransaction struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	GalleryURL   string `json:"GalleryURL"`
			// 	ItemID       string `json:"ItemID"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// 	Transaction  []struct {
			// 		AmountPaid struct {
			// 			Value      float64 `json:"Value"`
			// 			CurrencyID string  `json:"CurrencyID"`
			// 		} `json:"AmountPaid"`
			// 		BuyerUserID     string `json:"BuyerUserID"`
			// 		CreatedDate     string `json:"CreatedDate"`
			// 		OrderLineItemID string `json:"OrderLineItemID"`
			// 		QuantitySold    int    `json:"QuantitySold"`
			// 		TransactionID   string `json:"TransactionID"`
			// 	} `json:"Transaction"`
			// } `json:"FixedPriceTransaction"`

			// AskSellerQuestion struct {
			// 	EventType   string `json:"EventType"`
			// 	ItemID      string `json:"ItemID"`
			// 	MessageID   string `json:"MessageID"`
			// 	MessageType string `json:"MessageType"`
			// 	Timestamp   string `json:"Timestamp"`
			// 	Title       string `json:"Title"`
			// } `json:"AskSellerQuestion"`

			// BestOffer struct {
			// 	BestOffer struct {
			// 		BestOfferCodeType string `json:"BestOfferCodeType"`
			// 		BestOfferID       string `json:"BestOfferID"`
			// 		BuyerMessage      string `json:"BuyerMessage"`
			// 		BuyerUserID       string `json:"BuyerUserID"`
			// 		ExpirationTime    string `json:"ExpirationTime"`
			// 		Price             struct {
			// 			Value      float64 `json:"Value"`
			// 			CurrencyID string  `json:"CurrencyID"`
			// 		} `json:"Price"`
			// 		Quantity      int    `json:"Quantity"`
			// 		SellerMessage string `json:"SellerMessage"`
			// 		Status        string `json:"Status"`
			// 	} `json:"BestOffer"`
			// 	EventType string `json:"EventType"`
			// 	ItemID    string `json:"ItemID"`
			// 	Timestamp string `json:"Timestamp"`
			// } `json:"BestOffer"`

			// BestOfferDeclined struct {
			// 	BestOffer struct {
			// 		BestOfferCodeType string `json:"BestOfferCodeType"`
			// 		BestOfferID       string `json:"BestOfferID"`
			// 		BuyerMessage      string `json:"BuyerMessage"`
			// 		BuyerUserID       string `json:"BuyerUserID"`
			// 		ExpirationTime    string `json:"ExpirationTime"`
			// 		Price             struct {
			// 			Value      float64 `json:"Value"`
			// 			CurrencyID string  `json:"CurrencyID"`
			// 		} `json:"Price"`
			// 		Quantity      int    `json:"Quantity"`
			// 		SellerMessage string `json:"SellerMessage"`
			// 		Status        string `json:"Status"`
			// 	} `json:"BestOffer"`
			// 	EventType string `json:"EventType"`
			// 	ItemID    string `json:"ItemID"`
			// 	Timestamp string `json:"Timestamp"`
			// } `json:"BestOfferDeclined"`

			// BestOfferPlaced struct {
			// 	BestOffer struct {
			// 		BestOfferCodeType string `json:"BestOfferCodeType"`
			// 		BestOfferID       string `json:"BestOfferID"`
			// 		BuyerMessage      string `json:"BuyerMessage"`
			// 		BuyerUserID       string `json:"BuyerUserID"`
			// 		ExpirationTime    string `json:"ExpirationTime"`
			// 		Price             struct {
			// 			Value      float64 `json:"Value"`
			// 			CurrencyID string  `json:"CurrencyID"`
			// 		} `json:"Price"`
			// 		Quantity      int    `json:"Quantity"`
			// 		SellerMessage string `json:"SellerMessage"`
			// 		Status        string `json:"Status"`
			// 	} `json:"BestOffer"`
			// 	EventType string `json:"EventType"`
			// 	ItemID    string `json:"ItemID"`
			// 	Timestamp string `json:"Timestamp"`
			// } `json:"BestOfferPlaced"`

			// BidPlaced struct {
			// 	BidCount          int  `json:"BidCount"`
			// 	BuyItNowAvailable bool `json:"BuyItNowAvailable"`
			// 	CurrentPrice      struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime          string `json:"EndTime"`
			// 	EventType        string `json:"EventType"`
			// 	GalleryURL       string `json:"GalleryURL"`
			// 	HighBidderUserID string `json:"HighBidderUserID"`
			// 	ItemID           string `json:"ItemID"`
			// 	Quantity         int    `json:"Quantity"`
			// 	ReserveMet       bool   `json:"ReserveMet"`
			// 	SellerUserID     string `json:"SellerUserID"`
			// 	Timestamp        string `json:"Timestamp"`
			// 	Title            string `json:"Title"`
			// } `json:"BidPlaced"`

			// BidReceived struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime          string `json:"EndTime"`
			// 	EventType        string `json:"EventType"`
			// 	GalleryURL       string `json:"GalleryURL"`
			// 	HighBidderUserID string `json:"HighBidderUserID"`
			// 	ItemID           string `json:"ItemID"`
			// 	Quantity         int    `json:"Quantity"`
			// 	SellerUserID     string `json:"SellerUserID"`
			// 	Timestamp        string `json:"Timestamp"`
			// 	Title            string `json:"Title"`
			// } `json:"BidReceived"`

			// CounterOfferReceived struct {
			// 	BestOffer struct {
			// 		BestOfferCodeType string `json:"BestOfferCodeType"`
			// 		BestOfferID       string `json:"BestOfferID"`
			// 		BuyerMessage      string `json:"BuyerMessage"`
			// 		BuyerUserID       string `json:"BuyerUserID"`
			// 		ExpirationTime    string `json:"ExpirationTime"`
			// 		Price             struct {
			// 			Value      float64 `json:"Value"`
			// 			CurrencyID string  `json:"CurrencyID"`
			// 		} `json:"Price"`
			// 		Quantity      int    `json:"Quantity"`
			// 		SellerMessage string `json:"SellerMessage"`
			// 		Status        string `json:"Status"`
			// 	} `json:"BestOffer"`
			// 	EventType string `json:"EventType"`
			// 	ItemID    string `json:"ItemID"`
			// 	Timestamp string `json:"Timestamp"`
			// } `json:"CounterOfferReceived"`

			// EndOfAuction struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	GalleryURL   string `json:"GalleryURL"`
			// 	ItemID       string `json:"ItemID"`
			// 	Quantity     int    `json:"Quantity"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// 	Transaction  struct {
			// 		AmountPaid struct {
			// 			Value      float64 `json:"Value"`
			// 			CurrencyID string  `json:"CurrencyID"`
			// 		} `json:"AmountPaid"`
			// 		BuyerUserID     string `json:"BuyerUserID"`
			// 		CreatedDate     string `json:"CreatedDate"`
			// 		OrderLineItemID string `json:"OrderLineItemID"`
			// 		QuantitySold    int    `json:"QuantitySold"`
			// 		TransactionID   string `json:"TransactionID"`
			// 	} `json:"Transaction"`
			// } `json:"EndOfAuction"`

			// FeedbackLeft struct {
			// 	EventType      string `json:"EventType"`
			// 	FeedbackDetail struct {
			// 		CommentingUser string `json:"CommentingUser"`
			// 		CommentText    string `json:"CommentText"`
			// 		CommentType    string `json:"CommentType"`
			// 		FeedbackID     string `json:"FeedbackID"`
			// 		FeedbackScore  int    `json:"FeedbackScore"`
			// 		ItemID         string `json:"ItemID"`
			// 		ItemPrice      struct {
			// 			Value      float64 `json:"Value"`
			// 			CurrencyID string  `json:"CurrencyID"`
			// 		} `json:"ItemPrice"`
			// 		ItemTitle     string `json:"ItemTitle"`
			// 		Role          string `json:"Role"`
			// 		TransactionID string `json:"TransactionID"`
			// 	} `json:"FeedbackDetail"`
			// 	Timestamp string `json:"Timestamp"`
			// } `json:"FeedbackLeft"`

			// FeedbackStarChanged struct {
			// 	EventType string `json:"EventType"`
			// 	Timestamp string `json:"Timestamp"`
			// 	User      struct {
			// 		FeedbackRatingStar      string  `json:"FeedbackRatingStar"`
			// 		PositiveFeedbackPercent float64 `json:"PositiveFeedbackPercent"`
			// 		UserID                  string  `json:"UserID"`
			// 	} `json:"User"`
			// } `json:"FeedbackStarChanged"`

			// ItemAddedToWatchList struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	GalleryURL   string `json:"GalleryURL"`
			// 	ItemID       string `json:"ItemID"`
			// 	Quantity     int    `json:"Quantity"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// } `json:"ItemAddedToWatchList"`

			// ItemEnded struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	ItemID       string `json:"ItemID"`
			// 	Quantity     int    `json:"Quantity"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// } `json:"ItemEnded"`

			// ItemLost struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	GalleryURL   string `json:"GalleryURL"`
			// 	ItemID       string `json:"ItemID"`
			// 	Quantity     int    `json:"Quantity"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// } `json:"ItemLost"`

			// ItemMarkedPaid struct {
			// 	EventType     string `json:"EventType"`
			// 	ItemID        string `json:"ItemID"`
			// 	OrderID       string `json:"OrderID"`
			// 	SellerUserID  string `json:"SellerUserID"`
			// 	Timestamp     string `json:"Timestamp"`
			// 	Title         string `json:"Title"`
			// 	TransactionID string `json:"TransactionID"`
			// } `json:"ItemMarkedPaid"`

			// ItemRemovedFromWatchList struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	GalleryURL   string `json:"GalleryURL"`
			// 	ItemID       string `json:"ItemID"`
			// 	Quantity     int    `json:"Quantity"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// } `json:"ItemRemovedFromWatchList"`

			// ItemSold struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	GalleryURL   string `json:"GalleryURL"`
			// 	ItemID       string `json:"ItemID"`
			// 	Quantity     int    `json:"Quantity"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// } `json:"ItemSold"`

			// ItemUnsold struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	GalleryURL   string `json:"GalleryURL"`
			// 	ItemID       string `json:"ItemID"`
			// 	Quantity     int    `json:"Quantity"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// } `json:"ItemUnsold"`

			// ItemWon struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime      string `json:"EndTime"`
			// 	EventType    string `json:"EventType"`
			// 	GalleryURL   string `json:"GalleryURL"`
			// 	ItemID       string `json:"ItemID"`
			// 	Quantity     int    `json:"Quantity"`
			// 	SellerUserID string `json:"SellerUserID"`
			// 	Timestamp    string `json:"Timestamp"`
			// 	Title        string `json:"Title"`
			// } `json:"ItemWon"`

			// Outbid struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime             string `json:"EndTime"`
			// 	EventType           string `json:"EventType"`
			// 	GalleryURL          string `json:"GalleryURL"`
			// 	HighBidderEIASToken string `json:"HighBidderEIASToken"`
			// 	HighBidderUserID    string `json:"HighBidderUserID"`
			// 	ItemID              string `json:"ItemID"`
			// 	SellerUserID        string `json:"SellerUserID"`
			// 	Timestamp           string `json:"Timestamp"`
			// 	Title               string `json:"Title"`
			// } `json:"Outbid"`

			// SecondChanceOffer struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime                    string `json:"EndTime"`
			// 	EventType                  string `json:"EventType"`
			// 	GalleryURL                 string `json:"GalleryURL"`
			// 	ItemID                     string `json:"ItemID"`
			// 	Quantity                   int    `json:"Quantity"`
			// 	SecondChanceOriginalItemID string `json:"SecondChanceOriginalItemID"`
			// 	SellerUserID               string `json:"SellerUserID"`
			// 	Timestamp                  string `json:"Timestamp"`
			// 	Title                      string `json:"Title"`
			// } `json:"SecondChanceOffer"`

			// WatchedItemEndingSoon struct {
			// 	BidCount     int `json:"BidCount"`
			// 	CurrentPrice struct {
			// 		Value      float64 `json:"Value"`
			// 		CurrencyID string  `json:"CurrencyID"`
			// 	} `json:"CurrentPrice"`
			// 	EndTime          string `json:"EndTime"`
			// 	EventType        string `json:"EventType"`
			// 	GalleryURL       string `json:"GalleryURL"`
			// 	HighBidderUserID string `json:"HighBidderUserID"`
			// 	ItemID           string `json:"ItemID"`
			// 	SellerUserID     string `json:"SellerUserID"`
			// 	Timestamp        string `json:"Timestamp"`
			// 	ViewItemURL      string `json:"ViewItemURL"`
			// } `json:"WatchedItemEndingSoon"`

		} `json:"ClientAlertEvent"`
	} `json:"ClientAlerts"`
}
