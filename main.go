package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"formative-13/config"
	"formative-13/handlers"
	"formative-13/routes"
)

func main() {
	db, err := config.ConnectDatabase()

	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}

	defer db.Close()

	log.Println("Database berhasil terhubung")

	router := gin.Default()

	bioskopHandler := handlers.NewBioskopHandler(db)

	routes.SetupRoutes(router, bioskopHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server berjalan di port:", port)

	err = router.Run(":" + port)

	if err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
