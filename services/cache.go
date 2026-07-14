package services

import (
	"be_dashboard/database"
	"context"
	"encoding/json"
	"time"
)

const CacheTTL = 5 * time.Minute

// GetFromCache mengambil data dari Redis cache
// Returns: data (bytes), found (bool), error
func GetFromCache(ctx context.Context, cacheKey string) ([]byte, bool, error) {
	val, err := database.Redis.Get(ctx, cacheKey).Result()
	if err != nil {
		// Cache miss - key tidak ditemukan
		if err.Error() == "redis: nil" {
			return nil, false, nil
		}
		// Error lain (connection error, etc)
		return nil, false, err
	}

	return []byte(val), true, nil
}

// SetCache menyimpan data ke Redis dengan TTL 5 menit
func SetCache(ctx context.Context, cacheKey string, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return database.Redis.Set(ctx, cacheKey, string(jsonData), CacheTTL).Err()
}

// DeleteCache menghapus data dari Redis
func DeleteCache(ctx context.Context, cacheKey string) error {
	return database.Redis.Del(ctx, cacheKey).Err()
}

// GenerateCacheKey membuat cache key berdasarkan userID dan filter
func GenerateCacheKey(userID, prefix string, filter interface{}) (string, error) {
	filterJSON, err := json.Marshal(filter)
	if err != nil {
		return "", err
	}
	return prefix + ":" + userID + ":" + string(filterJSON), nil
}
