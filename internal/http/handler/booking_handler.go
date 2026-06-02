package handler

import (
	"net/http"

	"travel-backend/internal/http/dto"
	"travel-backend/internal/service"
	"travel-backend/pkg/response"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	bookingService service.BookingService
}

func NewBookingHandler(bs service.BookingService) *BookingHandler {
	return &BookingHandler{
		bookingService: bs,
	}
}

func (h *BookingHandler) Create(c echo.Context) error {
	var req dto.CreateBookingRequest

	// 1. Bind JSON request body ke DTO
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse("Format request tidak valid"))
	}

	// 2. Eksekusi Business Logic melalui Service
	booking, err := h.bookingService.CreateBooking(req)
	if err != nil {
		return c.JSON(http.StatusConflict, response.ErrorResponse(err.Error()))
	}

	// 3. Kembalikan Response Sukses
	return c.JSON(http.StatusCreated, response.SuccessResponse("Pemesanan berhasil dibuat", booking))
}