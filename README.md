# ebayapi-go
API Client library for eBay seller services including ClientAlerts, Trade, etc. APIs 


Initialize Client with Credentials:

    ebayClient = ebayapi.NewProductionClient(nil)
    mustMapEnv(&ebayClient.AppID, "APP_ID", "")
    mustMapEnv(&ebayClient.DevID, "DEV_ID", "")
    mustMapEnv(&ebayClient.CertID, "CERT_ID", "")
    mustMapEnv(&ebayClient.AuthToken, "TOKEN", "")
    
    func mustMapEnv(target *string, envKey string, useDefault string) {
      v := os.Getenv(envKey)
      if v == "" {
      	if useDefault == "" {
   			  panic(fmt.Sprintf("%s env var not set", envKey))
   		  }
        v = useDefault
      }
    	*target = v
    }
    
    
Initialize APIs

    alertsAPI := ebayapi.NewClientAlertsAPI(ebayClient, log)
    tradingAPI := ebayapi.NewTradingAPI(ebayClient, log)
