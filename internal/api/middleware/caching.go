package middleware

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/baobei23/lk21-go/internal/utils"
	"github.com/patrickmn/go-cache"
)


type responseWriter struct {
	http.ResponseWriter
	body *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

// CacheMiddleware caches the response
func CacheMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.RequestURI
		if cachedData, found := utils.Cache.Get(key); found {
			w.Header().Set("Content-Type", "application/json")
			w.Write(cachedData.([]byte))
			return
		}

		rw := &responseWriter{
			ResponseWriter: w,
			body:          &bytes.Buffer{},
		}

		next.ServeHTTP(rw, r)

		if rw.body.Len() > 0 {
			var data interface{}
			if err := json.Unmarshal(rw.body.Bytes(), &data); err == nil {
				utils.Cache.Set(key, rw.body.Bytes(), cache.DefaultExpiration)
			}
		}
	})
}
