package router

import (
	"net/http"

	"travel-backend/internal/http/handler"
	"travel-backend/pkg/route"
)

func PublicRoutes(pemesananHandler *handler.PemesananHandler, seedHandler *handler.SeedHandler) []route.Route {
	return []route.Route{
		// Skenario A: Tanpa Kontrol Konkurensi (Lost Update)
		{
			Method:  http.MethodPost,
			Path:    "/bookings/none",
			Handler: pemesananHandler.CreateNoControl,
		},
		// Skenario B: Pessimistic Concurrency Control (PCC)
		{
			Method:  http.MethodPost,
			Path:    "/bookings/pcc",
			Handler: pemesananHandler.CreatePCC,
		},
		// Skenario C: Optimistic Concurrency Control (OCC)
		{
			Method:  http.MethodPost,
			Path:    "/bookings/occ",
			Handler: pemesananHandler.CreateOCC,
		},
		// Endpoint untuk mereset/seeding tabel database dari Load Tester
		{
			Method:  http.MethodDelete,
			Path:    "/seed",
			Handler: seedHandler.ResetDatabase,
		},
	}
}