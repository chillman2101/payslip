package cache

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	CACHE_DURATION = 2 * time.Minute
)

type CacheHelper struct {
	client *redis.Client
}

func NewCacheHelper(redisClient *redis.Client) *CacheHelper {
	return &CacheHelper{
		client: redisClient,
	}
}

// generateParamsHash creates a consistent hash from map of parameters
func (c *CacheHelper) generateParamsHash(params map[string]string) string {
	if len(params) == 0 {
		return "empty"
	}

	// Get all keys
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build sorted key-value pairs
	var pairs []string
	for _, k := range keys {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, params[k]))
	}

	// Create hash
	hasher := md5.New()
	hasher.Write([]byte(strings.Join(pairs, "&")))
	return hex.EncodeToString(hasher.Sum(nil))
}

// generateKey creates a unique cache key based on the base key and parameters
func (c *CacheHelper) generateKey(baseKey string, params map[string]string) string {
	paramsHash := c.generateParamsHash(params)
	return fmt.Sprintf("%s:%s", baseKey, paramsHash)
}

// GetCache retrieves data from cache
func (c *CacheHelper) GetCache(key string, params map[string]string) (interface{}, error) {
	cacheKey := c.generateKey(key, params)
	val, err := c.client.Get(context.Background(), cacheKey).Result()
	if err != nil {
		return nil, err
	}

	var result interface{}
	if err := json.Unmarshal([]byte(val), &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SetCache stores data in cache
func (c *CacheHelper) SetCache(key string, params map[string]string, data interface{}) error {
	cacheKey := c.generateKey(key, params)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return c.client.Set(context.Background(), cacheKey, dataBytes, CACHE_DURATION).Err()
}

// DeleteCache removes data from cache
func (c *CacheHelper) DeleteCache(key string, params map[string]string) error {
	cacheKey := c.generateKey(key, params)
	return c.client.Del(context.Background(), cacheKey).Err()
}

// ScanCache scans data from cache by key
func (c *CacheHelper) ScanCache(key string, cursor uint64) ([]string, uint64, error) {
	return c.client.Scan(context.Background(), cursor, key, 0).Result()
}

// DeleteCache removes data from cache
func (c *CacheHelper) DeleteCacheWithoutGenerateKey(key string) error {
	cacheKey := key
	return c.client.Del(context.Background(), cacheKey).Err()
}
