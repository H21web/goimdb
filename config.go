package main

import (
    "time"
    
    "github.com/Jisin0/filmigo/imdb"
)

type Config struct {
    Port            string        `json:"port"`
    CacheEnabled    bool          `json:"cache_enabled"`
    CacheExpiration time.Duration `json:"cache_expiration"`
}

func NewConfiguredClient(config Config) *imdb.ImdbClient {
    opts := imdb.ImdbClientOpts{
        DisableCaching:   !config.CacheEnabled,
        CacheExpiration:  config.CacheExpiration,
    }
    
    return imdb.NewClient(opts)
}

func DefaultConfig() Config {
    return Config{
        Port:            ":8080",
        CacheEnabled:    true,
        CacheExpiration: 5 * time.Hour,
    }
}
