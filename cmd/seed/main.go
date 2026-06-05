package main

import (
	"fmt"
	"log"
	"time"

	"travel-backend/config"
	"travel-backend/internal/entity"
	"travel-backend/pkg/database"
)

func main() {
	// 1. Load Config & Connect DB
	cfg, err := config.NewConfig(".env")
	if err != nil {
		log.Fatalf("Gagal memuat konfigurasi: %v", err)
	}

	db, err := database.InitDatabase(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}

	fmt.Println("Menjalankan Auto Migrate...")
	// 2. Drop tabel lama agar data selalu fresh saat seeder dijalankan
	db.Migrator().DropTable(&entity.Booking{}, &entity.Schedule{}, &entity.Route{}, &entity.User{})
	
	// 3. Migrate ulang tabel
	err = db.AutoMigrate(&entity.User{}, &entity.Route{}, &entity.Schedule{}, &entity.Booking{})
	if err != nil {
		log.Fatalf("Migrasi gagal: %v", err)
	}

	fmt.Println("Menyuntikkan Data Dummy...")

	// 4. Buat Dummy User
	user := entity.User{
		Name:  "Tester Simulator",
		Email: "tester@simulator.local",
	}
	db.Create(&user)

	// 5. Buat Dummy Route
	route := entity.Route{
		Origin:      "Semarang",
		Destination: "Pangkalan Bun",
		DistanceKm:  650,
	}
	db.Create(&route)

	// 6. Buat Dummy Schedule (Ini target utama race condition kita)
	schedule := entity.Schedule{
		RouteID:        route.ID,
		DepartureTime:  time.Now().Add(48 * time.Hour), // Berangkat 2 hari lagi
		ArrivalTime:    time.Now().Add(50 * time.Hour),
		Price:          350000.00,
		TotalSeats:     50, // Kapasitas 50 kursi
		AvailableSeats: 50,
		Version:        1,  // Versi awal untuk OCC
	}
	db.Create(&schedule)

	fmt.Printf("✅ Seeding selesai!\n")
	fmt.Printf("- User ID: %d\n", user.ID)
	fmt.Printf("- Route ID: %d (%s -> %s)\n", route.ID, route.Origin, route.Destination)
	fmt.Printf("- Schedule ID: %d (Sisa Kursi: %d)\n", schedule.ID, schedule.AvailableSeats)
}