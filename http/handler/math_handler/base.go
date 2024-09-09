package math_handler

import "web/service/math_service"

func NewMathHandler() *MathHandler {
	return &MathHandler{
		service: math_service.NewMathService(),
	}
}

type MathHandler struct {
	service *math_service.MathService
}
