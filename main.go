package main

import (
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Jisin0/filmigo/imdb"
)

type MovieAPI struct {
    imdbClient *imdb.ImdbClient
}

func NewMovieAPI() *MovieAPI {
    client := imdb.NewClient()
    return &MovieAPI{
        imdbClient: client,
    }
}

func main() {
    api := NewMovieAPI()
    router := setupRoutes(api)
    
    log.Println("Starting Movie Detail API server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", router))
}

func setupRoutes(api *MovieAPI) *gin.Engine {
    router := gin.Default()
    
    // Middleware for CORS
    router.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    // API routes
    v1 := router.Group("/api/v1")
    {
        // Movie endpoints
        v1.GET("/movie/:id", api.GetMovieByID)
        v1.GET("/movies/search", api.SearchMovies)
        
        // Person endpoints
        v1.GET("/person/:id", api.GetPersonByID)
        v1.GET("/persons/search", api.SearchPersons)
        
        // Advanced search endpoints
        v1.POST("/movies/advanced-search", api.AdvancedSearchMovies)
        v1.POST("/persons/advanced-search", api.AdvancedSearchPersons)
        
        // Health check
        v1.GET("/health", api.HealthCheck)
    }
    
    return router
}
