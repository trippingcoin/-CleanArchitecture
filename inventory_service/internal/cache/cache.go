package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var InventoryCache *cache.Cache

const DefaultExpiration = 12 * time.Hour

const CleanupInterval = 15 * time.Minute

func InitCache() {
	InventoryCache = cache.New(DefaultExpiration, CleanupInterval)
}
