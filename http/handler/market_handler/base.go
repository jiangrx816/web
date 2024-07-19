package market_handler

import "web/service/market_service"

func NewMarketHandler() *MarketHandler {
	return &MarketHandler{
		service: market_service.NewMarketService(),
	}
}

type MarketHandler struct {
	service *market_service.MarketService
}
