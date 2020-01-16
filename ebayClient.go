package ebayapi

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Credentials holds user api token data
type Credentials struct {
	DevID,
	AppID,
	CertID,
	AuthToken string
}

// EbayClient enacts SOAP requests to the ebay API
type EbayClient struct {
	Credentials
	baseURL string
	SiteID  int
	logger  *logrus.Logger
}

// NewSandboxClient initializes EbayClient for the 'sandbox' dev environment/markets
func NewSandboxClient(log *logrus.Logger) *EbayClient {
	return &EbayClient{
		baseURL: "https://api.sandbox.ebay.com",
		logger:  log,
	}
}

// NewProductionClient initializes EbayClient for the production environment/markets
func NewProductionClient(log *logrus.Logger) *EbayClient {
	return &EbayClient{
		baseURL: "https://api.ebay.com",
		logger:  log,
	}
}

func (e *EbayClient) parseSOAPrequest(c Call) (*http.Request, error) {
	ec := ebayRequest{
		creds:   e.Credentials,
		command: c,
	}

	body := new(bytes.Buffer)
	body.Write([]byte(xml.Header))
	err := xml.NewEncoder(body).Encode(ec)

	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/ws/api.dll", e.baseURL),
		body,
	)

	req.Header.Add("X-EBAY-API-DEV-NAME", e.DevID)
	req.Header.Add("X-EBAY-API-APP-NAME", e.AppID)
	req.Header.Add("X-EBAY-API-CERT-NAME", e.CertID)
	req.Header.Add("X-EBAY-API-CALL-NAME", c.CallName())
	req.Header.Add("X-EBAY-API-SITEID", strconv.Itoa(e.SiteID))
	req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", strconv.Itoa(1113))
	req.Header.Add("Content-Type", "text/xml")

	if e.logger != nil {
		xmlString, err := PrettPrintXML(body.Bytes())
		if err == nil {
			fmt.Println()
			e.logger.Debug("REQUESTING: ", c.CallName())
			fmt.Printf("%s\n", xmlString)
		} else {
			e.logger.Debug("FAILED PrettyPrintXML REQUEST: ", err)
		}
	}

	return req, nil
}

// DoSOAPcall makes the requested XML request to the ebay API
// EbayResponse must then be typecast into the correct response type
func (e *EbayClient) DoSOAPcall(ctx context.Context, c Call) (EbayResponse, error) {
	req, err := e.parseSOAPrequest(c)
	if err != nil {
		return ebayResponse{}, err
	}

	req.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(req)

	if urlErr, ok := err.(*url.Error); ok {
		return ebayResponse{}, urlErr
	} else if resp.StatusCode != 200 {
		httpErr := httpError{
			statusCode: resp.StatusCode,
		}
		httpErr.body, _ = ioutil.ReadAll(resp.Body)

		return ebayResponse{}, httpErr
	}

	bodyContents, _ := ioutil.ReadAll(resp.Body)
	defer func() {
		cerr := resp.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	if e.logger != nil {
		xmlString, err := PrettPrintXML(bodyContents)
		if err == nil {
			fmt.Println()
			e.logger.Debug("RESPONSE: ", c.CallName())
			fmt.Printf("%s\n\n", xmlString)
		} else {
			e.logger.Debug("FAILED PrettyPrintXML RESPONSE: ", err)
		}
	}

	response, err := c.ParseResponse(bodyContents)

	if response.Failure() {
		return response, EbayErrors(response.ResponseErrors())
	}

	return response, err
}

// ParseRESTrequest prepares a request for REST calling
func (e *EbayClient) parseRESTrequest(endpoint string, req Call, method string) (*http.Request, error) {
	reqURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	reqURL.RawQuery = req.Body(&e.Credentials).(url.Values).Encode()
	httpReq, err := http.NewRequest(method, reqURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if e.logger != nil {
		e.logger.Debug("REQUESTING: ", req.CallName(), ": ", httpReq.URL.String())
	}

	return httpReq, nil
}

// DoRESTcall performs request with data encoded into the URL querystring
func (e *EbayClient) DoRESTcall(ctx context.Context, endpoint string, req Call, method string) (EbayResponse, error) {
	httpReq, err := e.parseRESTrequest(endpoint, req, method)
	if err != nil {
		return nil, err
	}

	httpReq.WithContext(ctx)
	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}

	bodyContents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer func() {
		cerr := resp.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()

	if e.logger != nil {
		e.logger.Debug("RESPONSE: ", req.CallName(), ": ", string(bodyContents))
	}

	response, err := req.ParseResponse(bodyContents)
	if err != nil {
		return nil, err
	}

	return response, nil
}
