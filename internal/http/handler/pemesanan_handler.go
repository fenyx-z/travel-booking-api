package handler

import (
	"net/http"

	"travel-backend/internal/http/dto"
	"travel-backend/internal/service"
	"travel-backend/pkg/response"

	"github.com/labstack/echo/v4"
)

type PemesananHandler struct {
	pemesananService service.PemesananService
}

func NewPemesananHandler(ps service.PemesananService) *PemesananHandler {
	return &PemesananHandler{
		pemesananService: ps,
	}
}

func (h *PemesananHandler) CreateNoControl(c echo.Context) error {
	var req dto.CreatePemesananRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Format request tidak valid"))
	}

	pemesanan, err := h.pemesananService.CreatePemesananNoControl(req)
	if err != nil {
		return c.JSON(http.StatusConflict, response.ErrorResponse(http.StatusConflict, err.Error()))
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse("Pemesanan berhasil dibuat (No Control)", pemesanan))
}

func (h *PemesananHandler) CreatePCC(c echo.Context) error {
	var req dto.CreatePemesananRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Format request tidak valid"))
	}

	pemesanan, err := h.pemesananService.CreatePemesananPCC(req)
	if err != nil {
		return c.JSON(http.StatusConflict, response.ErrorResponse(http.StatusConflict, err.Error()))
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse("Pemesanan berhasil dibuat (PCC)", pemesanan))
}

func (h *PemesananHandler) CreateOCC(c echo.Context) error {
	var req dto.CreatePemesananRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse(http.StatusBadRequest, "Format request tidak valid"))
	}

	pemesanan, err := h.pemesananService.CreatePemesananOCC(req)
	if err != nil {
		return c.JSON(http.StatusConflict, response.ErrorResponse(http.StatusConflict, err.Error()))
	}

	return c.JSON(http.StatusCreated, response.SuccessResponse("Pemesanan berhasil dibuat", pemesanan))
}
