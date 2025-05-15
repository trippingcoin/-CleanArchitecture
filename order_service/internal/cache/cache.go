package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var OrderCache *cache.Cache

const DefaultExpiration = 12 * time.Hour

const CleanupInterval = 15 * time.Minute

func InitCache() {
	OrderCache = cache.New(DefaultExpiration, CleanupInterval)
}
