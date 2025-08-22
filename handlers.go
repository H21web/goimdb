package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/Jisin0/filmigo/imdb"
)

// Response structures
type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
    Message string      `json:"message,omitempty"`
}

type SearchRequest struct {
    Query          string `json:"query" binding:"required"`
    IncludeVideos  bool   `json:"include_videos"`
}

// GetMovieByID retrieves detailed movie information by IMDb ID
func (api *MovieAPI) GetMovieByID(c *gin.Context) {
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

// SearchMovies performs a basic movie search
func (api *MovieAPI) SearchMovies(c *gin.Context) {
    query := c.Query("q")
    if query == "" {
        c.JSON(http.StatusBadRequest, APIResponse{
            Success: false,
            Error:   "Search query 'q' parameter is required",
        })
        return
    }

    includeVideos := c.Query("include_videos") == "true"
    
    config := &imdb.SearchConfigs{
        IncludeVideos: includeVideos,
    }

    results, err := api.imdbClient.SearchTitles(query, config)
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

// GetPersonByID retrieves detailed person information by IMDb ID
func (api *MovieAPI) GetPersonByID(c *gin.Context) {
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

// SearchPersons performs a basic person search
func (api *MovieAPI) SearchPersons(c *gin.Context) {
    query := c.Query("q")
    if query == "" {
        c.JSON(http.StatusBadRequest, APIResponse{
            Success: false,
            Error:   "Search query 'q' parameter is required",
        })
        return
    }

    includeVideos := c.Query("include_videos") == "true"
    
    config := &imdb.SearchConfigs{
        IncludeVideos: includeVideos,
    }

    results, err := api.imdbClient.SearchNames(query, config)
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

// AdvancedSearchMovies performs advanced movie search with filters
func (api *MovieAPI) AdvancedSearchMovies(c *gin.Context) {
    var opts imdb.AdvancedSearchTitleOpts
    
    if err := c.ShouldBindJSON(&opts); err != nil {
        c.JSON(http.StatusBadRequest, APIResponse{
            Success: false,
            Error:   "Invalid request body: " + err.Error(),
        })
        return
    }

    results, err := api.imdbClient.AdvancedSearchTitle(&opts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, APIResponse{
            Success: false,
            Error:   "Advanced search failed: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, APIResponse{
        Success: true,
        Data:    results,
    })
}

// AdvancedSearchPersons performs advanced person search with filters
func (api *MovieAPI) AdvancedSearchPersons(c *gin.Context) {
    var opts imdb.AdvancedSearchNameOpts
    
    if err := c.ShouldBindJSON(&opts); err != nil {
        c.JSON(http.StatusBadRequest, APIResponse{
            Success: false,
            Error:   "Invalid request body: " + err.Error(),
        })
        return
    }

    results, err := api.imdbClient.AdvancedSearchName(&opts)
    if err != nil {
        c.JSON(http.StatusInternalServerError, APIResponse{
            Success: false,
            Error:   "Advanced search failed: " + err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, APIResponse{
        Success: true,
        Data:    results,
    })
}

// HealthCheck endpoint for API health monitoring
func (api *MovieAPI) HealthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, APIResponse{
        Success: true,
        Message: "Movie Detail API is running",
    })
}
