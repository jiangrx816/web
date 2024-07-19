package api_handler

import "web/service/api_service"

func NewApiHandler() *ApiHandler {
	return &ApiHandler{
		service: api_service.NewApiService(),
	}
}

type ApiHandler struct {
	service *api_service.ApiService
}
