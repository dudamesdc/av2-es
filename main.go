package main

import (
	"log"
	"os"

	"github.com/dudamesdc/av2-es/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Recuperar porta do servidor a partir das variáveis de ambiente (caso queira configurar a porta também no .env)
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080" // Valor padrão
	}

	router := gin.Default()
	routes.InitRoutes(router)

	log.Printf("Starting server on %s", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
