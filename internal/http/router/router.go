package router

import (
	"net/http"

	"travel-backend/internal/http/handler"
	"travel-backend/pkg/route"
)

func PublicRoutes(bookingHandler *handler.BookingHandler) []route.Route {
	return []route.Route{
		{
			Method:  http.MethodPost,
			Path:    "/bookings",
			Handler: bookingHandler.Create,
			Roles:   []string{}, // Kosongkan karena tidak ada pengecekan role
		},
		// Endpoint lain untuk kebutuhan simulator bisa ditambahkan di sini
	}
}

func PrivateRoutes() []route.Route {
	// Dibiarkan kosong karena simulator tidak memakai auth
	return []route.Route{}
}