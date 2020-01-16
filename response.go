package ebayapi

import "time"

// EbayResponse interface for defining response types
type EbayResponse interface {
	Failure() bool
	ResponseErrors() EbayErrors
}

type ebayResponse struct {
	Timestamp time.Time
	Ack       string
	Errors    []ebayResponseError
}

func (r ebayResponse) Failure() bool {
	return r.Ack == "Failure"
}

func (r ebayResponse) ResponseErrors() EbayErrors {
	return r.Errors
}

type ebayResponseError struct {
	ShortMessage        string
	LongMessage         string
	ErrorCode           int
	SeverityCode        string
	ErrorClassification string
}
