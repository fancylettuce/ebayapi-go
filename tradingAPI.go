package ebayapi

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/fatih/structs"
	"github.com/sirupsen/logrus"
)

// TradingAPI enacts requests to the ebay Client Alerts API
type TradingAPI struct {
	client *EbayClient
	logger *logrus.Logger
}

// NewTradingAPI instantiates and configures Trading obj
func NewTradingAPI(cli *EbayClient, l *logrus.Logger) *TradingAPI {
	return &TradingAPI{
		client: cli,
		logger: l,
	}
}

// GetItem gets single item
func (api *TradingAPI) GetItem(ctx context.Context, req *GetItemRequest) (*GetItemResponse, error) {
	response, err := api.client.DoSOAPcall(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := response.(GetItemResponse)
	return &resp, nil
}

// GetOrders gets a selection of orders, including pagination
func (api *TradingAPI) GetOrders(ctx context.Context, req *GetOrdersRequest) ([]Order, error) {
	ords := []Order{}
	response, err := api.client.DoSOAPcall(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := response.(GetOrdersResponse)
	ords = append(ords, resp.OrderArray.Orders...)

	// Parallel async wait-all pattern for paginated response
	if resp.HasMoreOrders {
		var waitGroup sync.WaitGroup
		for p := 2; p <= resp.PaginationResult.TotalNumberOfPages; p++ {
			waitGroup.Add(1)
			page := p

			// Stagger requests to avoid request limit
			time.Sleep(200 * time.Millisecond)

			go func() {
				req.Pagination = &Pagination{PageNumber: page}
				response, err := api.client.DoSOAPcall(ctx, req)
				if err != nil {
					// Backoff and retry 1 time, then continue
					time.Sleep(500 * time.Millisecond)
					response, err = api.client.DoSOAPcall(ctx, req)
					if err != nil {
						api.logger.WithFields(structs.Map(err)).
							Error("FAILED to get next page of GetOrders from ebay TradingAPI: ", err.Error())
						waitGroup.Done()
						return
					}
				}
				resp = response.(GetOrdersResponse)
				ords = append(ords, resp.OrderArray.Orders...)
				waitGroup.Done()
			}()

		}
		waitGroup.Wait()
	}
	return ords, nil
}

// GetMyeBaySellingPage gets data about our ebay listings - single page
func (api *TradingAPI) GetMyeBaySellingPage(ctx context.Context, req *GetMyeBaySellingRequest) (*GetMyeBaySellingResponse, error) {
	response, err := api.client.DoSOAPcall(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := response.(GetMyeBaySellingResponse)
	return &resp, nil
}

// GetMyeBaySellingAll gets data about our ebay listings with all pages
func (api *TradingAPI) GetMyeBaySellingAll(ctx context.Context, req *GetMyeBaySellingRequest) ([]Item, error) {
	items := []Item{}
	response, err := api.client.DoSOAPcall(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := response.(GetMyeBaySellingResponse)
	items = append(items, resp.ActiveList.ItemArray.Items...)

	// Parallel async wait-all pattern for paginated response
	if resp.ActiveList.PaginationResult.TotalNumberOfPages > 1 {
		var waitGroup sync.WaitGroup
		for p := 2; p <= resp.ActiveList.PaginationResult.TotalNumberOfPages; p++ {
			waitGroup.Add(1)
			page := p

			// Stagger requests to avoid request limit
			time.Sleep(200 * time.Millisecond)

			go func() {
				req.ActiveList.Pagination = &Pagination{PageNumber: page}
				response, err := api.client.DoSOAPcall(ctx, req)
				if err != nil {
					// Backoff and retry 1 time, then continue
					time.Sleep(500 * time.Millisecond)
					response, err = api.client.DoSOAPcall(ctx, req)
					if err != nil {
						api.logger.WithFields(structs.Map(err)).
							Error("FAILED to get next page of GetMyeBaySelling from ebay TradingAPI: ", err.Error())
						waitGroup.Done()
						return
					}
				}
				resp = response.(GetMyeBaySellingResponse)
				items = append(items, resp.ActiveList.ItemArray.Items...)
				waitGroup.Done()
			}()

		}
		waitGroup.Wait()
	}
	return items, nil
}

// ReviseFixedPriceItem gets data about our ebay listings - single page
func (api *TradingAPI) ReviseFixedPriceItem(ctx context.Context, item *Item) (*ReviseFixedPriceItemResponse, error) {
	req := &ReviseFixedPriceItemRequest{Item: item}

	response, err := api.client.DoSOAPcall(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := response.(ReviseFixedPriceItemResponse)
	return &resp, nil
}

// ReviseInventoryStatus gets data about our ebay listings - single page
func (api *TradingAPI) ReviseInventoryStatus(ctx context.Context, invRevisions []*InventoryStatus) (*ReviseInventoryStatusResponse, error) {
	req := &ReviseInventoryStatusRequest{}
	for i, item := range invRevisions {
		// Add revised item data to request, max 4
		if i >= 4 {
			break
		}

		if item.ItemID == "" {
			return nil, errors.New("ERROR[ReviseInventoryStatus]: ItemID value missing")
		}
		if item.SKU == "" {
			return nil, errors.New("ERROR[ReviseInventoryStatus]: SKU value missing")
		}

		req.InventoryStatus = append(req.InventoryStatus, item)
	}

	response, err := api.client.DoSOAPcall(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := response.(ReviseInventoryStatusResponse)
	return &resp, nil
}

// CompleteSale gets data about our ebay listings - single page
func (api *TradingAPI) CompleteSale(ctx context.Context, req *CompleteSaleRequest) (*CompleteSaleResponse, error) {
	response, err := api.client.DoSOAPcall(ctx, req)
	if err != nil {
		return nil, err
	}
	resp := response.(CompleteSaleResponse)
	return &resp, nil
}
