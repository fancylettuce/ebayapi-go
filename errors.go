package ebayapi

import (
	"fmt"
	"strings"
)

// EbayErrors holds and handles ebay API errors
type EbayErrors []ebayResponseError

func (err EbayErrors) Error() string {
	var errs []string

	for _, e := range err {
		errs = append(errs, fmt.Sprintf("%#v", e))
	}

	return strings.Join(errs, ", ")
}

// RevisionError handles listing revision errors
func (err EbayErrors) RevisionError() bool {
	for _, err := range err {
		if err.ErrorCode == 10039 || err.ErrorCode == 10029 || err.ErrorCode == 21916916 || err.ErrorCode == 21916923 || err.ErrorCode == 21919028 {
			return true
		}
	}

	return false
}

// ListingEnded handles ebay API errors
func (err EbayErrors) ListingEnded() bool {
	for _, err := range err {
		if err.ErrorCode == 291 || err.ErrorCode == 240 {
			return true
		}
	}

	return false
}

// ListingDeleted handles ebay API errors
func (err EbayErrors) ListingDeleted() bool {
	for _, err := range err {
		if err.ErrorCode == 17 {
			return true
		}
	}

	return false
}

type httpError struct {
	statusCode int
	body       []byte
}

func (err httpError) Error() string {
	return fmt.Sprintf("%d - %s", err.statusCode, err.body)
}
