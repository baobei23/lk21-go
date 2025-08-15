package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Cache is a global cache instance
var Cache = cache.New(24*time.Hour, 25*time.Hour)
