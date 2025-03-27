package main

import (
	"fmt"
	"log"

	"github.com/Kevinmajesta/pixel/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// Middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Pastikan tidak ada tanda "/" di akhir
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	// Endpoint untuk menyimpan gambar pixel art
	app.Post("/save", handlers.SavePixelArt)

	fmt.Println("🚀 Server berjalan di http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
