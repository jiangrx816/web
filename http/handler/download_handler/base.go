package download_handler

import (
	"web/service/download_service"
)

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{
		service: download_service.NewDownloadService(),
	}
}

type DownloadHandler struct {
	service *download_service.DownloadService
}
