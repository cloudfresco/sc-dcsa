package common

import (
	"github.com/throttled/throttled"
	"github.com/throttled/throttled/v2/store/goredisstore"
	"go.uber.org/zap"
)

// GetHTTPRateLimiter - Get HTTP Rate Limiter
func GetHTTPRateLimiter(store *goredisstore.GoRedisStore, maxRate int, maxBurst int) throttled.HTTPRateLimiter {
	quota := throttled.RateQuota{MaxRate: throttled.PerMin(maxRate), MaxBurst: maxBurst}

	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		log.Error("Error", zap.Int("msgnum", 753), zap.Error(err))
	}

	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{RemoteAddr: true},
	}
	return httpRateLimiter
}
