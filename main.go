package main

import (
	"fmt"
	"log"
	"os"

	"github.com/khrees2412/convas/database"
	"github.com/khrees2412/convas/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func init() {
    // loads values from .env into the system
    if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}


func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	PORT := os.Getenv("PORT") 
	fmt.Println("Application started...")
	app.Listen(fmt.Sprintf(":%s", PORT))
}
