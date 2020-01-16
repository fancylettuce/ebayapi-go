package ebayapi

import (
	"context"

	"github.com/sirupsen/logrus"
)

// ClientAlertsAPI enacts requests to the ebay Client Alerts API
type ClientAlertsAPI struct {
	clientAlertsURL        string
	clientAlertsSandboxURL string
	client                 *EbayClient
	logger                 *logrus.Logger

	clientAlertsAuthToken string
	sessionID             string
	sessionData           string
}

// NewClientAlertsAPI instantiates and configures ClientAlerts obj
func NewClientAlertsAPI(cli *EbayClient, l *logrus.Logger) *ClientAlertsAPI {
	return &ClientAlertsAPI{
		clientAlertsURL: "http://clientalerts.ebay.com/ws/ecasvc/ClientAlerts",
		// SANDBOX:
		//clientAlertsURL: "http://clientalerts.sandbox.ebay.com/ws/ecasvc/ClientAlerts",
		client: cli,
		logger: l,
	}
}

// GetClientAlertsAuthToken should be called at least weekly - SOAP
func (api *ClientAlertsAPI) GetClientAlertsAuthToken(ctx context.Context) (*GetClientAlertsAuthTokenResponse, error) {
	response, err := api.client.DoSOAPcall(ctx, &GetClientAlertsAuthTokenRequest{})
	if err != nil {
		return nil, err
	}
	resp := response.(GetClientAlertsAuthTokenResponse)

	// Stores token inside this api object
	api.clientAlertsAuthToken = resp.ClientAlertsAuthToken
	return &resp, nil
}

// Login to Client Alerts API at least daily - REST
func (api *ClientAlertsAPI) Login(ctx context.Context) (*LoginResponse, error) {
	response, err := api.client.DoRESTcall(ctx, api.clientAlertsURL, &LoginRequest{ClientAlertsAuthToken: api.clientAlertsAuthToken}, "GET")
	if err != nil {
		return nil, err
	}
	loginResp := response.(LoginResponse)

	// Stores auth data for next call
	api.sessionID = loginResp.SessionID
	api.sessionData = loginResp.SessionData
	return &loginResp, nil
}

// GetUserAlerts from the Client Alerts API - REST
func (api *ClientAlertsAPI) GetUserAlerts(ctx context.Context) (*GetUserAlertsResponse, error) {
	response, err := api.client.DoRESTcall(ctx, api.clientAlertsURL,
		&GetUserAlertsRequest{
			SessionID:   api.sessionID,
			SessionData: api.sessionData,
		}, "GET")
	if err != nil {
		return nil, err
	}
	alertsResp := response.(GetUserAlertsResponse)

	// Stores auth data for next call
	api.sessionData = alertsResp.SessionData
	return &alertsResp, nil
}
