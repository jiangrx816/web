package poetry_handler

import "web/service/poetry_service"

func NewPoetryHandler() *PoetryHandler {
	return &PoetryHandler{
		service: poetry_service.NewPoetryService(),
	}
}

type PoetryHandler struct {
	service *poetry_service.PoetryService
}
