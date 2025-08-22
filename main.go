package main

import (
    "log"
    "net/http"
    
    "github.com/gin-gonic/gin"
    "github.com/Jisin0/filmigo/imdb"
)

type MovieAPI struct {
    imdbClient *imdb.ImdbClient
}

type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func main() {
    client := imdb.NewClient()
    api := &MovieAPI{imdbClient: client}
    
    router := gin.Default()
    
    // CORS middleware
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
    
    v1 := router.Group("/api/v1")
    {
        v1.GET("/health", api.healthCheck)
        v1.GET("/movie/:id", api.getMovie)
        v1.GET("/person/:id", api.getPerson)
        v1.GET("/search/movies", api.searchMovies)
        v1.GET("/search/persons", api.searchPersons)
    }
    
    // Use PORT environment variable for Render
    port := "8080"
    if envPort := gin.Mode(); envPort != "" {
        port = "8080"
    }
    
    log.Printf("Starting Movie Detail API on port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}

func (api *MovieAPI) healthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, APIResponse{
        Success: true,
        Data:    "Movie Detail API is running",
    })
}

func (api *MovieAPI) getMovie(c *gin.Context) {
    movieID := c.Param("id")
    if movieID == "" {
        c.JSON(http.StatusBadRequest, APIResponse{
            Success: false,
            Error:   "Movie ID is required",
        })
        return
    }

    movie, err := api.imdbClient.GetMovie(movieID)
    if err != nil {
        c.JSON(http.StatusNotFound, APIResponse{
            Success: false,
            Error:   "Movie not found: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, APIResponse{
        Success: true,
        Data:    movie,
    })
}

func (api *MovieAPI) getPerson(c *gin.Context) {
    personID := c.Param("id")
    if personID == "" {
        c.JSON(http.StatusBadRequest, APIResponse{
            Success: false,
            Error:   "Person ID is required",
        })
        return
    }

    person, err := api.imdbClient.GetPerson(personID)
    if err != nil {
        c.JSON(http.StatusNotFound, APIResponse{
            Success: false,
            Error:   "Person not found: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, APIResponse{
        Success: true,
        Data:    person,
    })
}

func (api *MovieAPI) searchMovies(c *gin.Context) {
    query := c.Query("q")
    if query == "" {
        c.JSON(http.StatusBadRequest, APIResponse{
            Success: false,
            Error:   "Search query 'q' parameter is required",
        })
        return
    }

    results, err := api.imdbClient.SearchTitles(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, APIResponse{
            Success: false,
            Error:   "Search failed: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, APIResponse{
        Success: true,
        Data:    results,
    })
}

func (api *MovieAPI) searchPersons(c *gin.Context) {
    query := c.Query("q")
    if query == "" {
        c.JSON(http.StatusBadRequest, APIResponse{
            Success: false,
            Error:   "Search query 'q' parameter is required",
        })
        return
    }

    results, err := api.imdbClient.SearchNames(query)
    if err != nil {
        c.JSON(http.StatusInternalServerError, APIResponse{
            Success: false,
            Error:   "Search failed: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, APIResponse{
        Success: true,
        Data:    results,
    })
}
