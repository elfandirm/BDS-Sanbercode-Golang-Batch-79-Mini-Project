package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"formative-13/config"
	"formative-13/handlers"
	"formative-13/routes"
)

func main() {
	// Koneksi database
	db, err := config.ConnectDatabase()

	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	defer db.Close()

	println("Database berhasil terhubung")

	// Membuat Gin
	router := gin.Default()

	// Membuat handler bioskop
	bioskopHandler := handlers.NewBioskopHandler(db)

	// Setup routes
	routes.SetupRoutes(router, bioskopHandler)

	// Menjalankan server
	err = router.Run(":8080")

	if err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}