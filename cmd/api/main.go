package main

import (
	"log"
	"net/http"
	"os"

	"api-pokemon-meta-go/internal/handler"
	"api-pokemon-meta-go/internal/repository"
	"api-pokemon-meta-go/internal/service"
	"api-pokemon-meta-go/pkg/database"

	_ "api-pokemon-meta-go/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag"
)

// @title Pokémon Meta API
// @version 1.0
// @description API de analítica competitiva de Pokémon migrada a Go.
// @host localhost:8080
// @BasePath /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No se encontró archivo .env, usando variables de entorno")
	}

	db := database.Init()

	repo := repository.NewPostgresPokemonRepository(db)
	svc := service.NewPokemonService(repo)
	pokemonHandler := handler.NewPokemonHandler(svc)

	router := gin.Default()

	router.GET("/debug-swagger", func(c *gin.Context) {
		doc, err := swag.ReadDoc("swagger")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error_real_de_swagger": err.Error(),
				"mensaje":               "¡Toma captura de este error y pásamelo!",
			})
			return
		}
		c.Data(http.StatusOK, "application/json", []byte(doc))
	})

	api := router.Group("/api/v1")
	{
		api.GET("/meta/snapshot", pokemonHandler.GetMetaSnapshot)
		api.GET("/anti-meta/speed-creepers", pokemonHandler.GetSpeedCreepers)
		api.GET("/anti-meta/wallbreakers", pokemonHandler.GetWallbreakers)
		api.GET("/analytics/trend", pokemonHandler.GetPokemonTrend)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Pokemon API corriendo en el puerto %s", port)
	router.Run(":" + port)
}
